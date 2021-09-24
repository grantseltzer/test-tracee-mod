package main

import (
	"fmt"

	"github.com/grantseltzer/tracee/tracee-ebpf/tracee"
)

func main() {

	fmt.Println(tracee.Config{
		PerfBufferize: 2,
	})
}
