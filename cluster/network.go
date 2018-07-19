package ecluster

import (
	"encoding/json"

	"github.com/docker/docker/api/types"
)

func CheckNetwork(s string) bool {

	nl, err := listNetworks()
	if err != nil {
		return false
	}

	for _, n := range *nl {
		if n.Name == s {
			return false
		}
	}

	return false
}

func CreateNetwork(s string) (string, error) {

	n := map[string]interface{}{
		"Name":   s,
		"Driver": "overlay",
		"Options": map[string]string{
			"encrypted": "true",
		},
	}

	b, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	d, err := doPOSTDocker("http://swarm/networks/create", b)

	res := &struct {
		Id string
	}{}

	err = d.Decode(res)

	return res.Id, err
}

func listNetworks() (*[]types.NetworkResource, error) {
	d, err := doGETDocker("http://swarm/networks")

	n := &[]types.NetworkResource{}

	err = d.Decode(n)
	if err != nil {
		return nil, err
	}

	return n, err
}
