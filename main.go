package main

import (
	"fmt"

	"github.com/b4rti/entropy/plugin"
)

func main() {
	go entropy.NewEntropyPlugin().Serve()

	fmt.Scanln()
}
