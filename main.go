package main

import (
	"fmt"
	"log"
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

	ecluster.GetInfo()
	ecluster.GetNodes()

	if !ecluster.CheckNetwork(c.NetworkName) {
		id, err := ecluster.CreateNetwork(c.NetworkName)
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(id)
	}

	eplugin.NewEntropyPlugin().Serve()
}
