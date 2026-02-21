# CC3064 - Laboratorio de Sincronización: Filósofos Comensales

**Equipo 10**  
**Universidad del Valle de Guatemala**  
**2024**

Implementación del Problema de los Filósofos Comensales utilizando el patrón Monitor con sistema de prioridades para evitar inanición.

## 📋 Descripción

Este proyecto implementa la solución clásica al problema de los Filósofos Comensales usando:

- **Backend en Go**: Monitor con variables de condición y sistema de prioridades
- **Frontend en React**: Visualización en tiempo real de la simulación
- **Docker**: Containerización para despliegue sencillo

### Características Principales

✅ **Monitor Pattern**: Sincronización con `sync.Mutex` y `sync.Cond`  
✅ **Sistema de Prioridades**: Evita inanición usando fórmula `priority = eatCount[i] * 10000 - waitTime[i]/100`  
✅ **API REST**: Endpoints `/start`, `/stop`, `/status`  
✅ **Visualización Interactiva**: Interfaz web con disposición circular  
✅ **Docker**: Configuración multi-stage para producción  

## 🏗️ Arquitectura

```
.
├── backend/
│   ├── monitor/           # Implementación del Monitor
│   ├── simulation/        # Lógica de simulación
│   ├── handlers/          # HTTP handlers
│   └── main.go           # Entry point
├── frontend/
│   ├── src/
│   │   ├── components/    # Componentes React
│   │   ├── assets/        # Imágenes (EAT, HUNGRY, THINK)
│   │   ├── App.jsx        # Componente principal
│   │   └── App.css        # Estilos
│   └── package.json
├── docker-compose.yml     # Orquestación de servicios
├── Dockerfile            # Build backend
└── Makefile              # Comandos útiles

```

## 🚀 Inicio Rápido

### Opción 1: Docker Compose (Recomendado)

```bash
# Iniciar backend y frontend
docker-compose up --build

# Acceder a:
# Backend: http://localhost:8080
# Frontend: http://localhost:5173
```

### Opción 2: Desarrollo Local

#### Backend

```bash
# Instalar dependencias Go
go mod download

# Ejecutar backend
go run main.go
# o
make run

# Backend disponible en http://localhost:8080
```

#### Frontend

```bash
cd frontend

# Instalar dependencias
npm install

# Iniciar dev server
npm run dev

# Frontend disponible en http://localhost:5173
```

## 📊 API Endpoints

### POST /start
Inicia una simulación con parámetros.

**Query Parameters:**
- `n`: Número de filósofos (2-100)
- `duration`: Duración en segundos (60-3600)

**Ejemplo:**
```bash
curl -X POST "http://localhost:8080/start?n=5&duration=120"
```

**Respuesta:**
```json
{
  "message": "Simulation started with 5 philosophers for 120 seconds"
}
```

### POST /stop
Detiene la simulación actual.

```bash
curl -X POST "http://localhost:8080/stop"
```

### GET /status
Obtiene el estado actual de todos los filósofos.

```bash
curl http://localhost:8080/status
```

**Respuesta:**
```json
{
  "running": true,
  "n": 5,
  "states": ["EATING", "THINKING", "HUNGRY", "THINKING", "EATING"],
  "eatCount": [12, 8, 5, 10, 11],
  "waitTime": [1500000000, 800000000, 2200000000, 900000000, 1100000000],
  "priorities": [118500, 78000, 28000, 91000, 99000]
}
```

## 🎯 Monitor & Prioridades

### Implementación del Monitor

El Monitor encapsula:
- **Estado interno**: `state[]`, `eatCount[]`, `waitTime[]`
- **Mutex**: Control de exclusión mutua
- **Condition Variables**: `self[i]` para cada filósofo

### Sistema de Prioridades

```go
priority = eatCount[i] * 10000 - waitTime[i]/100
```

**Menor número = Mayor prioridad**

- Los filósofos que han comido menos tienen mayor prioridad
- El tiempo de espera también afecta la prioridad
- Garantiza progreso justo y evita inanición

### Transiciones de Estado

```
THINKING → HUNGRY → EATING → THINKING
```

## 🎨 Frontend

### Visualización Circular

Los filósofos se muestran alrededor de una mesa circular con:
- **Imágenes de estado**: PNG animados (THINKING, HUNGRY, EATING)
- **Colores dinámicos**: Azul (THINKING), Amarillo (HUNGRY), Verde (EATING)
- **Animaciones**: Pulse effects para cada estado
- **Estadísticas**: eatCount y priority por filósofo

### Panel de Estadísticas

- Distribución de estados en tiempo real
- Promedio de comidas y prioridades
- Identificación del filósofo con más comidas
- Filósofo con mayor prioridad actual

## 🧪 Testing

### Backend

```bash
# Ejecutar tests
make test

# Con coverage
make test-coverage
```

### Frontend

```bash
cd frontend
npm run test
```

## 🔧 Desarrollo

### Hot Reload

#### Backend
```bash
# Instalar air (opcional)
go install github.com/air-verse/air@latest

# Ejecutar con hot reload
air

# o usar make
make dev
```

#### Frontend
```bash
cd frontend
npm run dev
# El servidor de Vite tiene hot reload por defecto
```

### Linting

```bash
# Backend
make lint

# Frontend
cd frontend
npm run lint
```

## 📖 Documentación Adicional

- [Arquitectura del Backend](ARCHITECTURE.md)
- [Guía de Docker](DOCKER.md)
- [Frontend README](frontend/README.md)

## 🐳 Docker

### Build & Run

```bash
# Build imagen
docker build -t dining-philosophers .

# Run container
docker run -p 8080:8080 dining-philosophers

# Con docker-compose
docker-compose up -d
docker-compose logs -f
docker-compose down
```

### Makefile Commands

```bash
make build          # Compilar binary
make run           # Ejecutar localmente
make docker-build  # Build imagen Docker
make docker-run    # Run container
make test          # Ejecutar tests
make clean         # Limpiar artifacts
make help          # Ver todos los comandos
```

## 🧩 Requisitos del Lab

Este proyecto cumple con todos los requisitos del laboratorio:

1. ✅ **Exclusión Mutua**: Implementada con `sync.Mutex`
2. ✅ **Variables de Condición**: Usando `sync.Cond` nativas de Go
3. ✅ **Monitor Pattern**: Encapsulación completa del estado
4. ✅ **Prevención de Deadlock**: Lógica de `test()` asegura progreso
5. ✅ **Mitigación de Inanición**: Sistema de prioridades basado en teoría
6. ✅ **Sin librerías de alto nivel**: Solo primitivas básicas (`sync` package)
7. ✅ **Logs detallados**: Timestamps en microsegundos
8. ✅ **API RESTful**: Endpoints para control y monitoreo

## 👥 Equipo 10

- [Nombre] - [Carnet]
- [Nombre] - [Carnet]
- [Nombre] - [Carnet]

## 📄 License

Este proyecto es parte del curso CC3064 - Laboratorio de Sistemas Operativos.  
Universidad del Valle de Guatemala - 2024

## 🙏 Referencias

- Dijkstra, E. W. (1971). "Hierarchical ordering of sequential processes"
- Hoare, C. A. R. (1974). "Monitors: An operating system structuring concept"
- Material del curso CC3064 - Slides sobre Monitors
