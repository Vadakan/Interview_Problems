package main

import (
	"fmt"
	"github.com/Consul/Discover"
	"github.com/hashicorp/consul/api"
)

func main() {

	config := api.DefaultConfig()
	config.Address = "localhost:8500"

	// faked name and port for example
	cli, err := Discover.NewClient(config, "localhost", "messaging", 8080)
	fmt.Println(cli)
	if err != nil {
		panic(err)
	}

	if err := Discover.Exec(cli); err != nil {
		panic(err)
	}

	// faked name and port for example
	cli, err = Discover.NewClient(config, "localhost", "payment", 8080)
	fmt.Println(cli)
	if err != nil {
		panic(err)
	}

	if err := Discover.Exec(cli); err != nil {
		panic(err)
	}

}