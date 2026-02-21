package monitor

import (
	"log"
	"sync"
	"time"
)

// Estados de los filósofos
const (
	THINKING = iota
	HUNGRY
	EATING
)

// DiningMonitor implementa el patrón Monitor para el problema de los Dining Philosophers
type DiningMonitor struct {
	n          int
	state      []int
	mutex      sync.Mutex
	cond       []*sync.Cond
	waitStart  []time.Time
	waitTime   []time.Duration
	// Sistema de prioridades para espera condicional
	eatCount []int // Contador de veces que ha comido cada filósofo
}

// NewDiningMonitor crea e inicializa un nuevo Monitor para n filósofos
func NewDiningMonitor(n int) *DiningMonitor {
	m := &DiningMonitor{
		n:         n,
		state:     make([]int, n),
		cond:      make([]*sync.Cond, n),
		waitStart: make([]time.Time, n),
		waitTime:  make([]time.Duration, n),
		eatCount:  make([]int, n),
	}

	for i := 0; i < n; i++ {
		m.state[i] = THINKING
		m.cond[i] = sync.NewCond(&m.mutex)
		m.eatCount[i] = 0
	}
	return m
}

// left retorna el índice del vecino izquierdo en topología circular
func (m *DiningMonitor) left(i int) int {
	return (i + m.n - 1) % m.n
}

// right retorna el índice del vecino derecho en topología circular
func (m *DiningMonitor) right(i int) int {
	return (i + 1) % m.n
}

// calculatePriority calcula el número de prioridad para espera condicional
// Menor número = mayor prioridad
// Prioridad basada en: 1) veces que ha comido (principal), 2) tiempo esperado (secundario)
func (m *DiningMonitor) calculatePriority(i int) int {
	// Factor principal: eatCount (quien menos ha comido tiene menor número)
	// Factor secundario: waitTime acumulado (quien más ha esperado tiene menor número)
	basePriority := m.eatCount[i] * 10000
	waitPenalty := int(m.waitTime[i].Milliseconds() / 100)
	return basePriority - waitPenalty // Resta waitPenalty para dar prioridad a quien más ha esperado
}

// shouldWakeUp verifica si el filósofo i tiene suficiente prioridad para ser despertado
// Implementa calendarización con prioridades
func (m *DiningMonitor) shouldWakeUp(i int) bool {
	if m.state[i] != HUNGRY {
		return false
	}

	// Puede comer solo si ambos vecinos no están comiendo
	if m.state[m.left(i)] == EATING || m.state[m.right(i)] == EATING {
		return false
	}

	// Verificar prioridad: solo despertar si tiene mayor prioridad que otros hambrientos que compiten
	myPriority := m.calculatePriority(i)

	// Comparar con vecinos hambrientos
	left := m.left(i)
	right := m.right(i)

	// Si el vecino izquierdo también está hambriento y puede comer,
	// verificar quién tiene mayor prioridad
	if m.state[left] == HUNGRY && m.state[m.left(left)] != EATING && m.state[i] != EATING {
		if m.calculatePriority(left) < myPriority {
			return false // El vecino izquierdo tiene mayor prioridad
		}
	}

	// Si el vecino derecho también está hambriento y puede comer,
	// verificar quién tiene mayor prioridad
	if m.state[right] == HUNGRY && m.state[i] != EATING && m.state[m.right(right)] != EATING {
		if m.calculatePriority(right) < myPriority {
			return false // El vecino derecho tiene mayor prioridad
		}
	}

	return true
}

// test evalúa si el filósofo i puede transicionar a EATING
// Implementación de calendarización con prioridades
func (m *DiningMonitor) test(i int) {
	if m.shouldWakeUp(i) {
		m.state[i] = EATING
		m.waitTime[i] += time.Since(m.waitStart[i])
		priority := m.calculatePriority(i)
		log.Printf("Philosopher %d granted access (priority=%d, eatCount=%d, waitTime=%.2fs)\n",
			i, priority, m.eatCount[i], m.waitTime[i].Seconds())
		m.cond[i].Signal()
	}
}

// Pickup es llamado cuando un filósofo quiere comer (entrar a sección crítica)
func (m *DiningMonitor) Pickup(i int) {
	m.mutex.Lock()

	m.state[i] = HUNGRY
	m.waitStart[i] = time.Now()
	log.Printf("Philosopher %d -> HUNGRY\n", i)

	m.test(i)

	for m.state[i] != EATING {
		m.cond[i].Wait()
	}

	log.Printf("Philosopher %d ENTER critical section (EATING)\n", i)

	m.mutex.Unlock()
}

// Putdown es llamado cuando un filósofo termina de comer (salir de sección crítica)
func (m *DiningMonitor) Putdown(i int) {
	m.mutex.Lock()

	m.state[i] = THINKING
	m.eatCount[i]++ // Incrementar contador para sistema de prioridades
	log.Printf("Philosopher %d EXIT critical section (THINKING) [total meals: %d]\n", i, m.eatCount[i])

	// Evaluar vecinos con sistema de prioridades
	m.test(m.left(i))
	m.test(m.right(i))

	m.mutex.Unlock()
}

// Snapshot retorna una copia del estado actual del Monitor de forma segura
func (m *DiningMonitor) Snapshot() map[string]interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	states := make([]string, m.n)
	priorities := make([]int, m.n)
	for i := 0; i < m.n; i++ {
		switch m.state[i] {
		case THINKING:
			states[i] = "THINKING"
		case HUNGRY:
			states[i] = "HUNGRY"
		case EATING:
			states[i] = "EATING"
		}
		priorities[i] = m.calculatePriority(i)
	}

	return map[string]interface{}{
		"states":     states,
		"waitTime":   m.waitTime,
		"eatCount":   m.eatCount,
		"priorities": priorities,
	}
}
