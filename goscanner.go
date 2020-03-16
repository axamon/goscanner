// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goscanner

import (
	"log"
	"net"
	"time"
)

// CheckPort connects to the desided address.
func CheckPort(protocol, addr string, timeout int) {

	_, err := net.DialTimeout(protocol, addr, time.Duration(timeout)*time.Second)

	if err == nil {
		log.Printf("Connected on %s\n", addr)
	}

	return
}
