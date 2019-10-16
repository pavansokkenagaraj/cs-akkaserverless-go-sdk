//
// Copyright 2019 Lightbend Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudstate

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/cloudstateio/go-support/cloudstate/protocol"
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	filedescr "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	SupportLibraryVersion = "0.1.0"
	SupportLibraryName    = "cloudstate-go-support"
)

// CloudState is an instance of a CloudState User Function
type CloudState struct {
	server                   *grpc.Server
	entityDiscoveryResponder *EntityDiscoveryService
	eventSourcedHandler      *EventSourcedHandler
}

// New returns a new CloudState instance.
func New(options Options) (*CloudState, error) {
	cs := &CloudState{
		server:                   grpc.NewServer(),
		entityDiscoveryResponder: newEntityDiscoveryResponder(options),
		eventSourcedHandler:      NewEventSourcedHandler(),
	}
	protocol.RegisterEntityDiscoveryServer(cs.server, cs.entityDiscoveryResponder)
	log.Println("RegisterEntityDiscoveryServer")
	protocol.RegisterEventSourcedServer(cs.server, cs.eventSourcedHandler)
	log.Println("RegisterEventSourcedServer")
	return cs, nil
}

// Options go get a CloudState instance configured.
type Options struct {
	ServiceName    string
	ServiceVersion string
}

// DescriptorConfig configures service and dependent descriptors.
type DescriptorConfig struct {
	Service        string
	ServiceMsg     descriptor.Message
	Domain         []string
	DomainMessages []descriptor.Message
}

func (dc DescriptorConfig) AddDomainMessage(m descriptor.Message) DescriptorConfig {
	dc.DomainMessages = append(dc.DomainMessages, m)
	return dc
}

func (dc DescriptorConfig) AddDomainDescriptor(filename string) DescriptorConfig {
	dc.Domain = append(dc.Domain, filename)
	return dc
}

// RegisterEventSourcedEntity registers an event sourced entity for CloudState.
func (cs *CloudState) RegisterEventSourcedEntity(ese *EventSourcedEntity, config DescriptorConfig) (err error) {
	ese.registerOnce.Do(func() {
		if err = ese.initZeroValue(); err != nil {
			return
		}
		if err = cs.eventSourcedHandler.registerEntity(ese); err != nil {
			return
		}
		if err = cs.entityDiscoveryResponder.registerEntity(ese, config); err != nil {
			return
		}
	})
	return
}

// Run runs the CloudState instance.
func (cs *CloudState) Run() error {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		return fmt.Errorf("unable to get environment variable \"HOST\"")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return fmt.Errorf("unable to get environment variable \"PORT\"")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	log.Printf("starting grpcServer at: %s:%s", host, port)
	if e := cs.server.Serve(lis); e != nil {
		return fmt.Errorf("failed to grpcServer.Serve for: %v", lis)
	}
	return nil
}

// EntityDiscoveryService implements the CloudState discovery protocol.
type EntityDiscoveryService struct {
	fileDescriptorSet *filedescr.FileDescriptorSet
	entitySpec        *protocol.EntitySpec
	message           *descriptor.Message
}

// newEntityDiscoveryResponder returns a new and initialized EntityDiscoveryService.
func newEntityDiscoveryResponder(options Options) *EntityDiscoveryService {
	responder := &EntityDiscoveryService{}
	responder.entitySpec = &protocol.EntitySpec{
		Entities: make([]*protocol.Entity, 0),
		ServiceInfo: &protocol.ServiceInfo{
			ServiceName:           options.ServiceName,
			ServiceVersion:        options.ServiceVersion,
			ServiceRuntime:        fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
			SupportLibraryName:    SupportLibraryName,
			SupportLibraryVersion: SupportLibraryVersion,
		},
	}
	responder.fileDescriptorSet = &filedescr.FileDescriptorSet{
		File: make([]*filedescr.FileDescriptorProto, 0),
	}
	return responder
}

// Discover returns an entity spec for
func (r *EntityDiscoveryService) Discover(c context.Context, pi *protocol.ProxyInfo) (*protocol.EntitySpec, error) {
	log.Printf("Received discovery call from sidecar [%s w%s] supporting CloudState %v.%v\n",
		pi.ProxyName,
		pi.ProxyVersion,
		pi.ProtocolMajorVersion,
		pi.ProtocolMinorVersion,
	)
	for _, filename := range []string{
		"google/protobuf/empty.proto",
		"google/protobuf/any.proto",
		"google/protobuf/descriptor.proto",
		"google/api/annotations.proto",
		"google/api/http.proto",
		"cloudstate/event_sourced.proto",
		"cloudstate/entity.proto",
		"cloudstate/entity_key.proto",
	} {
		if err := r.registerFileDescriptorProto(filename); err != nil {
			return nil, err
		}
	}
	log.Printf("Responding with: %v\n", r.entitySpec.GetServiceInfo())
	return r.entitySpec, nil
}

// ReportError logs any user function error reported by the CloudState proxy.
func (r *EntityDiscoveryService) ReportError(c context.Context, fe *protocol.UserFunctionError) (*empty.Empty, error) {
	log.Printf("ReportError: %v\n", fe)
	return &empty.Empty{}, nil
}

func (r *EntityDiscoveryService) updateSpec() (err error) {
	protoBytes, err := proto.Marshal(r.fileDescriptorSet)
	if err != nil {
		return errors.New("unable to Marshal FileDescriptorSet")
	}
	r.entitySpec.Proto = protoBytes
	return nil
}

func (r *EntityDiscoveryService) resolveFileDescriptors(dc DescriptorConfig) error {
	// service
	if dc.Service != "" {
		if err := r.registerFileDescriptorProto(dc.Service); err != nil {
			return err
		}
	} else {
		if dc.ServiceMsg != nil {
			if err := r.registerFileDescriptor(dc.ServiceMsg); err != nil {
				return err
			}
		}
	}
	// and dependent domain descriptors
	for _, dp := range dc.Domain {
		if err := r.registerFileDescriptorProto(dp); err != nil {
			return err
		}
	}
	for _, dm := range dc.DomainMessages {
		if err := r.registerFileDescriptor(dm); err != nil {
			return err
		}
	}
	return nil
}

func (r *EntityDiscoveryService) registerEntity(e *EventSourcedEntity, config DescriptorConfig) error {
	if err := r.resolveFileDescriptors(config); err != nil {
		return fmt.Errorf("failed to resolveFileDescriptor for DescriptorConfig: %+v: %w", config, err)
	}
	persistenceID := e.entityName
	if e.PersistenceID != "" {
		persistenceID = e.PersistenceID
	}
	r.entitySpec.Entities = append(r.entitySpec.Entities, &protocol.Entity{
		EntityType:    EventSourced,
		ServiceName:   e.ServiceName,
		PersistenceId: persistenceID,
	})
	return r.updateSpec()
}

func (r *EntityDiscoveryService) registerFileDescriptorProto(filename string) error {
	descriptorProto, err := unpackFile(proto.FileDescriptor(filename))
	if err != nil {
		return fmt.Errorf("failed to registerFileDescriptorProto for filename: %s: %w", filename, err)
	}
	r.fileDescriptorSet.File = append(r.fileDescriptorSet.File, descriptorProto)
	return r.updateSpec()
}

func (r *EntityDiscoveryService) registerFileDescriptor(msg descriptor.Message) error {
	fd, _ := descriptor.ForMessage(msg) // this can panic
	if r := recover(); r != nil {
		return fmt.Errorf("descriptor.ForMessage panicked (%v) for: %+v", r, msg)
	}
	r.fileDescriptorSet.File = append(r.fileDescriptorSet.File, fd)
	return nil
}
