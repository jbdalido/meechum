package main

import (
	"github.com/jbdalido/meechum"
	"github.com/jbdalido/meechum/api"
	"github.com/jbdalido/meechum/handler/loghandler"
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
	runtime, err := meechum.NewRuntime(*backend, *hosts)
	if err != nil {
		log.Fatalf("Meechum failed to start: %s", err)
	}

	// Subscribe to group, retrieve informations
	err = runtime.Subscribe([]string{*groups})
	if err != nil {
		log.Fatalf("Meechum started but failed to subscribe to groups %s : %s", *groups, err)
	}

	meechum.RegisterHandler(logHandler.NewLogHandler())

	// Runtime loop
	go func() {
		err = runtime.Run()
		if err != nil {
			log.Fatalf("Meechum stopped", err)
		}
	}()
	apiServer, err := api.NewApi(runtime, "127.0.0.1:9000")
	if err != nil {
		log.Fatalf("Meechum failed to start api: %s", err)
	}
	apiServer.Start()

}
