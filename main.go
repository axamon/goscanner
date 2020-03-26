// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

// Version of the app will be updated via -ldflags at build time.
var Version = "Development"

func main() {

	// ctx è il contesto padre da passare a ogni goroutine.
	ctx, cancel := context.WithCancel(context.Background())

	// cancel chiude ogni processo pendente a fine funzione.
	defer cancel()

	// v è un flag booleano per richiedere la versione attuale di goscanner.
	var v = flag.Bool("v", false, "mostra la versione attuale di goscanner")

	// protocol è un flag per selezionare il protocollo di connessione da usare.
	var protocol = flag.String("p", "tcp", "protocollo ip da usare per la connessione")

	// timeout è il tempo massimo in secondi per verificare l'apertura della porta
	// sull'host.
	var timeout = flag.Int("t", 1, "timeout in secondi per ogni verifica")

	// Parsa i flag.
	flag.Parse()

	// Se viene richiesta la versione di goscanner la mostra a video ed esce.
	if *v {
		fmt.Printf("goscanner version: %s\n", Version)
		fmt.Printf("Author: %s\n", "Alberto Bregliano")
		os.Exit(0)
	}

	// arguments è la lista di valori che seguono i flags.
	arguments := flag.Args()

	// Se non vengono passati paramtri riepiloga la sintassi.
	if len(arguments) < 1 {
		fmt.Println("Nessun parametro passato")
		fmt.Println("Sintassi: goscanner <ip> <porte ...>")
		os.Exit(1)
	}

	// 	VerificheParalleleCtx esegue i controlli delle porte in parallelo.
	VerificheParalleleCtx(ctx, arguments, *protocol, *timeout)
}
