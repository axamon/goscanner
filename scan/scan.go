// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scan verifica se esistono porte in ascolto su host di rete.
package scan

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

// CheckRequest è la struttura per creare una richiesta di verifica di
// porta logica Port su indirizzo Host tramite Protocol in medo di Timeout secondi.
type CheckRequest struct {
	Protocol string
	Host     string
	Port     string
	Timeout  int
}

// CheckPortCtx avvia la verifica sui parametri passati in r.
func CheckPortCtx(ctx context.Context, r CheckRequest) error {

	// se la porta logica non è usabile restituisce un errore ed esce.
	if _, err := strconv.Atoi(r.Port); err != nil {
		return fmt.Errorf("%s non è una porta logica utilizzabile", r.Port)
	}

	// addr è l'indirizzo a cui connttersi composto da host:porta.
	var addr string
	addr = r.Host + ":" + r.Port

	// Avvia la connessione per testare se la porta è aperta o meno.
	_, err := net.DialTimeout(r.Protocol, addr, time.Duration(r.Timeout)*time.Second)

	if err != nil {
		return fmt.Errorf("porta %s chiusa", r.Port)
	}

	// Se non ci sono errori restituisce un log con la data di test.
	log.Printf("Connesso con successo su: %s, porta aperta: %s \n", addr, r.Port)

	return nil
}
