package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	// Leer el número de filósofos desde variables de entorno (Configurable)
	numStr := os.Getenv("NUM_FILOSOFOS")
	numFilosofos := 5 // Valor por defecto si no se pasa la variable
	if n, err := strconv.Atoi(numStr); err == nil && n >= 2 {
		numFilosofos = n
	}

	fmt.Printf("Iniciando simulación con %d filósofos...\n", numFilosofos)
	monitor := NewMonitor(numFilosofos)

	// Lanzar a los filósofos como hilos (Goroutines)
	for i := 0; i < numFilosofos; i++ {
		go func(id int) {
			for {
				// Pensar
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

				monitor.TakeForks(id)
				fmt.Printf(">>> Filósofo %d está COMIENDO <<<\n", id)

				// Comer
				time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)

				monitor.PutForks(id)
				fmt.Printf("<<< Filósofo %d terminó de comer y está PENSANDO >>>\n", id)
			}
		}(i)
	}

	// Endpoint para el Frontend / Backend Web
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // Para que el frontend pueda leerlo
		json.NewEncoder(w).Encode(monitor.GetStates())
	})

	fmt.Println("Servidor backend corriendo en http://localhost:8080/status")
	http.ListenAndServe(":8080", nil)
}
