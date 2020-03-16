// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goscanner checks ports connectivity on hosts.
package goscanner

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

// CheckRequest is the structure to use to pass a port check request.
type CheckRequest struct {
	Protocol string
	Host     string
	Port     string
	Timeout  int
}

// CheckPortCtx connects to the desided address.
func CheckPortCtx(ctx context.Context, r CheckRequest) error {

	if _, err := strconv.Atoi(r.Port); err != nil {
		return fmt.Errorf("%s is not a usable port", r.Port)
	}

	var addr string
	addr = r.Host + ":" + r.Port

	_, err := net.DialTimeout(r.Protocol, addr, time.Duration(r.Timeout)*time.Second)

	if err == nil {
		log.Printf("Connected on %s\n", addr)
	}

	return nil
}
