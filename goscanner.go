package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var Version = "Development"

var wg sync.WaitGroup

func main() {

	var v = flag.Bool("v", false, "shows vwersion")
	var protocol = flag.String("t", "tcp", "ip protocol to use")

	flag.Parse()

	if *v {
		fmt.Printf("goscanner version: %s\n", Version)
		fmt.Printf("Author: %s\n", "Alberto Bregliano")
		os.Exit(0)
	}

	argoments := flag.Args()
	ip := argoments[0]
	ports := argoments[1:]

	for _, port := range ports {
		wg.Add(1)
		time.Sleep(5 * time.Millisecond)
		addr := ip + ":" + port
		go connect(*protocol, addr)
	}

	wg.Wait()
}

func connect(protocol, addr string) {
	defer wg.Done()

	_, err := net.Dial(protocol, addr)

	if err == nil {
		log.Printf("Connected on %s\n", addr)
	}
}
