# Guía de Arquitectura Modular

## 🏗️ Decisiones de Diseño

### Separación de Responsabilidades

Cada paquete tiene una única responsabilidad clara:

1. **`monitor/`**: Lógica de sincronización pura
   - No conoce HTTP ni detalles de la simulación
   - Es reutilizable en otros contextos
   - Implementa solo el patrón Monitor

2. **`simulation/`**: Orquestación de filósofos
   - Usa el Monitor pero no implementa su lógica
   - Gestiona goroutines y ciclo de vida
   - Independiente del transporte (HTTP)

3. **`handlers/`**: Capa de presentación
   - Maneja HTTP y validaciones
   - Traduce entre web y dominio
   - No conoce detalles de sincronización

4. **`main.go`**: Punto de entrada
   - Solo configuración e inicialización
   - Conecta los componentes
   - Mínimo código (35 líneas)

## 📦 Diagrama de Dependencias

```
main.go
  └─> handlers/
       └─> simulation/
            └─> monitor/
```

**Flujo de dependencias unidireccional:**
- `monitor/` no depende de nadie (núcleo puro)
- `simulation/` solo depende de `monitor/`
- `handlers/` solo depende de `simulation/`
- `main.go` solo depende de `handlers/`

## 🔄 Flujo de Ejecución

### Inicio de Simulación
```
HTTP Request → handlers.StartHandler()
                  ↓
              Validar parámetros
                  ↓
              simulation.New()
                  ↓
              monitor.NewDiningMonitor()
                  ↓
              simulation.Start()
                  ↓
              N × goroutine philosopher()
                  ↓
              monitor.Pickup() / Putdown()
```

### Consulta de Estado
```
HTTP Request → handlers.StatusHandler()
                  ↓
              simulation.WriteStatus()
                  ↓
              monitor.Snapshot()
                  ↓
              JSON Response
```

## 🎯 Ventajas de Esta Arquitectura

### 1. Testabilidad
Cada paquete puede testearse independientemente:

```go
// Test del monitor sin HTTP ni simulación
func TestMonitor_Deadlock(t *testing.T) {
    m := monitor.NewDiningMonitor(5)
    // ...test lógica pura
}

// Test de la simulación sin HTTP
func TestSimulation_Start(t *testing.T) {
    s := simulation.New(3, 10*time.Second)
    // ...test orquestación
}

// Test de handlers sin lógica de negocio
func TestStartHandler_Validation(t *testing.T) {
    // ...test validaciones HTTP
}
```

### 2. Mantenibilidad
- Cambios en HTTP no afectan el Monitor
- Cambios en el algoritmo no afectan los handlers
- Fácil encontrar y modificar código específico

### 3. Reutilización
El Monitor puede usarse en otros contextos:

```go
// CLI tool
func main() {
    m := monitor.NewDiningMonitor(5)
    // uso directo sin HTTP
}

// gRPC service
func (s *Server) StartDining(ctx context.Context, req *pb.Request) {
    sim := simulation.New(req.N, req.Duration)
    // misma lógica, diferente transporte
}
```

### 4. Escalabilidad
Fácil agregar nuevos componentes:

```
handlers/
  ├── handlers.go          # Existente
  ├── metrics_handler.go   # Nuevo: exportar métricas Prometheus
  └── websocket_handler.go # Nuevo: updates en tiempo real
```

### 5. Comprensión
- 35 líneas en main.go (vs 360 originales)
- Cada archivo tiene ~100-200 líneas
- Nombres descriptivos y documentación clara

## 📚 Comparación: Monolítico vs Modular

### Antes (Monolítico)
```
main.go (360 líneas)
├── Constantes de estados
├── DiningMonitor (130 líneas)
├── Simulation (60 líneas)
├── HTTP Handlers (120 líneas)
├── Validaciones (40 líneas)
└── Main (10 líneas)
```

**Problemas:**
- ❌ Difícil navegar 360 líneas
- ❌ Todo acoplado
- ❌ Imposible testear partes individuales
- ❌ Cambios en HTTP afectan Monitor

### Después (Modular)
```
main.go (35 líneas)
monitor/ (190 líneas)
simulation/ (85 líneas)
handlers/ (165 líneas)
```

**Ventajas:**
- ✅ Archivos pequeños y enfocados
- ✅ Bajo acoplamiento
- ✅ Tests unitarios posibles
- ✅ Cambios aislados

## 🔍 Guía Rápida: ¿Dónde está qué?

### "¿Dónde está la lógica del Monitor?"
→ `monitor/dining_monitor.go`

### "¿Dónde están las validaciones?"
→ `handlers/handlers.go` (líneas 8-19 constantes, ~40-80 validaciones)

### "¿Dónde se crean las goroutines?"
→ `simulation/simulation.go` (función `Start()`)

### "¿Dónde se configuran los endpoints?"
→ `main.go` (líneas 15-17)

### "¿Dónde está el sistema de prioridades?"
→ `monitor/dining_monitor.go` (función `calculatePriority()`)

## 🎓 Para la Presentación

Si el profesor pregunta sobre la arquitectura:

> "Organizamos el código en una arquitectura modular de tres capas:
>
> 1. **Capa de Dominio** (`monitor/`): Implementación pura del Monitor según las slides, sin dependencias externas.
>
> 2. **Capa de Aplicación** (`simulation/`): Orquestación de los filósofos usando el Monitor.
>
> 3. **Capa de Presentación** (`handlers/`): API HTTP con validaciones.
>
> Esta separación facilita el testing, mantenimiento y cumple con principios SOLID de diseño de software."

## 📖 Referencias

- **Clean Architecture** (Robert C. Martin)
- **Domain-Driven Design** (Eric Evans)
- **Go Project Layout** (golang-standards/project-layout)
