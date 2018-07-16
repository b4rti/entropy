package entropy

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// Remove - removes the volume
func (d Driver) Remove(req *volume.RemoveRequest) error {
	debug(req)

	return nil
}
