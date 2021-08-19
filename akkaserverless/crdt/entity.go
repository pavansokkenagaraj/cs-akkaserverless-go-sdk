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

package crdt

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/protocol"
)

// Entity captures an Entity with its ServiceName. It is used to be registered
// as an CRDT entity on an Akkaserverless instance.
type Entity struct {
	// ServiceName is the fully qualified name of the service that implements
	// this entities interface. Setting it is mandatory.
	ServiceName ServiceName
	// EntityFunc creates a new entity.
	EntityFunc          func(id EntityID) EntityHandler
	PassivationStrategy protocol.PassivationStrategy
}

type Option func(s *Entity)

func (e *Entity) Options(options ...Option) {
	for _, opt := range options {
		opt(e)
	}
}

func WithPassivationStrategyTimeout(duration time.Duration) Option {
	return func(e *Entity) {
		e.PassivationStrategy = protocol.PassivationStrategy{
			Strategy: &protocol.PassivationStrategy_Timeout{
				Timeout: &protocol.TimeoutPassivationStrategy{
					Timeout: duration.Milliseconds(),
				},
			},
		}
	}
}

// EntityHandler has to be implemented by any type that wants to get
// registered as a crdt.Entity
// tag::entity-handler[]
type EntityHandler interface {
	HandleCommand(ctx *CommandContext, name string, msg proto.Message) (*any.Any, error)
	Default(ctx *Context) (CRDT, error)
	Set(ctx *Context, state CRDT) error
}

// end::entity-handler[]
