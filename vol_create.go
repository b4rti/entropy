package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

func (d DDVSDriver) Create(req *volume.CreateRequest) error {
	debug(req)

	return nil
}
