// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"goscanner"
	"os"
	"sync"
)

// Version of the app will be updated via -ldflags at build time.
var Version = "Development"

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var v = flag.Bool("v", false, "shows version")
	var protocol = flag.String("p", "tcp", "ip protocol to use")
	var timeout = flag.Int("t", 1, "timeout in seconds for single portcheck")

	flag.Parse()

	if *v {
		fmt.Printf("goscanner version: %s\n", Version)
		fmt.Printf("Author: %s\n", "Alberto Bregliano")
		os.Exit(0)
	}

	arguments := flag.Args()
	ip := arguments[0]
	ports := arguments[1:]

	var wg sync.WaitGroup
	wg.Add(len(ports))

	for _, port := range ports {
		var r goscanner.CheckRequest
		r = goscanner.CheckRequest{
			Protocol: *protocol,
			Host:     ip,
			Port:     port,
			Timeout:  *timeout,
		}

		go func() {
			defer wg.Done()
			goscanner.CheckPortCtx(ctx, r)
		}()
	}

	wg.Wait()
}
