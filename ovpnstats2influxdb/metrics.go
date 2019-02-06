package ovpnstats2influxdb

import (
	"github.com/thor77/ovpnstats"
)

// Metric represents metric events with all required information to write it into an InfluxDB
type Metric struct {
	Fields map[string]interface{}
	Tags   map[string]string
}

func createMetrics(clients []ovpnstats.ClientInfo, routes []ovpnstats.RoutingInfo) []Metric {
	return []Metric{Metric{map[string]interface{}{"clients": len(clients), "routes": len(routes)}, nil}}
}
