// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goscanner checks ports connectivity on hosts.
package goscanner_test

import (
	"context"
	"fmt"
	"goscanner"
	"log"
	"net"
	"testing"
)

func init() {
	// Start the new server.
	srv, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1123})
	if err != nil {
		log.Println("error starting TCP server")
		return
	}

	// Run the server in Goroutine to stop tests from blocking
	// test execution.
	go func() {
		srv.AcceptTCP()
	}()
}

func TestCheckPortCtx(t *testing.T) {

	type args struct {
		ctx context.Context
		r   goscanner.CheckRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"one port", args{context.TODO(), goscanner.CheckRequest{Protocol: "tcp", Host: "localhost", Port: "8080", Timeout: 2}}, true},
		{"not a port", args{context.TODO(), goscanner.CheckRequest{Protocol: "tcp", Host: "localhost", Port: "notaport", Timeout: 2}}, true},
		{"open port", args{context.TODO(), goscanner.CheckRequest{Protocol: "tcp", Host: "localhost", Port: "1123", Timeout: 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := goscanner.CheckPortCtx(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CheckPortCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// ExampleCheckPortCtx è un esempio di utilizzo del package goscanner.
func ExampleCheckPortCtx() {

	ctx := context.TODO()

	openPort := goscanner.CheckRequest{
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     "1123",
	}

	closedPort := goscanner.CheckRequest{
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     "1124",
	}

	err1 := goscanner.CheckPortCtx(ctx, openPort)
	err2 := goscanner.CheckPortCtx(ctx, closedPort)

	fmt.Println(err1)
	fmt.Println(err2)
	// Output:
	// <nil>
	// porta 1124 chiusa
}
