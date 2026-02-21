# Frontend - Dining Philosophers Visualization

Frontend React para visualizar la simulación del Problema de los Filósofos Comensales en tiempo real.

## 🎨 Características

- **Visualización Circular**: Disposición de filósofos alrededor de una mesa circular
- **Estados Animados**: Animaciones y colores para THINKING, HUNGRY, y EATING
- **Control de Simulación**: Configurar número de filósofos (2-100) y duración (60-3600s)
- **Estadísticas en Tiempo Real**: Distribución de estados, promedios, y destacados
- **Responsive**: Adaptable a diferentes tamaños de pantalla

## 🛠️ Tecnologías

- **React 19.2.0**: Framework de UI
- **Vite**: Build tool y dev server
- **CSS3**: Animaciones y gradientes
- **Fetch API**: Comunicación con backend

## 📁 Estructura

```
frontend/
├── src/
│   ├── assets/           # Imágenes de estados (EAT, HUNGRY, THINK)
│   ├── components/
│   │   ├── PhilosophersCircle.jsx   # Visualización circular
│   │   ├── SimulationControls.jsx   # Controles de inicio/parada
│   │   └── Stats.jsx                # Panel de estadísticas
│   ├── App.jsx           # Componente principal
│   ├── App.css           # Estilos globales
│   └── main.jsx          # Entry point
├── package.json
└── vite.config.js
```

## 🚀 Instalación y Ejecución

### Prerrequisitos

- Node.js 18+ y npm

### Opción 1: Desarrollo Local

1. **Instalar dependencias:**
   ```bash
   npm install
   ```

2. **Iniciar el servidor de desarrollo:**
   ```bash
   npm run dev
   ```

3. **Abrir en el navegador:**
   ```
   http://localhost:5173
   ```

### Opción 2: Build de Producción

1. **Construir para producción:**
   ```bash
   npm run build
   ```

2. **Preview del build:**
   ```bash
   npm run preview
   ```

## 🔌 Configuración del Backend

El frontend se conecta al backend Go en `http://localhost:8080`. Asegúrate de que el backend esté corriendo antes de iniciar la simulación.

Para iniciar el backend:
```bash
cd ..
docker-compose up
# o
make run
```

## 📊 Componentes

### PhilosophersCircle

Muestra los filósofos en disposición circular con:
- Imagen del estado actual (PNG)
- ID del filósofo
- Estado actual con color
- Número de comidas (eatCount)
- Prioridad actual

### SimulationControls

Panel de control con:
- Input para número de filósofos (2-100)
- Input para duración en segundos (60-3600)
- Botones de Iniciar/Detener
- Validación de parámetros
- Información de estados

### Stats

Panel de estadísticas que muestra:
- Distribución de estados (THINKING/HUNGRY/EATING)
- Promedios (comidas, prioridad)
- Filósofo con más comidas
- Filósofo con mayor prioridad
- Tiempos promedio (si disponible)

## 🎨 Estados Visuales

Cada filósofo tiene un color y animación según su estado:

- **THINKING** 💭: Azul (#2196F3) - Animación lenta
- **HUNGRY** 😋: Amarillo (#FFC107) - Animación rápida
- **EATING** 🍽️: Verde (#4CAF50) - Animación media

## 🔄 Polling

El frontend hace polling al endpoint `/status` cada 500ms cuando la simulación está activa para actualizar el estado en tiempo real.

## 📱 Responsive

El diseño se adapta a diferentes tamaños:
- **Desktop (>1200px)**: Layout con panel lateral
- **Tablet (768-1200px)**: Layout vertical
- **Mobile (<768px)**: Layout simplificado con elementos escalados

## 🐛 Troubleshooting

### "Failed to connect to backend"
- Verifica que el backend esté corriendo en puerto 8080
- Verifica que no haya problemas de CORS

### Las imágenes no se cargan
- Verifica que los archivos PNG estén en `src/assets/`
- Los nombres deben ser: `EAT.png`, `HUNGRY.png`, `THINK.png`

### Error de validación
- Número de filósofos debe estar entre 2 y 100
- Duración debe estar entre 60 y 3600 segundos

## 📄 License

Parte del Laboratorio de Sincronización - CC3064, Universidad del Valle de Guatemala
Equipo 10 - 2024

