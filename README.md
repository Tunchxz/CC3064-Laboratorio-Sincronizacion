# Laboratorio de Sincronización: Dinning Philosophers

Implementación del problema clásico de los Dinning Philosophers utilizando el patrón Monitor con sistema de prioridades para evitar inanición. El proyecto incluye un backend en Go que maneja la sincronización y lógica de simulación, junto con un frontend en React para visualización interactiva en tiempo real.

## Estructura del Proyecto

```
.
├── backend/                      # Backend en Go
│   ├── main.go                   # Punto de entrada del servidor
│   ├── go.mod                    # Dependencias Go
│   ├── Dockerfile                # Imagen Docker del backend
│   ├── docker-compose.yml        # Configuración Docker local
│   ├── handlers/
│   │   └── handlers.go           # Endpoints /start, /stop, /status
│   ├── middleware/
│   │   └── cors.go               # Configuración CORS
│   ├── monitor/
│   │   └── dining_monitor.go     # Implementación del patrón Monitor
│   └── simulation/
│       └── simulation.go         # Gestión de filósofos con goroutines
│
├── frontend/                               # Frontend en React
│   ├── src/
│   │   ├── App.jsx                         # Componente principal
│   │   ├── App.css                         # Estilos globales
│   │   ├── main.jsx
│   │   ├── components/
│   │   │   ├── PhilosopherssCircle.jsx     # Visualización circular
│   │   │   ├── SimulationControls.jsx      # Controles de simulación
│   │   │   └── Stats.jsx                   # Estadísticas en tiempo real
│   │   └── assets/
│   ├── package.json                        # Dependencias npm
│   ├── vite.config.js                      # Configuración Vite
│   ├── Dockerfile                          # Imagen Docker del frontend
│   └── nginx.conf
│
└── docker-compose.yml            # Orquestación de servicios completa
```

## Funcionalidades

### Backend (Go)

- **Monitor con Sistema de Prioridades**: Implementación del patrón Monitor que gestiona la adquisición y liberación de tenedores de forma segura, con un sistema de prioridades basado en el tiempo de espera y número de veces que ha comido cada filósofo.

- **API REST**:
  - `POST/GET /start?n=<num>&duration=<sec>`: Inicia simulación con N filósofos (2-100) y duración en segundos (60-1200)
  - `POST/GET /stop`: Detiene la simulación en ejecución
  - `GET /status`: Obtiene estado actual de todos los filósofos y métricas
  - `GET /health`: Health check del servicio

- **Goroutines**: Cada filósofo se ejecuta en su propia goroutine con ciclo pensar-tomar tenedores-comer-soltar tenedores

- **Logs Detallados**: Timestamps con microsegundos para análisis de concurrencia

#### Arquitectura de Sincronización

El sistema implementa el patrón Monitor para evitar condiciones de carrera y deadlocks:

- **Mutex Global**: Protege el estado compartido de todos los filósofos
- **Variables de Condición**: Una por filósofo para espera condicional
- **Sistema de Prioridades**: 
  - Prioridad = `eatCount[i] * 10000 - waitTime[i]/100`
  - Mayor prioridad = menor oportunidad de comer (previene inanición)
- **Test de Disponibilidad**: Verifica que ambos vecinos no estén comiendo
- **Wake-up Explícito**: Notifica a filósofos vecinos cuando se liberan tenedores

### Frontend (React + Vite)

- **Visualización Circular**: Disposición circular de filósofos con sus estados (THINKING, HUNGRY, EATING)

- **Controles Interactivos**: 
  - Selector de número de filósofos
  - Selector de duración
  - Botones Start/Stop con validaciones

- **Polling Automático**: Actualización del estado cada 500ms mediante peticiones al endpoint `/status`

- **Estadísticas en Tiempo Real**:
  - Estado de cada filósofo
  - Tiempo de espera
  - Número de veces que comió
  - Prioridad actual

- **Indicadores Visuales**: Colores y animaciones para representar estados

## Requisitos Previos

### Opción 1: Docker (Recomendado)
- Docker 20.10 o superior
- Docker Compose 2.0 o superior

### Opción 2: Desarrollo Local
- Go 1.23 o superior
- Node.js 20.x o superior
- npm 10.x o superior

## Instrucciones de Ejecución

### Opción 1: Docker Compose (Recomendado)

Esta es la forma más sencilla de ejecutar el proyecto completo:

```bash
# Clonar el repositorio
git clone https://github.com/Tunchxz/CC3064-Laboratorio-Sincronizacion
cd CC3064-Laboratorio-Sincronizacion

# Iniciar ambos servicios
docker-compose up --build

# La aplicación estará disponible en:
# - Frontend: http://localhost:5173
# - Backend API: http://localhost:8080
```

Para detener los servicios:

```bash
# Detener servicios
docker-compose down

# Detener y limpiar volúmenes
docker-compose down -v
```

### Opción 2: Desarrollo Local

#### Backend

```bash
# Navegar al directorio backend
cd backend

# Instalar dependencias
go mod download

# Ejecutar el servidor
go run main.go

# El backend estará disponible en http://localhost:8080
```

#### Frontend

```bash
# En otra terminal, navegar al directorio frontend
cd frontend

# Instalar dependencias
npm install

# Iniciar el servidor de desarrollo
npm run dev

# El frontend estará disponible en http://localhost:5173
```

### Opción 3: Docker Individual

#### Backend
```bash
cd backend
docker build -t philosopherss-backend .
docker run -p 8080:8080 philosopherss-backend
```

#### Frontend
```bash
cd frontend
docker build -t philosopherss-frontend .
docker run -p 5173:80 philosopherss-frontend
```

## Equipo 10

- Cristian Túnchez (231359)  
- Javier Linares (231135)  
- Estuardo Castro (23890)
