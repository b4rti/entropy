package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config - set by YAML
type Config struct {
	WorkDir          string `yaml:"WorkDir"`
	NetworkName      string `yaml:"NetworkName"`
	NetworkDriver    string `yaml:"NetworkDriver"`
	NetworkRange     string `yaml:"NetworkRange"`
	NetworkEncrypted bool   `yaml:"NetworkEncrypted"`
}

// CreateDefaultConfig - creat default struct
func CreateDefaultConfig(p string) *Config {
	c := Config{
		WorkDir:          "/var/lib/entropy",
		NetworkName:      "entropy",
		NetworkDriver:    "overlay",
		NetworkRange:     "10.255.0.0/16",
		NetworkEncrypted: true,
	}

	d, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	ioutil.WriteFile(p, d, 0644)

	return &c
}

// LoadConfig - config struct
func LoadConfig(p string) *Config {
	d, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	c := &Config{}

	yaml.Unmarshal(d, c)

	return c
}
