package main

import (
	"log"

	"github.com/tomaszjonak/api-experiment/payments-client/client"
)

func main() {
	config := &client.TransportConfig{
		Host:     "127.0.0.1:50123",
		BasePath: "/api/v0",
		Schemes:  []string{"http"},
	}

	cl := client.NewHTTPClientWithConfig(nil, config)

	resp, err := cl.Operations.GetPayments(nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, payment := range resp.Payload {
		log.Println(*payment)
	}
}
