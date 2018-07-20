package main

import (
	"log"
	"os"
	"path"

	"github.com/b4rti/entropy/cluster"
	"github.com/b4rti/entropy/plugin"
)

const ConfigPath = "/etc/entropy/config.yml"

func main() {
	c := &Config{}

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		c = CreateDefaultConfig(ConfigPath)
	} else {
		c = LoadConfig(ConfigPath)
	}

	if _, err := os.Stat(c.WorkDir); os.IsNotExist(err) {
		os.MkdirAll(c.WorkDir, 0644)
	}
	if _, err := os.Stat(path.Join(c.WorkDir, "volumes")); os.IsNotExist(err) {
		os.MkdirAll(path.Join(c.WorkDir, "volumes"), 0644)
	}
	if _, err := os.Stat(path.Join(c.WorkDir, "mounts")); os.IsNotExist(err) {
		os.MkdirAll(path.Join(c.WorkDir, "mounts"), 0644)
	}

	if !ecluster.CheckNetwork(c.NetworkName) {
		_, err := ecluster.CreateNetwork(c.NetworkName)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	eplugin.NewEntropyPlugin().Serve()
}
