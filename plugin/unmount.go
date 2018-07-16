package entropy

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// Unmount - unmounts thew give volume
func (d Driver) Unmount(req *volume.UnmountRequest) error {
	debug(req)

	return nil
}
