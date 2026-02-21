package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/tunchxz/CC3064-Laboratorio-Sincronizacion/simulation"
)

// Constantes de validación
const (
	MIN_PHILOSOPHERS = 2
	MAX_PHILOSOPHERS = 100
	MIN_DURATION     = 60   // segundos
	MAX_DURATION     = 1200 // 20 minutos máximo
)

// Variable global con sincronización para manejo concurrente
var (
	currentSimulation *simulation.Simulation
	simulationMutex   sync.Mutex
)

// StartHandler maneja las peticiones para iniciar una nueva simulación
func StartHandler(w http.ResponseWriter, r *http.Request) {
	// Validar método HTTP
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Use GET or POST", http.StatusMethodNotAllowed)
		return
	}

	nParam := r.URL.Query().Get("n")
	durationParam := r.URL.Query().Get("duration")

	// Validación estricta del número de filósofos
	n := 5 // valor por defecto
	if nParam != "" {
		val, err := strconv.Atoi(nParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid parameter 'n': must be an integer, got '%s'", nParam), http.StatusBadRequest)
			return
		}
		if val < MIN_PHILOSOPHERS {
			http.Error(w, fmt.Sprintf("Invalid parameter 'n': must be at least %d philosophers", MIN_PHILOSOPHERS), http.StatusBadRequest)
			return
		}
		if val > MAX_PHILOSOPHERS {
			http.Error(w, fmt.Sprintf("Invalid parameter 'n': maximum %d philosophers allowed", MAX_PHILOSOPHERS), http.StatusBadRequest)
			return
		}
		n = val
	}

	// Validación estricta de la duración
	duration := 60 * time.Second // valor por defecto
	if durationParam != "" {
		val, err := strconv.Atoi(durationParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid parameter 'duration': must be an integer (seconds), got '%s'", durationParam), http.StatusBadRequest)
			return
		}
		if val < MIN_DURATION {
			http.Error(w, fmt.Sprintf("Invalid parameter 'duration': minimum %d seconds", MIN_DURATION), http.StatusBadRequest)
			return
		}
		if val > MAX_DURATION {
			http.Error(w, fmt.Sprintf("Invalid parameter 'duration': maximum %d seconds allowed", MAX_DURATION), http.StatusBadRequest)
			return
		}
		duration = time.Duration(val) * time.Second
	}

	// Manejo sincronizado de la variable global simulation
	simulationMutex.Lock()
	defer simulationMutex.Unlock()

	// Verificar si ya hay una simulación en ejecución
	if currentSimulation != nil && currentSimulation.IsRunning() {
		http.Error(w, "A simulation is already running. Stop it first with /stop", http.StatusConflict)
		return
	}

	currentSimulation = simulation.New(n, duration)
	currentSimulation.Start()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":       "started",
		"philosophers": n,
		"duration":     duration.String(),
		"message":      fmt.Sprintf("Simulation started with %d philosophers for %v", n, duration),
	})
}

// StopHandler maneja las peticiones para detener la simulación actual
func StopHandler(w http.ResponseWriter, r *http.Request) {
	// Validar método HTTP
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Use GET or POST", http.StatusMethodNotAllowed)
		return
	}

	// Manejo sincronizado de la variable global simulation
	simulationMutex.Lock()
	defer simulationMutex.Unlock()

	if currentSimulation == nil {
		http.Error(w, "No simulation exists", http.StatusNotFound)
		return
	}

	if !currentSimulation.IsRunning() {
		http.Error(w, "Simulation is not running", http.StatusConflict)
		return
	}

	currentSimulation.Stop()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "stopped",
		"message": "Simulation stopped successfully",
	})
}

// StatusHandler maneja las peticiones para obtener el estado actual de la simulación
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Validar método HTTP
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Use GET", http.StatusMethodNotAllowed)
		return
	}

	// Manejo sincronizado de la variable global simulation
	simulationMutex.Lock()
	defer simulationMutex.Unlock()

	if currentSimulation == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "no_simulation",
			"message": "No simulation exists. Start one with /start?n=5&duration=60",
		})
		return
	}

	currentSimulation.WriteStatus(w)
}
