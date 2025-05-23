// Copyright 2022-2023 The Inspektor Gadget authors
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

// Package tracer is deprecated.
//
// Deprecated: Switch to image-based gadgets instead. Check
// https://github.com/inspektor-gadget/inspektor-gadget/tree/main/examples/gadgets
package tracer

import (
	gadgetregistry "github.com/inspektor-gadget/inspektor-gadget/pkg/gadget-registry"
	"github.com/inspektor-gadget/inspektor-gadget/pkg/gadgets"
	"github.com/inspektor-gadget/inspektor-gadget/pkg/gadgets/trace/bind/types"
	"github.com/inspektor-gadget/inspektor-gadget/pkg/params"
	"github.com/inspektor-gadget/inspektor-gadget/pkg/parser"
)

const (
	ParamPID          = "pid"
	ParamPorts        = "ports"
	ParamIgnoreErrors = "ignore-errors"
)

type GadgetDesc struct {
	gadgets.GadgetDeprecated
}

func (g *GadgetDesc) Name() string {
	return "bind"
}

func (g *GadgetDesc) Category() string {
	return gadgets.CategoryTrace
}

func (g *GadgetDesc) Type() gadgets.GadgetType {
	return gadgets.TypeTrace
}

func (g *GadgetDesc) Description() string {
	return "Trace socket bindings"
}

func (g *GadgetDesc) ParamDescs() params.ParamDescs {
	return params.ParamDescs{
		{
			Key:          ParamPID,
			Title:        "PID",
			DefaultValue: "0",
			Description:  "Show only bind events generated by this particular PID",
			TypeHint:     params.TypeInt32,
		},
		{
			Key:          ParamPorts,
			Alias:        "P",
			DefaultValue: "",
			Description:  "Trace only bind events involving these ports",
			Validator:    params.ValidateSlice(params.ValidateUintRange(1, 65535)),
		},
		{
			Key:          ParamIgnoreErrors,
			Title:        "Ignore Errors",
			Alias:        "i",
			DefaultValue: "true",
			Description:  "Show only events where the bind succeeded",
			TypeHint:     params.TypeBool,
		},
	}
}

func (g *GadgetDesc) Parser() parser.Parser {
	return parser.NewParser[types.Event](types.GetColumns())
}

func (g *GadgetDesc) EventPrototype() any {
	return &types.Event{}
}

func init() {
	gadgetregistry.Register(&GadgetDesc{})
}
