// Copyright 2024 The Inspektor Gadget authors
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

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/inspektor-gadget/inspektor-gadget/pkg/datasource"
	igjson "github.com/inspektor-gadget/inspektor-gadget/pkg/datasource/formatters/json"
	gadgetcontext "github.com/inspektor-gadget/inspektor-gadget/pkg/gadget-context"
	"github.com/inspektor-gadget/inspektor-gadget/pkg/operators"
	"github.com/inspektor-gadget/inspektor-gadget/pkg/operators/simple"
	grpcruntime "github.com/inspektor-gadget/inspektor-gadget/pkg/runtime/grpc"
)

func do() error {
	const opPriority = 50000
	myOperator := simple.New("myOperator", simple.OnInit(func(gadgetCtx operators.GadgetContext) error {
		for _, d := range gadgetCtx.GetDataSources() {
			jsonFormatter, _ := igjson.New(d)

			d.Subscribe(func(source datasource.DataSource, data datasource.Data) error {
				jsonOutput := jsonFormatter.Marshal(data)
				fmt.Printf("%s\n", jsonOutput)
				return nil
			}, opPriority)
		}
		return nil
	}))

	// Used to close the connection in a clean way
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	gadgetCtx := gadgetcontext.New(
		ctx,
		"ghcr.io/inspektor-gadget/gadget/trace_open:main",
		gadgetcontext.WithDataOperators(myOperator),
	)

	runtime := grpcruntime.New()
	runtimeGlobalParams := runtime.GlobalParamDescs().ToParams()
	runtimeGlobalParams.Get(grpcruntime.ParamRemoteAddress).Set("tcp://127.0.0.1:8888")
	// Unix sockets are also supported:
	//runtimeGlobalParams.Get(grpcruntime.ParamRemoteAddress).Set("unix:///var/run/ig/ig.socket")

	if err := runtime.Init(runtimeGlobalParams); err != nil {
		return fmt.Errorf("runtime init: %w", err)
	}
	defer runtime.Close()

	if err := runtime.RunGadget(gadgetCtx, nil, nil); err != nil {
		return fmt.Errorf("running gadget: %w", err)
	}

	return nil
}

func main() {
	if err := do(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
		os.Exit(1)
	}
}
