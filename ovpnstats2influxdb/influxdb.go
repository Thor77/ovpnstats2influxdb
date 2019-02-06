package ovpnstats2influxdb

import (
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

// createBatchPoints converts metrics into InfluxDB points
func createBatchPoints(measurementName string, metrics []Metric) ([]*influxdb.Point, error) {
	var result []*influxdb.Point
	measurement := measurementName
	timestamp := time.Now()
	for _, metric := range metrics {
		point, err := influxdb.NewPoint(measurement, metric.Tags, metric.Fields, timestamp)
		if err != nil {
			return nil, err
		}
		result = append(result, point)
	}
	return result, nil
}
