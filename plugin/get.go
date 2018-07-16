package entropy

import (
	"github.com/docker/go-plugins-helpers/volume"
)

// Get - return information about give volume
func (d Driver) Get(req *volume.GetRequest) (*volume.GetResponse, error) {
	debug(req)

	vol, err := getVolume(req.Name)
	if err != nil {
		return nil, err
	}

	res := &volume.GetResponse{
		Volume: vol,
	}

	return res, nil
}

func getVolume(n string) (*volume.Volume, error) {
	var err error
	var c, m string
	var s map[string]interface{}
	var v *volume.Volume

	c, err = getCreatedAt(n)
	if err != nil {
		return nil, err
	}

	m, err = getMountpoint(n)
	if err != nil {
		return nil, err
	}

	s, err = getStatus(n)
	if err != nil {
		return nil, err
	}

	v = &volume.Volume{
		CreatedAt:  c,
		Mountpoint: m,
		Name:       n,
		Status:     s,
	}

	return v, err
}

func getCreatedAt(n string) (string, error) {
	var p string
	var err error

	return p, err
}

func getMountpoint(n string) (string, error) {
	var p string
	var err error

	return p, err
}

func getStatus(n string) (map[string]interface{}, error) {
	var s map[string]interface{}
	var err error

	return s, err
}
