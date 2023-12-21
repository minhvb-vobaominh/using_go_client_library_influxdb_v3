package main

import (
	"context"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

func main() {
	const INFLUXDB_URL = "https://eu-central-1-1.aws.cloud2.influxdata.com"
	const INFLUXDB_API_TOKEN = "XPNjWrBSTP4q7KtJf-6zBxjjYiG3EnufZj594QzhrktyeYCC_QOKrHauRLKNZ3OJB5p-8uZaLtX5P3PTZrzCBA=="
	const INFLUXDB_BUCKET = "node_metrics"

	// Create a new client using an InfluxDB server base URL and an authentication token
	client, error_connect := influxdb3.New(influxdb3.ClientConfig{
		Host:     INFLUXDB_URL,
		Token:    INFLUXDB_API_TOKEN,
		Database: INFLUXDB_BUCKET,
	})

	if error_connect != nil {
		panic(error_connect)
	}

	// Close client at the end and escalate error if present
	defer func(client *influxdb3.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	// Create point using full params constructor
	p := influxdb3.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45.0},
		time.Now())

	// Write point synchronously
	error_statement := client.WritePoints(context.Background(), p)
	if error_statement != nil {
		panic(error_statement)
	}
}
