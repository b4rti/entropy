package main

import (
	"fmt"
	"os"

	"github.com/b4rti/entropy/cluster"

	"github.com/b4rti/entropy/plugin"
)

func main() {
	c := &Config{}

	if _, err := os.Stat("/etc/entropy/config.yml"); os.IsNotExist(err) {
		c = CreateDefaultConfig()
	} else {
		c = LoadConfig()
	}

	if !ecluster.CheckNetwork(c.NetworkName) {
		ecluster.CreateNetwork(c.NetworkName)
	}

	go eplugin.NewEntropyPlugin().Serve()

	fmt.Scanln()
}
