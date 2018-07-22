package eplugin

import (
	"encoding/json"
	"fmt"
	"log"
	"os/user"
	"strconv"

	"github.com/docker/go-plugins-helpers/volume"
)

const (
	// EntropyPath - path to workdir
	EntropyPath = "/var/lib/ddsv"
)

// EntropyPlugin - plugin struct
type EntropyPlugin struct {
	handler *volume.Handler
}

// NewEntropyPlugin - return new Plugin
func NewEntropyPlugin() *EntropyPlugin {
	d := Driver{}
	h := volume.NewHandler(d)

	return &EntropyPlugin{
		handler: h,
	}
}

// e *Serve - serves unix socket
func (e *EntropyPlugin) Serve() {
	u, err := user.Lookup("root")
	if err != nil {
		log.Panicln(err.Error())
	}

	g, err := strconv.Atoi(u.Gid)
	if err != nil {
		log.Panicln(err.Error())
	}

	err = e.handler.ServeUnix("entropy", g)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func debug(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	fmt.Println(string(b))
}
