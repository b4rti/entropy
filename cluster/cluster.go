package ecluster

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
)

// GetInfo - GetInfo return a *json.Decoder attached to net.Response.Body of a Info Request
func GetInfo() *types.Info {

	d := doRequest("http://info")
	i := &types.Info{}
	d.Decode(i)

	return i
}

// GetNodes - GetNodes return a *json.Decoder attached to net.Response.Body of a Nodes Request
func GetNodes() *[]swarm.Node {

	d := doRequest("http://swarm/nodes")
	n := &[]swarm.Node{}
	d.Decode(n)

	return n
}

func doRequest(path string) *json.Decoder {
	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/run/docker.sock")
			},
		},
	}

	res, err := c.Get(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	return json.NewDecoder(res.Body)
}
