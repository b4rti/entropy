package entropy

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// List - returns a list of all volumes
func (d Driver) List() (*volume.ListResponse, error) {
	vols := []*volume.Volume{}

	res := &volume.ListResponse{
		Volumes: vols,
	}

	return res, nil
}
