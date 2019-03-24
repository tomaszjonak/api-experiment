package main

import (
	"log"
	"time"

	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

const (
	listenEndpoint = "127.0.0.1:50123"
)

type routeLogger struct {
	handler http.Handler
}

func (rl *routeLogger) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println(request)
	rl.handler.ServeHTTP(writer, request)
}

func main() {
	log.Println("Hello, world")

	log.Println("Creating db connection pool")
	cluster := gocql.NewCluster("127.0.0.1")
	_, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connections created")

	log.Println("Creating HTTP server")
	router := mux.NewRouter()
	router.Use(func(handler http.Handler) http.Handler {
		return &routeLogger{handler: handler}
	})
	router.HandleFunc("/payments", func(http.ResponseWriter, *http.Request) {}).Methods("GET")

	server := http.Server{
		Handler:      router,
		Addr:         listenEndpoint,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// go func() {
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	// }()
	log.Println("HTTP server up and running")

	log.Println("Bye, world")
}
