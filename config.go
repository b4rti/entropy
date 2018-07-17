package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	WorkdirPath string
	PluginPath  string
	NetworkName string
}

func CreateDefaultConfig() *Config {
	c := Config{
		WorkdirPath: "/var/lib/entropy",
		PluginPath:  "/var/lib/docker/plugins/entropy",
		NetworkName: "entropy",
	}

	d, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	ioutil.WriteFile("/etc/entropy/config.yml", d, 0644)

	return &c
}

func LoadConfig() *Config {
	d, err := ioutil.ReadFile("/etc/entropy/config.yml")
	if err != nil {
		log.Fatal(err.Error())
	}

	c := &Config{}

	yaml.Unmarshal(d, c)

	return c
}
