package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/b4rti/entropy/cluster"
	"github.com/b4rti/entropy/plugin"
)

func main() {
	go eplugin.NewEntropyPlugin().Serve()

	i := ecluster.GetInfo()
	debug(i)

	n := ecluster.GetNodes()
	debug(n)

	fmt.Scanln()

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
