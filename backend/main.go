package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tunchxz/CC3064-Laboratorio-Sincronizacion/handlers"
	"github.com/tunchxz/CC3064-Laboratorio-Sincronizacion/middleware"
)

func main() {
	// Configurar logging con timestamps de microsegundos para análisis detallado
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Registrar handlers HTTP con CORS habilitado
	http.HandleFunc("/start", middleware.CORSHandler(handlers.StartHandler))
	http.HandleFunc("/stop", middleware.CORSHandler(handlers.StopHandler))
	http.HandleFunc("/status", middleware.CORSHandler(handlers.StatusHandler))
	
	// Health check endpoint
	http.HandleFunc("/health", middleware.CORSHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"dining-philosophers"}`))
	}))

	// Iniciar servidor HTTP
	log.Println("===========================================")
	log.Println("Dining Philosophers Monitor")
	log.Println("Server running on :8080")
	log.Println("CORS enabled for all origins")
	log.Println("===========================================")
	log.Println("Endpoints:")
	log.Println("  GET/POST /start?n=5&duration=60")
	log.Println("  GET/POST /stop")
	log.Println("  GET      /status")
	log.Println("  GET      /health")
	log.Println("===========================================")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
