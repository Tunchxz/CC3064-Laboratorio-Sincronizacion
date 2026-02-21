package main

import (
	"fmt"
	"sync"
)

type Monitor struct {
	mu    sync.Mutex
	cond  []*sync.Cond
	state []string // "THINKING", "HUNGRY", "EATING"
	n     int
}

func NewMonitor(n int) *Monitor {
	m := &Monitor{
		state: make([]string, n),
		cond:  make([]*sync.Cond, n),
		n:     n,
	}
	for i := 0; i < n; i++ {
		m.state[i] = "THINKING"
		m.cond[i] = sync.NewCond(&m.mu)
	}
	return m
}

func (m *Monitor) TakeForks(i int) {
	fmt.Printf("[LOG] Filósofo %d intentando entrar a SECCIÓN CRÍTICA (TakeForks)...\n", i)
	m.mu.Lock()
	fmt.Printf("[LOG] Filósofo %d ENTRÓ a SECCIÓN CRÍTICA (Mutex adquirido)\n", i)

	m.state[i] = "HUNGRY"
	m.test(i)

	// Si no pudo pasar a EATING en test(), espera a que un vecino lo despierte
	for m.state[i] != "EATING" {
		fmt.Printf("[LOG] Filósofo %d no puede comer. Libera Mutex y espera (Wait)...\n", i)
		m.cond[i].Wait()
		fmt.Printf("[LOG] Filósofo %d despertó y RECUPERÓ Mutex.\n", i)
	}

	fmt.Printf("[LOG] Filósofo %d SALIENDO de SECCIÓN CRÍTICA (TakeForks)...\n", i)
	m.mu.Unlock()
}

func (m *Monitor) PutForks(i int) {
	fmt.Printf("[LOG] Filósofo %d intentando entrar a SECCIÓN CRÍTICA (PutForks)...\n", i)
	m.mu.Lock()
	fmt.Printf("[LOG] Filósofo %d ENTRÓ a SECCIÓN CRÍTICA (Mutex adquirido)\n", i)

	m.state[i] = "THINKING"

	// Al terminar, verifica si los vecinos estaban esperando para comer
	m.test((i + m.n - 1) % m.n) // Vecino izquierdo
	m.test((i + 1) % m.n)       // Vecino derecho

	fmt.Printf("[LOG] Filósofo %d SALIENDO de SECCIÓN CRÍTICA (PutForks)...\n", i)
	m.mu.Unlock()
}

func (m *Monitor) test(i int) {
	left := (i + m.n - 1) % m.n
	right := (i + 1) % m.n

	// Condición crítica: Solo come si tiene hambre Y los vecinos no están comiendo
	if m.state[i] == "HUNGRY" && m.state[left] != "EATING" && m.state[right] != "EATING" {
		m.state[i] = "EATING"
		m.cond[i].Signal() // Despierta al filósofo i si estaba en Wait()
	}
}

func (m *Monitor) GetStates() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Retornamos una copia para evitar condiciones de carrera al leer desde el web server
	copyState := make([]string, len(m.state))
	copy(copyState, m.state)
	return copyState
}
