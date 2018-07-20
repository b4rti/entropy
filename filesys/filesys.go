package efilesys

import "github.com/docker/docker/api/types"

// GetLocalVolume - get local volume plugin struct
func GetLocalVolume() *types.Volume {
	v := &types.Volume{}

	return v
}

// GetClusterVolume - get cluster volume plugin struct
func GetClusterVolume() *types.Volume {
	v := &types.Volume{}

	return v
}

// GetLocalVolumeList - get local list volume plugin struct
func GetLocalVolumeList() *[]types.Volume {
	v := &[]types.Volume{}

	return v
}

// GetClusterVolumeList - get cluster list volume plugin struct
func GetClusterVolumeList() *[]types.Volume {
	v := &[]types.Volume{}

	return v
}
