package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// Create - creates a new volume with the given name and options
func (d DDVSDriver) Create(req *volume.CreateRequest) error {
	debug(req)

	return nil
}
