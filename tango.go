package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Tango is an open-source MQTT server

Usage:
	tango [options] command [arg...]

The available commands are:

    serve         start the Tango server

The available options are:

    -port         port that Tango listens to for connections
`)
	os.Exit(1)
}

func serve(port *int) {
	fmt.Fprintf(os.Stdout, "Listening on %v\n", *port)
	// TODO
}

func main() {
	port := flag.Int("port", 1883, "the port number")
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	cmd := flag.Args()[0]

	switch cmd {
	case "serve":
		serve(port)
	default:
		fmt.Fprintf(os.Stderr, "Invalid command\n\n")
		usage()
	}
}
