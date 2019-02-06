package ovpnstats2influxdb

import (
	"fmt"

	"github.com/thor77/ovpnstats"
)

// RunTelegraf prints Telegraf-compatible metric-output to stdout
func RunTelegraf(path string) error {
	clients, routes, err := ovpnstats.ParseStatusFile(path)
	if err != nil {
		return err
	}

	metrics := createMetrics(clients, routes)

	// convert metrics to influxdb line protocol
	points, err := createBatchPoints("openvpn", metrics)
	if err != nil {
		return err
	}

	// output line protocol lines
	for _, point := range points {
		fmt.Println(point.String())
	}
	return nil
}
