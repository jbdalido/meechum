package main

import (
	"github.com/jbdalido/meechum"
)

import (
	"flag"
	"log"
)

const (
	version = "0.1.0"
)

var (
	backend = flag.String("b", "consul", "Type of supported backend (etcd or consul)")
	hosts   = flag.String("h", "127.0.0.1:8500", "Host:Port of the selected backend")
	groups  = flag.String("g", "basics", "comma separated values for groups")
)

func main() {
	log.Printf("Starting Meechum v%s ...", version)

	// Setup the runtime and connect to the backend
	mc, err := meechum.NewRuntime(backen, hosts)
	if err != nil {
		log.Fatalf("Meechum failed to start: %s", err)
	}

	// Subscribe to group, retrieve informations
	err = meechum.Subscribe(groups)
	if err != nil {
		log.Fatalf("Meechum started but failed to subscribe to groups %s : %s", groups, err)
	}

	// Runtime loop
	err = meechum.Run()
	if err != nil {
		log.Fatalf("Meechum stopped", err)
	}
}
