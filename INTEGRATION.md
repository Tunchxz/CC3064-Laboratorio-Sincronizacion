# Guía de Integración Completa: Frontend + Backend

Esta guía explica cómo ejecutar el sistema completo de visualización de Filósofos Comensales.

## 📋 Tabla de Contenidos

1. [Arquitectura del Sistema](#arquitectura-del-sistema)
2. [Prerrequisitos](#prerrequisitos)
3. [Opción 1: Docker Compose (Recomendado)](#opción-1-docker-compose-recomendado)
4. [Opción 2: Desarrollo Local](#opción-2-desarrollo-local)
5. [Verificación](#verificación)
6. [Troubleshooting](#troubleshooting)

## 🏗️ Arquitectura del Sistema

```
┌─────────────┐         HTTP         ┌─────────────┐
│   Browser   │ <───────────────────> │   Frontend  │
│             │    http://localhost   │   (React)   │
└─────────────┘         :5173         └─────────────┘
                                             │
                                             │ REST API
                                             │ (Fetch)
                                             ▼
                                      ┌─────────────┐
                                      │   Backend   │
                                      │    (Go)     │
                                      │   :8080     │
                                      └─────────────┘
                                             │
                                             ▼
                                      ┌─────────────┐
                                      │   Monitor   │
                                      │   Pattern   │
                                      └─────────────┘
```

### Componentes

- **Frontend (React + Vite)**: 
  - Puerto: 5173 (dev) / 80 (production)
  - Polling cada 500ms al endpoint `/status`
  - Visualización circular con animaciones

- **Backend (Go + HTTP)**: 
  - Puerto: 8080
  - API REST con endpoints `/start`, `/stop`, `/status`
  - Monitor con sistema de prioridades

## 🔧 Prerrequisitos

### Para Docker Compose
- Docker 20.10+
- Docker Compose 2.0+

### Para Desarrollo Local
- Go 1.21+
- Node.js 18+ y npm
- Git

## 🚀 Opción 1: Docker Compose (Recomendado)

Esta es la forma más sencilla de ejecutar todo el sistema.

### 1.1 Build & Run

```bash
# Desde el directorio raíz del proyecto
docker-compose up --build
```

Este comando:
1. ✅ Construye la imagen del backend (Go multi-stage)
2. ✅ Construye la imagen del frontend (React + Nginx)
3. ✅ Inicia ambos servicios
4. ✅ Configura networking entre ellos
5. ✅ Espera a que el backend esté healthy antes de iniciar frontend

### 1.2 Acceder a los Servicios

- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080
- **Health Check**: http://localhost:8080/status

### 1.3 Ver Logs

```bash
# Logs de ambos servicios
docker-compose logs -f

# Solo backend
docker-compose logs -f backend

# Solo frontend
docker-compose logs -f frontend
```

### 1.4 Detener

```bash
# Detener servicios (mantiene imágenes)
docker-compose down

# Detener y eliminar imágenes
docker-compose down --rmi all

# Detener y eliminar todo (incluyendo volúmenes)
docker-compose down -v --rmi all
```

## 🛠️ Opción 2: Desarrollo Local

Para desarrollo activo con hot-reload.

### 2.1 Iniciar Backend

```bash
# Terminal 1
cd backend

# Instalar dependencias (primera vez)
go mod download

# Ejecutar backend
go run main.go

# ✅ Backend corriendo en http://localhost:8080
```

**Verificar:**
```bash
curl http://localhost:8080/status
```

### 2.2 Iniciar Frontend

```bash
# Terminal 2
cd frontend

# Instalar dependencias (primera vez)
npm install

# Iniciar dev server con hot reload
npm run dev

# ✅ Frontend corriendo en http://localhost:5173
```

**Abrir en navegador:**
```
http://localhost:5173
```

### 2.3 Desarrollo con Hot Reload

#### Backend (Opcional: Air)

```bash
# Instalar Air para hot reload
go install github.com/air-verse/air@latest

# Ejecutar con air
cd backend
air
```

#### Frontend

El servidor de Vite ya incluye hot reload por defecto. Cualquier cambio en los archivos `.jsx` o `.css` se reflejará automáticamente.

## ✅ Verificación

### 3.1 Verificar Backend

```bash
# Health check
curl http://localhost:8080/status

# Respuesta esperada:
# {"running":false,"n":0,"states":null,"eatCount":null,"waitTime":null,"priorities":null}
```

### 3.2 Iniciar Simulación desde CLI

```bash
# Iniciar simulación con 5 filósofos por 120 segundos
curl -X POST "http://localhost:8080/start?n=5&duration=120"

# Ver estado
curl http://localhost:8080/status

# Detener
curl -X POST "http://localhost:8080/stop"
```

### 3.3 Verificar Frontend

1. **Abrir navegador**: http://localhost:5173
2. **Verificar elementos**:
   - ✅ Panel izquierdo con controles
   - ✅ Panel central vacío (esperando simulación)
   - ✅ Campos de configuración (n y duration)
3. **Iniciar simulación**:
   - Configurar `n = 5` filósofos
   - Configurar `duration = 60` segundos
   - Clic en "Iniciar Simulación"
4. **Observar**:
   - ✅ Filósofos aparecen en círculo
   - ✅ Animaciones de estado (THINKING/HUNGRY/EATING)
   - ✅ Estadísticas se actualizan en tiempo real
   - ✅ Colores cambian según estado

## 🐛 Troubleshooting

### Error: "Failed to connect to backend"

**Causa**: Frontend no puede conectar al backend.

**Solución**:
```bash
# Verificar que el backend esté corriendo
curl http://localhost:8080/status

# Si no responde, iniciar backend
cd backend
go run main.go
```

### Error: "Port 8080 already in use"

**Causa**: Otro proceso está usando el puerto 8080.

**Solución**:
```bash
# Encontrar proceso
lsof -i :8080
# o en Linux
netstat -tulpn | grep 8080

# Matar proceso
kill -9 <PID>

# O cambiar puerto en backend
PORT=8081 go run main.go
```

### Error: "Port 5173 already in use"

**Causa**: Otro servidor Vite está corriendo.

**Solución**:
```bash
# Matar procesos de Node
pkill -f vite

# O cambiar puerto
cd frontend
npm run dev -- --port 3000
```

### Frontend muestra pero no se actualiza

**Causa**: Polling no está funcionando.

**Debug**:
1. Abrir DevTools (F12)
2. Ver consola para errores
3. Verificar Network tab - debe haber requests a `/status` cada 500ms
4. Verificar CORS si el backend está en diferente host

### Imágenes no se cargan en el frontend

**Causa**: Assets no encontrados.

**Verificar**:
```bash
cd frontend/src/assets
ls -la
# Debe mostrar: EAT.png, HUNGRY.png, THINK.png
```

### Docker Compose no encuentra el Dockerfile

**Causa**: Contextos incorrectos.

**Verificar estructura**:
```bash
.
├── backend/
│   └── Dockerfile
├── frontend/
│   └── Dockerfile
└── docker-compose.yml
```

## 🔄 Workflow de Desarrollo Típico

### Desarrollo Activo

```bash
# Terminal 1: Backend con hot reload
cd backend
air  # o go run main.go

# Terminal 2: Frontend con hot reload
cd frontend
npm run dev

# Terminal 3: Testing / curl requests
curl -X POST "http://localhost:8080/start?n=5&duration=60"
```

### Testing Rápido con Docker

```bash
# Build rápido y test
docker-compose up --build -d
docker-compose logs -f

# Cuando termine
docker-compose down
```

### Deploy de Producción

```bash
# Build optimizado
docker-compose build --no-cache

# Deploy
docker-compose up -d

# Verificar
docker-compose ps
docker-compose logs
```

## 📝 Comandos Útiles

```bash
# Ver todos los containers
docker ps -a

# Ver logs en tiempo real
docker-compose logs -f --tail=100

# Reconstruir solo un servicio
docker-compose up --build backend

# Entrar al container
docker-compose exec backend sh
docker-compose exec frontend sh

# Ver uso de recursos
docker stats

# Limpiar todo
docker system prune -a --volumes
```

## 🎯 Siguientes Pasos

Después de tener todo corriendo:

1. **Explorar la interfaz**: Prueba diferentes valores de `n` y `duration`
2. **Observar prioridades**: Mira cómo el sistema previene inanición
3. **Leer el código**: Revisa los componentes en `frontend/src/components/`
4. **Modificar y experimentar**: El hot reload facilita los cambios
5. **Monitorear logs**: Observa las transiciones de estado en tiempo real

## 📚 Referencias

- [Backend README](../backend/README.md)
- [Frontend README](../frontend/README.md)
- [Arquitectura](../ARCHITECTURE.md)
- [Docker Guide](../DOCKER.md)

---
**CC3064 - Equipo 10 - 2024**
