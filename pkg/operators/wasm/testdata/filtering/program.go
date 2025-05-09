// Copyright 2025 The Inspektor Gadget authors
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
	api "github.com/inspektor-gadget/inspektor-gadget/wasmapi/go"
)

// Keep in sync with pkg/operators/wasm/wasm_test.go TestFiltering
const (
	mntnsDiscarded    = uint64(555)
	mntnsNotDiscarded = uint64(777)
)

//go:wasmexport gadgetStart
func gadgetStart() int32 {
	if ret := api.ShouldDiscardMntNsID(mntnsDiscarded); !ret {
		api.Errorf("mntns should be discarded")
		return 1
	}
	if ret := api.ShouldDiscardMntNsID(mntnsNotDiscarded); ret {
		api.Errorf("mntns should not be discarded")
		return 1
	}
	return 0
}

func main() {}
