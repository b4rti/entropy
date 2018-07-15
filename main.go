package main

import (
	"encoding/json"
	"fmt"
	"os/user"
	"strconv"

	"github.com/docker/go-plugins-helpers/volume"
)

func main() {
	d := DDVSDriver{}

	h := volume.NewHandler(d)

	u, err := user.Lookup("root")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = h.ServeUnix("ddvs", gid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(err.Error())
}

func debug(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	fmt.Println(string(b))
}
