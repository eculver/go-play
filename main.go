package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const routersFile = "/etc/uber/hyperbahn/hosts.json"

var defaultRouters = []string{"10.10.0.1:23100"}

func main() {
	fmt.Printf("it's a map!: %v\n", map[string]string{
		"foo": "bar",
	})

	fmt.Printf("Initial Nodes: %s\n", getInitialNodes())
}

func getInitialNodes() []string {
	if _, err := os.Stat(routersFile); os.IsNotExist(err) {
		return defaultRouters
	}

	bs, err := ioutil.ReadFile(routersFile)
	if err != nil {
		fmt.Println("COULD NOT READ ROUTERS FILE")
		return defaultRouters
	}

	var initialNodes []string
	json.Unmarshal(bs, &initialNodes)

	return initialNodes
}
