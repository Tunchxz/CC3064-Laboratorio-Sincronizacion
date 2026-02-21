# 🐳 Guía de Docker

Esta guía explica cómo usar Docker y Docker Compose para ejecutar el backend de Dining Philosophers.

## 📋 Prerequisitos

- Docker instalado ([Instalar Docker](https://docs.docker.com/get-docker/))
- Docker Compose instalado (viene incluido con Docker Desktop)

## 🏗️ Arquitectura del Dockerfile

El Dockerfile usa **multi-stage build** para optimizar el tamaño de la imagen:

### Stage 1: Builder
- Imagen base: `golang:1.21-alpine`
- Compila el código Go
- Genera el binario `dining_philosophers`

### Stage 2: Runtime
- Imagen base: `alpine:latest` (muy ligera, ~5MB)
- Solo contiene el binario compilado
- Ejecuta con usuario no-root (seguridad)
- Imagen final: ~15-20MB

**Ventajas**:
- ✅ Imagen final pequeña (no incluye compilador)
- ✅ Más rápido de desplegar
- ✅ Más seguro (menos superficie de ataque)
- ✅ Usuario no-root

## 🚀 Comandos Docker

### Construcción

```bash
# Construcción básica
docker build -t dining-philosophers .

# Construcción sin caché (desde cero)
docker build --no-cache -t dining-philosophers .

# Ver imágenes creadas
docker images | grep dining-philosophers
```

### Ejecución

```bash
# Ejecutar en primer plano (ver logs en tiempo real)
docker run -p 8080:8080 dining-philosophers

# Ejecutar en background (daemon)
docker run -d -p 8080:8080 --name dp-backend dining-philosophers

# Ejecutar con auto-eliminación al detener
docker run -it --rm -p 8080:8080 dining-philosophers

# Ejecutar en puerto diferente (ej: 3000)
docker run -p 3000:8080 dining-philosophers
```

### Gestión de contenedores

```bash
# Listar contenedores en ejecución
docker ps

# Listar todos los contenedores (incluidos detenidos)
docker ps -a

# Ver logs de un contenedor
docker logs dp-backend
docker logs -f dp-backend  # seguir logs en tiempo real

# Detener contenedor
docker stop dp-backend

# Reiniciar contenedor
docker restart dp-backend

# Eliminar contenedor
docker rm dp-backend

# Detener y eliminar
docker rm -f dp-backend
```

### Limpieza

```bash
# Eliminar imagen
docker rmi dining-philosophers

# Eliminar imágenes no usadas
docker image prune

# Eliminar todo (contenedores, imágenes, volúmenes)
docker system prune -a
```

## 🚀 Docker Compose

Docker Compose simplifica la gestión con un solo comando.

### Comandos básicos

```bash
# Iniciar servicio
docker-compose up

# Iniciar en background
docker-compose up -d

# Reconstruir imagen y iniciar
docker-compose up --build

# Ver logs
docker-compose logs
docker-compose logs -f  # seguir logs

# Ver estado
docker-compose ps

# Detener servicio (mantiene contenedor)
docker-compose stop

# Detener y eliminar contenedor
docker-compose down

# Detener, eliminar contenedor y volúmenes
docker-compose down -v
```

### Healthcheck

El `docker-compose.yml` incluye un healthcheck que verifica cada 30 segundos si el servidor responde:

```bash
# Ver estado de salud
docker-compose ps

# Output esperado:
# NAME                        STATUS
# dining-philosophers-backend Up (healthy)
```

## 🔧 Personalización

### Cambiar puerto

**Opción 1: En docker run**
```bash
docker run -p 9090:8080 dining-philosophers
# Ahora accesible en http://localhost:9090
```

**Opción 2: Modificar docker-compose.yml**
```yaml
ports:
  - "9090:8080"  # host:container
```

### Variables de entorno

Puedes agregar variables de entorno en `docker-compose.yml`:

```yaml
environment:
  - LOG_LEVEL=debug
  - MAX_PHILOSOPHERS=50
```

O al ejecutar con docker:
```bash
docker run -e LOG_LEVEL=debug -p 8080:8080 dining-philosophers
```

## 🐛 Troubleshooting

### "Address already in use"
El puerto 8080 está ocupado. Soluciones:

```bash
# Opción 1: Usar puerto diferente
docker run -p 9090:8080 dining-philosophers

# Opción 2: Encontrar qué usa el puerto
lsof -i :8080  # Linux/Mac
netstat -ano | findstr :8080  # Windows

# Opción 3: Detener proceso que usa el puerto
```

### "Cannot connect to Docker daemon"
Docker no está corriendo:

```bash
# Linux
sudo systemctl start docker

# Mac/Windows
# Abrir Docker Desktop
```

### Contenedor se detiene inmediatamente
Ver los logs para diagnosticar:

```bash
docker logs dp-backend
```

### Rebuild forzado
Si cambias código y no se refleja:

```bash
# Con docker
docker build --no-cache -t dining-philosophers .

# Con docker-compose
docker-compose build --no-cache
docker-compose up
```

## 📊 Monitoreo

### Ver uso de recursos

```bash
# CPU, memoria, red en tiempo real
docker stats

# Solo el contenedor específico
docker stats dp-backend
```

### Inspeccionar contenedor

```bash
# Ver configuración completa
docker inspect dp-backend

# Ver solo IP
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dp-backend
```

## 🔒 Seguridad

El Dockerfile implementa buenas prácticas de seguridad:

✅ **Multi-stage build**: Solo el binario en la imagen final
✅ **Usuario no-root**: Ejecuta como `appuser` (UID 1000)
✅ **Imagen minimal**: Alpine Linux (superficie de ataque reducida)
✅ **Sin dependencias innecesarias**: Solo el binario y certificados CA

## 🚀 Despliegue

### Subir a Docker Hub

```bash
# Tag de la imagen
docker tag dining-philosophers tu-usuario/dining-philosophers:latest

# Login
docker login

# Push
docker push tu-usuario/dining-philosophers:latest
```

### Usar en servidor

```bash
# Pull de la imagen
docker pull tu-usuario/dining-philosophers:latest

# Ejecutar
docker run -d -p 8080:8080 --restart=always tu-usuario/dining-philosophers:latest
```

## 📚 Recursos

- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Best practices for writing Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [Multi-stage builds](https://docs.docker.com/build/building/multi-stage/)
