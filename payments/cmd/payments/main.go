package main

import (
	"flag"
	"log"

	"github.com/tomaszjonak/api-experiment/payments/gen/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/tomaszjonak/api-experiment/payments/gen/restapi"
)

func main() {
	var (
		host = flag.String("host", "127.0.0.1", "Host to listen on")
		port = flag.Int("port", 50123, "Port to listen on")
	)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalf("Couldn't load swagger spec (err: %v)", err)
	}

	api := operations.NewPaymentsAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Host = *host
	server.Port = *port

	defer server.Shutdown()

	flag.Parse()
	if err := server.Serve(); err != nil {
		log.Fatalf("Couldn't start serving (err: %v)", err)
	}
}
