// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package tracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadIterUDPv4 returns the embedded CollectionSpec for IterUDPv4.
func LoadIterUDPv4() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_IterUDPv4Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load IterUDPv4: %w", err)
	}

	return spec, err
}

// LoadIterUDPv4Objects loads IterUDPv4 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *IterUDPv4Objects
//     *IterUDPv4Programs
//     *IterUDPv4Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadIterUDPv4Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadIterUDPv4()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// IterUDPv4Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type IterUDPv4Specs struct {
	IterUDPv4ProgramSpecs
	IterUDPv4MapSpecs
}

// IterUDPv4Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type IterUDPv4ProgramSpecs struct {
	DumpUdp4 *ebpf.ProgramSpec `ebpf:"dump_udp4"`
}

// IterUDPv4MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type IterUDPv4MapSpecs struct {
}

// IterUDPv4Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadIterUDPv4Objects or ebpf.CollectionSpec.LoadAndAssign.
type IterUDPv4Objects struct {
	IterUDPv4Programs
	IterUDPv4Maps
}

func (o *IterUDPv4Objects) Close() error {
	return _IterUDPv4Close(
		&o.IterUDPv4Programs,
		&o.IterUDPv4Maps,
	)
}

// IterUDPv4Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadIterUDPv4Objects or ebpf.CollectionSpec.LoadAndAssign.
type IterUDPv4Maps struct {
}

func (m *IterUDPv4Maps) Close() error {
	return _IterUDPv4Close()
}

// IterUDPv4Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadIterUDPv4Objects or ebpf.CollectionSpec.LoadAndAssign.
type IterUDPv4Programs struct {
	DumpUdp4 *ebpf.Program `ebpf:"dump_udp4"`
}

func (p *IterUDPv4Programs) Close() error {
	return _IterUDPv4Close(
		p.DumpUdp4,
	)
}

func _IterUDPv4Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed iterudpv4_bpfel.o
var _IterUDPv4Bytes []byte
