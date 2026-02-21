package simulation

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/tunchxz/CC3064-Laboratorio-Sincronizacion/monitor"
)

// Simulation gestiona la ejecución de la simulación de Dining Philosophers
type Simulation struct {
	Monitor  *monitor.DiningMonitor
	n        int
	running  bool
	stopChan chan struct{}
	wg       sync.WaitGroup
	duration time.Duration
}

// New crea una nueva simulación con n filósofos y duración especificada
func New(n int, duration time.Duration) *Simulation {
	return &Simulation{
		Monitor:  monitor.NewDiningMonitor(n),
		n:        n,
		stopChan: make(chan struct{}),
		duration: duration,
	}
}

// philosopher simula el comportamiento de un filósofo individual
func (s *Simulation) philosopher(id int) {
	defer s.wg.Done()

	for {
		select {
		case <-s.stopChan:
			return
		default:
		}

		// Pensar (tiempo variable por filósofo)
		time.Sleep(time.Duration(500+id*100) * time.Millisecond)

		// Intentar comer
		s.Monitor.Pickup(id)

		// Comer (tiempo variable por filósofo)
		time.Sleep(time.Duration(800+id*100) * time.Millisecond)

		// Terminar de comer
		s.Monitor.Putdown(id)
	}
}

// Start inicia la simulación con goroutines para cada filósofo
func (s *Simulation) Start() {
	if s.running {
		return
	}
	s.running = true

	// Iniciar goroutine para cada filósofo
	for i := 0; i < s.n; i++ {
		s.wg.Add(1)
		go s.philosopher(i)
	}

	// Goroutine para detener automáticamente después de la duración
	go func() {
		time.Sleep(s.duration)
		s.Stop()
	}()
}

// Stop detiene la simulación de forma segura
func (s *Simulation) Stop() {
	if !s.running {
		return
	}
	close(s.stopChan)
	s.wg.Wait()
	s.running = false
	log.Println("Simulation stopped")
}

// IsRunning retorna si la simulación está actualmente en ejecución
func (s *Simulation) IsRunning() bool {
	return s.running
}

// WriteStatus escribe el estado actual de la simulación en formato JSON
func (s *Simulation) WriteStatus(w http.ResponseWriter) {
	data := s.Monitor.Snapshot()
	// Agregar el campo running al snapshot
	data["running"] = s.running
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
