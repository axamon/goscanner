// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"goscanner"
	"os"
	"strconv"
	"sync"
	"time"
)

// Version of the app will be updated via -ldflags at build time.
var Version = "Development"

func main() {

	var v = flag.Bool("v", false, "shows version")
	var protocol = flag.String("p", "tcp", "ip protocol to use")
	var timeout = flag.Int("t", 1, "timeout in seconds for single portcheck")

	flag.Parse()

	if *v {
		fmt.Printf("goscanner version: %s\n", Version)
		fmt.Printf("Author: %s\n", "Alberto Bregliano")
		os.Exit(0)
	}

	argoments := flag.Args()
	ip := argoments[0]
	ports := argoments[1:]

	var wg sync.WaitGroup

	for _, port := range ports {

		time.Sleep(5 * time.Millisecond)

		if _, err := strconv.Atoi(port); err != nil {
			fmt.Printf("Error %s is not a usable port\n", port)
			continue
		}

		addr := ip + ":" + port
		wg.Add(1)
		go func() {
			defer wg.Done()
			goscanner.CheckPort(*protocol, addr, *timeout)
		}()
	}

	wg.Wait()
}
