package main

import (
	"fmt"

	"github.com/grantseltzer/tracee/tracee-ebpf/external"
	"github.com/grantseltzer/tracee/tracee-ebpf/tracee"
	"github.com/grantseltzer/tracee/tracee-rules/engine"
	"github.com/grantseltzer/tracee/tracee-rules/types"
)

func main() {
	fmt.Println(external.Event{})
	fmt.Println(tracee.Tracee{})
	fmt.Println(types.SignatureMetadata{})
	fmt.Println(engine.Engine{})
}
