import { useEffect, useState } from 'react';
import './App.css';
import PhilosophersCircle from './components/PhilosophersCircle';
import SimulationControls from './components/SimulationControls';
import Stats from './components/Stats';

const API_URL = 'http://localhost:8080';

function App() {
  const [philosophers, setPhilosophers] = useState([]);
  const [isRunning, setIsRunning] = useState(false);
  const [stats, setStats] = useState(null);
  const [config, setConfig] = useState({
    n: 5,
    duration: 60
  });

  // Polling para obtener el estado actual
  useEffect(() => {
    if (!isRunning) return;

    const interval = setInterval(async () => {
      try {
        const response = await fetch(`${API_URL}/status`);
        const data = await response.json();
        
        if (data.states) {
          const philosophersData = data.states.map((state, index) => ({
            id: index,
            state: state,
            eatCount: data.eatCount[index],
            waitTime: formatDuration(data.waitTime[index]),
            priority: data.priorities[index]
          }));
          setPhilosophers(philosophersData);
          setStats(data);
        }
      } catch (error) {
        console.error('Error fetching status:', error);
      }
    }, 500); // Actualizar cada 500ms

    return () => clearInterval(interval);
  }, [isRunning]);

  // Formatear duración de nanosegundos a segundos
  const formatDuration = (ns) => {
    return (ns / 1000000000).toFixed(2) + 's';
  };

  const startSimulation = async () => {
    try {
      const response = await fetch(
        `${API_URL}/start?n=${config.n}&duration=${config.duration}`,
        { method: 'POST' }
      );
      
      if (!response.ok) {
        const data = await response.json().catch(() => ({ message: 'Unknown error' }));
        alert(`Error: ${data.message || response.statusText}`);
        return;
      }
      
      const data = await response.json();
      setIsRunning(true);
      
      // Inicializar filósofos
      const initialPhilosophers = Array.from({ length: config.n }, (_, i) => ({
        id: i,
        state: 'THINKING',
        eatCount: 0,
        waitTime: '0.00s',
        priority: 0
      }));
      setPhilosophers(initialPhilosophers);
    } catch (error) {
      console.error('Error starting simulation:', error);
      const errorMessage = error.message.includes('fetch') 
        ? 'No se puede conectar con el backend. Verifica:\n\n' +
          '1. El backend está corriendo en http://localhost:8080\n' +
          '2. No hay problemas de firewall\n' +
          '3. El puerto 8080 no está siendo usado por otra aplicación\n\n' +
          'Si usas Docker, ejecuta: docker-compose up'
        : `Error: ${error.message}`;
      alert(errorMessage);
    }
  };

  const stopSimulation = async () => {
    try {
      await fetch(`${API_URL}/stop`, { method: 'POST' });
      setIsRunning(false);
    } catch (error) {
      console.error('Error stopping simulation:', error);
    }
  };

  return (
    <div className="app">
      <header className="app-header">
        <h1>🍴 Dining Philosophers Monitor</h1>
        <p className="subtitle">Implementación con Monitor y Sistema de Prioridades</p>
      </header>

      <div className="app-content">
        <div className="left-panel">
          <SimulationControls
            isRunning={isRunning}
            config={config}
            setConfig={setConfig}
            onStart={startSimulation}
            onStop={stopSimulation}
          />
          <Stats philosophers={philosophers} stats={stats} />
        </div>

        <div className="main-panel">
          <PhilosophersCircle philosophers={philosophers} />
        </div>
      </div>

      <footer className="app-footer">
        <p>CC3064 - Equipo 10 - Web Backend Concurrente</p>
      </footer>
    </div>
  );
}

export default App;

