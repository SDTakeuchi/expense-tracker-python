package main

import (
	"http/net"
	"log"
)

func main() {
	http.HandleFunc("/")
	log.Fatalf()
}