package eplugin

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// Path - returns the mountpoint
func (d Driver) Path(req *volume.PathRequest) (*volume.PathResponse, error) {
	debug(req)

	res := &volume.PathResponse{
		Mountpoint: "",
	}

	return res, nil
}
