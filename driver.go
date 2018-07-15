package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

const (
	// DDSVPath - WTF golint
	DDSVPath = "/var/lib/ddsv"
)

// DDVSDriver - driver struct
type DDVSDriver struct {
	volume.Driver
}
