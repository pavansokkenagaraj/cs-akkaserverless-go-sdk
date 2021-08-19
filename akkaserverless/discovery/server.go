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

// Package discovery implements the Akkaserverless entity discovery server.
package discovery

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	filedescr "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/action"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/crdt"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/eventsourced"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/protocol"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/value"
)

const (
	SupportLibraryName    = "akkaserverless-go-sdk"
	SupportLibraryVersion = "0.0.1"
	ProtocolMajorVersion  = 0
	ProtocolMinorVersion  = 7
)

// EntityDiscoveryServer implements the Akkaserverless discovery protocol.
type EntityDiscoveryServer struct {
	mu                sync.RWMutex
	fileDescriptorSet *filedescr.FileDescriptorSet
	entitySpec        *protocol.Spec

	protocol.UnimplementedDiscoveryServer
}

// NewServer returns a new and initialized EntityDiscoveryServer.
func NewServer(config protocol.Config) *EntityDiscoveryServer {
	return &EntityDiscoveryServer{
		entitySpec: &protocol.Spec{
			Components: make([]*protocol.Component, 0),
			ServiceInfo: &protocol.ServiceInfo{
				ServiceName:           config.ServiceName,
				ServiceVersion:        config.ServiceVersion,
				ServiceRuntime:        fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
				SupportLibraryName:    SupportLibraryName,
				SupportLibraryVersion: SupportLibraryVersion,
				ProtocolMajorVersion:  ProtocolMajorVersion,
				ProtocolMinorVersion:  ProtocolMinorVersion,
			},
		},
		fileDescriptorSet: &filedescr.FileDescriptorSet{
			File: make([]*filedescr.FileDescriptorProto, 0),
		},
	}
}

// Discover returns an entity spec for registered entities.
func (s *EntityDiscoveryServer) Discover(_ context.Context, info *protocol.ProxyInfo) (*protocol.Spec, error) {
	log.Printf("Received discovery call from sidecar [%s %s] supporting Akkaserverless %v.%v\n",
		info.ProxyName,
		info.ProxyVersion,
		info.ProtocolMajorVersion,
		info.ProtocolMinorVersion,
	)
	log.Printf("Responding with: %v\n", s.entitySpec.GetServiceInfo())
	// TODO: s.entitySpec can be written potentially but should not after we started to run the server;
	//  check how to enforce that after protocol.Run has started.
	return s.entitySpec, nil
}

// ReportError logs any user function error reported by the Akkaserverless proxy.
func (s *EntityDiscoveryServer) ReportError(_ context.Context, error *protocol.UserFunctionError) (*empty.Empty, error) {
	log.Printf("ReportError: %v\n", error)
	return &empty.Empty{}, nil
}

func (s *EntityDiscoveryServer) updateSpec() (err error) {
	protoBytes, err := proto.Marshal(s.fileDescriptorSet)
	if err != nil {
		return errors.New("unable to Marshal FileDescriptorSet")
	}
	s.entitySpec.Proto = protoBytes
	return nil
}

func (s *EntityDiscoveryServer) resolveFileDescriptors(config protocol.DescriptorConfig) error {
	if config.Service != "" {
		if err := s.registerFileDescriptorProto(config.Service); err != nil {
			return err
		}
	}
	// Add dependent domain descriptors.
	for _, dp := range config.Domain {
		if err := s.registerFileDescriptorProto(dp); err != nil {
			return err
		}
	}
	for _, dm := range config.DomainMessages {
		if err := s.registerFileDescriptor(dm); err != nil {
			return err
		}
	}
	return nil
}

func (s *EntityDiscoveryServer) RegisterEventSourcedEntity(entity *eventsourced.Entity, config protocol.DescriptorConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.resolveFileDescriptors(config); err != nil {
		return fmt.Errorf("failed to resolve FileDescriptor for DescriptorConfig: %+v: %w", config, err)
	}
	e := &protocol.Component{
		ComponentType: protocol.EventSourcedEntities,
		ServiceName:   entity.ServiceName.String(),
	}
	entitySettings := &protocol.Component_Entity{
		Entity: &protocol.EntitySettings{
			EntityType: entity.PersistenceID,
		},
	}
	if entity.PassivationStrategy.Strategy != nil {
		entitySettings.Entity.PassivationStrategy = &entity.PassivationStrategy
	}
	e.ComponentSettings = entitySettings

	s.entitySpec.Components = append(s.entitySpec.Components, e)
	return s.updateSpec()
}

func (s *EntityDiscoveryServer) RegisterCRDTEntity(entity *crdt.Entity, config protocol.DescriptorConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.resolveFileDescriptors(config); err != nil {
		return fmt.Errorf("failed to resolveFileDescriptor for DescriptorConfig: %+v: %w", config, err)
	}
	e := &protocol.Component{
		ComponentType: protocol.CRDT,
		ServiceName:   entity.ServiceName.String(),
	}

	entitySettings := &protocol.Component_Entity{
		Entity: &protocol.EntitySettings{
			// TODO: as per https://github.com/cloudstateio/go-support/pull/67#issuecomment-749838999 this is temporary.
			EntityType: entity.ServiceName.String(),
		},
	}
	if entity.PassivationStrategy.Strategy != nil {
		entitySettings.Entity.PassivationStrategy = &entity.PassivationStrategy
	}
	e.ComponentSettings = entitySettings

	s.entitySpec.Components = append(s.entitySpec.Components, e)
	return s.updateSpec()
}

func (s *EntityDiscoveryServer) RegisterActionEntity(entity *action.Entity, config protocol.DescriptorConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.resolveFileDescriptors(config); err != nil {
		return fmt.Errorf("failed to resolveFileDescriptor for DescriptorConfig: %+v: %w", config, err)
	}
	s.entitySpec.Components = append(s.entitySpec.Components, &protocol.Component{
		ComponentType: protocol.Action,
		ServiceName:   entity.ServiceName.String(),
	})
	return s.updateSpec()
}

func (s *EntityDiscoveryServer) RegisterValueEntity(entity *value.Entity, config protocol.DescriptorConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.resolveFileDescriptors(config); err != nil {
		return fmt.Errorf("failed to resolveFileDescriptor for DescriptorConfig: %+v: %w", config, err)
	}
	e := &protocol.Component{
		ComponentType: protocol.Value,
		ServiceName:   entity.ServiceName.String(),
	}
	entitySettings := &protocol.Component_Entity{
		Entity: &protocol.EntitySettings{
			EntityType: entity.PersistenceID,
		},
	}
	if entity.PassivationStrategy.Strategy != nil {
		entitySettings.Entity.PassivationStrategy = &entity.PassivationStrategy
	}
	e.ComponentSettings = entitySettings

	s.entitySpec.Components = append(s.entitySpec.Components, e)
	return s.updateSpec()
}

func (s *EntityDiscoveryServer) hasRegistered(filename string) bool {
	for _, f := range s.fileDescriptorSet.File {
		if f.GetName() == filename {
			return true
		}
	}
	return false
}

func (s *EntityDiscoveryServer) registerFileDescriptorProto(filename string) error {
	if s.hasRegistered(filename) {
		return nil
	}
	desc, err := unpackFile(proto.FileDescriptor(filename))
	if err != nil {
		return fmt.Errorf("failed to registerFileDescriptorProto for filename: %s: %w", filename, err)
	}
	s.fileDescriptorSet.File = append(s.fileDescriptorSet.File, desc)
	for _, dep := range desc.Dependency {
		err := s.registerFileDescriptorProto(dep)
		if err != nil {
			return err
		}
	}
	return s.updateSpec()
}

func (s *EntityDiscoveryServer) registerFileDescriptor(msg descriptor.Message) (err error) {
	fd, _ := descriptor.ForMessage(msg) // this can panic.
	if s.hasRegistered(fd.GetName()) {
		return nil
	}
	s.fileDescriptorSet.File = append(s.fileDescriptorSet.File, fd)
	return nil
}
