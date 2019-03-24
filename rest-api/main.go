package main

import (
	"log"

	"github.com/gocql/gocql"
)

func main() {
	log.Println("Hello, world")
	cluster := gocql.NewCluster("127.0.0.1")

	_, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bye, world")
}
