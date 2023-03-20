package main

import (
	"fmt"
	"os"
)

type Endpoint struct {
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ip_address"`
	LastLogin string `json:"last_login"`
}

func getEndpoint() Endpoint {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	endpoint := Endpoint{
		Hostname:  hostname,
		IPAddress: "1.1.1.1",
		LastLogin: "19 MAR 23",
	}
	return endpoint
}

func main() {
	fmt.Println(getEndpoint())

	// check EventForwarding-Plugin log

	// sendEndpointInfo
}
