package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// Unmount - unmounts thew give volume
func (d DDVSDriver) Unmount(req *volume.UnmountRequest) error {
	debug(req)

	return nil
}
