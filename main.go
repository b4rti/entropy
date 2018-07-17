package main

import (
	"fmt"

	"github.com/b4rti/entropy/plugin"
)

func main() {
	go eplugin.NewEntropyPlugin().Serve()

	fmt.Scanln()

}
