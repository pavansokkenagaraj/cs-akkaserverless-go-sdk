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

package synth

import (
	"github.com/golang/protobuf/proto"
	"github.com/lightbend/akkaserverless-go-sdk/tck-cloudstate/crdt"
)

func flagRequest(messages ...proto.Message) *crdt.FlagRequest {
	r := &crdt.FlagRequest{
		Actions: make([]*crdt.FlagRequestAction, 0),
	}
	for _, i := range messages {
		switch t := i.(type) {
		case *crdt.Get:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.FlagRequestAction{Action: &crdt.FlagRequestAction_Get{Get: t}})
		case *crdt.Delete:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.FlagRequestAction{Action: &crdt.FlagRequestAction_Delete{Delete: t}})
		case *crdt.FlagEnable:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.FlagRequestAction{Action: &crdt.FlagRequestAction_Enable{Enable: t}})
		default:
			panic("no type matched")
		}
	}
	return r
}