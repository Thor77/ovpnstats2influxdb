package main

import (
	"flag"
	"log"

	"github.com/thor77/ovpnstats2influxdb/ovpnstats2influxdb"
)

func main() {
	// read configuration
	pathFlagPointer := flag.String("path", "", "Path to openvpn-status.log")

	flag.Parse()

	if *pathFlagPointer == "" {
		log.Fatal("No path provided")
	}
        err := ovpnstats2influxdb.RunTelegraf(*pathFlagPointer)
        if err != nil {
            log.Fatal(err)
        }
}
