package ecluster

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
)

func GetInfo() *types.Info {

	b := doRequest("http://info")
	d := json.NewDecoder(b)
	dsi := &types.Info{}
	d.Decode(dsi)

	return dsi
}

func GetNodes() *[]swarm.Node {

	b := doRequest("http://nodes")
	d := json.NewDecoder(b)
	nl := &[]swarm.Node{}
	d.Decode(nl)

	return nl
}

func doRequest(path string) io.ReadCloser {
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

	return res.Body
}
