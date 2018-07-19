package ecluster

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
)

// GetInfo - GetInfo return a *json.Decoder attached to net.Response.Body of a Info Request
func GetInfo() (*types.Info, error) {

	d, err := doGETDocker("http://swarm/info")
	if err != nil {
		return nil, err
	}

	i := &types.Info{}
	d.Decode(i)

	return i, err
}

// GetNodes - GetNodes return a *json.Decoder attached to net.Response.Body of a Nodes Request
func GetNodes() (*[]swarm.Node, error) {

	d, err := doGETDocker("http://swarm/nodes")
	if err != nil {
		return nil, err
	}

	n := &[]swarm.Node{}
	d.Decode(n)

	return n, err
}

func doGETDocker(p string) (*json.Decoder, error) {
	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/run/docker.sock")
			},
		},
	}

	res, err := c.Get(p)
	if err != nil {
		return nil, err
	}

	return json.NewDecoder(res.Body), err
}

func doPOSTDocker(p string, d []byte) (*json.Decoder, error) {
	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/run/docker.sock")
			},
		},
	}

	b := bytes.NewBuffer(d)
	res, err := c.Post(p, "application/json", b)
	if err != nil {
		return nil, err
	}

	return json.NewDecoder(res.Body), err
}
