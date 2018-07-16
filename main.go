package main

import (
	"fmt"

	"github.com/b4rti/ddvs/plugin"
)

func main() {
	go entropy.NewEntropyPlugin().Serve()

	fmt.Scanln()
}
