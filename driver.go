package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

const (
	DDSVPath = "/var/lib/ddsv"
)

type DDVSDriver struct {
	volume.Driver
}
