package entropy

import "github.com/docker/go-plugins-helpers/volume"

// Capabilities - returns Driver Capabilities - global
func (d Driver) Capabilities() *volume.CapabilitiesResponse {
	caps := volume.Capability{
		Scope: "global",
	}

	res := &volume.CapabilitiesResponse{
		Capabilities: caps,
	}

	return res
}
