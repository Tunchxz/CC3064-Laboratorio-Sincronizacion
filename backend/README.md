# Dining Philosophers Monitor - CC3064

Implementación del problema de los Dining Philosophers usando el patrón Monitor con sistema de prioridades.

## 📁 Arquitectura Modular

```
.
├── main.go                     # Punto de entrada de la aplicación
├── go.mod                      # Definición del módulo Go
├── Makefile                    # Comandos útiles de desarrollo
├── Dockerfile                  # Configuración Docker
├── docker-compose.yml          # Orquestación Docker
├── .dockerignore              # Archivos ignorados en build
├── README.md                   # Documentación principal
├── ARCHITECTURE.md             # Guía de arquitectura
├── DOCKER.md                   # Guía detallada de Docker
├── monitor/
│   └── dining_monitor.go      # Implementación del Monitor
├── simulation/
│   └── simulation.go          # Gestión de la simulación
└── handlers/
    └── handlers.go            # HTTP handlers y validaciones
```

## 🔧 Módulos

### 📦 `monitor/` - Núcleo del Monitor
Implementa el patrón Monitor según las Slides 1-5 de la clase:
- **DiningMonitor**: Estructura del Monitor con mutex y variables de condición
- **Sistema de prioridades**: Espera condicional con números de prioridad
- **Calendarización**: Gestión explícita de qué filósofo despertar
- **Métodos públicos**: `Pickup()` y `Putdown()`

### 📦 `simulation/` - Gestión de Simulación
Maneja la ejecución de la simulación:
- **Simulation**: Estructura que coordina N filósofos
- **Goroutines**: Un hilo por cada filósofo
- **Control de ciclo de vida**: Start(), Stop(), IsRunning()

### 📦 `handlers/` - API HTTP
Expone la funcionalidad vía HTTP REST:
- **StartHandler**: Iniciar simulación con validaciones
- **StopHandler**: Detener simulación en ejecución
- **StatusHandler**: Obtener estado actual
- **Validaciones**: Parámetros, métodos HTTP, concurrencia

## 🚀 Uso

### Opción 1: Directo con Go

#### Compilar
```bash
go build -o dining_philosophers
```

#### Ejecutar
```bash
./dining_philosophers
# O directamente:
go run main.go
```

### Opción 2: Con Makefile (recomendado para desarrollo) 🛠️

```bash
# Ver todos los comandos disponibles
make help

# Compilar
make build

# Ejecutar
make run

# Docker: build + run
make start

# Docker Compose
make compose-up

# Limpiar todo
make full-clean
```

### Opción 3: Con Docker 🐳

#### Construir imagen
```bash
docker build -t dining-philosophers .
```

#### Ejecutar contenedor
```bash
docker run -p 8080:8080 dining-philosophers
```

#### Ejecutar con logs visibles
```bash
docker run -it --rm -p 8080:8080 dining-philosophers
```

#### Detener contenedor
```bash
# Listar contenedores en ejecución
docker ps

# Detener contenedor específico
docker stop <container_id>
```

### Opción 4: Con Docker Compose 🚀

#### Iniciar servicio
```bash
docker-compose up
```

#### Iniciar en background
```bash
docker-compose up -d
```

#### Ver logs
```bash
docker-compose logs -f
```

#### Detener servicio
```bash
docker-compose down
```

> 📚 **Para más información sobre Docker** (troubleshooting, personalización, despliegue), consulta [DOCKER.md](DOCKER.md)

### Endpoints

#### Iniciar simulación
```bash
GET/POST /start?n=5&duration=60

Parámetros:
- n: Número de filósofos (2-100, default: 5)
- duration: Duración en segundos (60-3600, default: 60)

Ejemplo:
curl "http://localhost:8080/start?n=5&duration=120"
```

#### Detener simulación
```bash
GET/POST /stop

Ejemplo:
curl -X POST "http://localhost:8080/stop"
```

#### Obtener estado
```bash
GET /status

Ejemplo:
curl "http://localhost:8080/status"

Respuesta:
{
  "states": ["THINKING", "EATING", "HUNGRY", "THINKING", "EATING"],
  "waitTime": [123ms, 456ms, 789ms, 234ms, 567ms],
  "eatCount": [5, 3, 2, 4, 3],
  "priorities": [49990, 29980, 19970, 39985, 29975]
}
```

## 🎯 Características Implementadas

### ✅ Requisitos del Laboratorio
- ✅ Exclusión mutua con Monitor
- ✅ Prevención de deadlock (adquisición atómica)
- ✅ Prevención de starvation (sistema de prioridades)
- ✅ Logs de secciones críticas
- ✅ Backend concurrente (HTTP + goroutines)
- ✅ N filósofos configurable
- ✅ Duración ≥ 60 segundos

### ✅ Teoría de Clase (Slides 1-5)
- ✅ Monitor con mutex y variables de condición
- ✅ Espera condicional con números de prioridad
- ✅ Calendarización explícita de procesos
- ✅ Patrón similar a ResourceAllocator

### ✅ Mejoras de Ingeniería
- ✅ Arquitectura modular y escalable
- ✅ Validaciones exhaustivas
- ✅ Manejo de concurrencia seguro
- ✅ API REST con códigos HTTP apropiados
- ✅ Métricas de monitoreo

## 📊 Sistema de Prioridades

La prioridad se calcula como:

```
priority = (eatCount × 10000) - (waitTime / 100ms)
```

**Menor número = Mayor prioridad**

Ejemplo:
- Filósofo que ha comido 2 veces y esperado 3s: `20000 - 30 = 19970` ← Alta prioridad
- Filósofo que ha comido 5 veces y esperado 1s: `50000 - 10 = 49990` ← Baja prioridad

## 👥 Equipo
**Equipo 10 - CC3064**
- Lenguaje: Web (Go)
- Problema: Dining Philosophers
- Solución: Monitor con prioridades
