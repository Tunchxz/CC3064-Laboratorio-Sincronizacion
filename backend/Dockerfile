# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Instalar dependencias de compilación
RUN apk add --no-cache git

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de módulos Go
COPY go.mod ./

# Descargar dependencias (si las hubiera)
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dining_philosophers .

# Stage 2: Runtime
FROM alpine:latest

# Instalar certificados CA para HTTPS (si se necesitan)
RUN apk --no-cache add ca-certificates

# Crear usuario no-root para seguridad
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /home/appuser

# Copiar binario desde el stage de build
COPY --from=builder /app/dining_philosophers .

# Cambiar permisos
RUN chown appuser:appuser dining_philosophers

# Cambiar a usuario no-root
USER appuser

# Exponer puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./dining_philosophers"]
