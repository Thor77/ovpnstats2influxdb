ovpnstats2influxdb
==================

Read metrics from OpenVPN `openvpn-status.log` and output them as InfluxDB metrics.

## Installation
* Build it yourself with `go build main.go` and move the binary somewhere into your `$PATH`

## Usage
Add it as an exec plugin to your `telegraf.conf`:
```
[[inputs.exec]]
    commands = ["/usr/bin/ovpnstats -path /etc/openvpn/server/openvpn-status.log"]
    data_format = "influx"
    timeout = 1
```
