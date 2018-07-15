package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

func (d DDVSDriver) List() (*volume.ListResponse, error) {
	vols := []*volume.Volume{}

	res := &volume.ListResponse{
		Volumes: vols,
	}

	return res, nil
}
