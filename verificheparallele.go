// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"goscanner/scan"
	"sync"
	
)

// VerificheParalleleCtx esegue le verifiche delle porte in parallelo.
func VerificheParalleleCtx(ctx context.Context, arguments []string, protocol string, timeout int) {

	// ip è l'indirizzo dell'host da contattare.
	ip := arguments[0]

	// ports è la lista di porte da testare.
	ports := arguments[1:]

	// wg è un WaitGroup per la gestione del parallelismo delle goroutines.
	var wg sync.WaitGroup

	// Il numero di goroutines parallele viene fissato a quello di porte da verificare,
	// la lunghezza della slice ports.
	wg.Add(len(ports))

	// Per ogni porta nella lista di porte da testare avvia una richiesta parallela.
	for _, port := range ports {

		// r è una richiesta di verifica di tipo CheckRequest che viene istanziata
		// con il suo ZeroValue.
		var r scan.CheckRequest

		// gli elementi di r vengono fissati.
		r = scan.CheckRequest{
			Protocol: protocol,
			Host:     ip,
			Port:     port,
			Timeout:  timeout,
		}

		// Esegue una go routine anonima.
		go func() {
			// Al termine della gorountine decrementa il contatore wg di uno.
			defer wg.Done()
			// Avvia la verifica propagando il contesto ctx e inviando la
			// richiesta r.
			scan.CheckPortCtx(ctx, r)
		}()
	}

	// Attende che il contatore wg arrivi a zero prima di proseguire,
	// Tutte le goroutine create devono terminare autonomamente prima di poter proseguire.
	wg.Wait()
}
