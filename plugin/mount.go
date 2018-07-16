package entropy

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// Mount - mounts give volume and return the mountpoint
func (d Driver) Mount(req *volume.MountRequest) (*volume.MountResponse, error) {
	debug(req)

	res := &volume.MountResponse{
		Mountpoint: "",
	}

	return res, nil
}
