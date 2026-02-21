import { useState } from 'react';

const API_URL = 'http://localhost:8080';

function SimulationControls({ config, setConfig, isRunning, onStart, onStop }) {
  const [errors, setErrors] = useState({});

  const validateConfig = () => {
    const newErrors = {};
    
    if (config.n < 2 || config.n > 100) {
      newErrors.n = 'El número de filósofos debe estar entre 2 y 100';
    }
    
    if (config.duration < 60 || config.duration > 3600) {
      newErrors.duration = 'La duración debe estar entre 60 y 3600 segundos';
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleStart = () => {
    if (validateConfig()) {
      onStart();
    }
  };

  const handleNumberChange = (e) => {
    const value = parseInt(e.target.value) || 2;
    setConfig({ ...config, n: value });
    
    // Limpiar error cuando el usuario cambia el valor
    if (errors.n) {
      setErrors({ ...errors, n: null });
    }
  };

  const handleDurationChange = (e) => {
    const value = parseInt(e.target.value) || 60;
    setConfig({ ...config, duration: value });
    
    // Limpiar error cuando el usuario cambia el valor
    if (errors.duration) {
      setErrors({ ...errors, duration: null });
    }
  };

  return (
    <div className="simulation-controls">
      <h2>Configuración</h2>
      
      <div className="control-group">
        <label htmlFor="num-philosophers">
          Número de Filósofos:
        </label>
        <input
          id="num-philosophers"
          type="number"
          min="2"
          max="100"
          value={config.n}
          onChange={handleNumberChange}
          disabled={isRunning}
          className={errors.n ? 'error' : ''}
        />
        {errors.n && <span className="error-message">{errors.n}</span>}
        <small>Entre 2 y 100 filósofos</small>
      </div>

      <div className="control-group">
        <label htmlFor="duration">
          Duración (segundos):
        </label>
        <input
          id="duration"
          type="number"
          min="60"
          max="3600"
          value={config.duration}
          onChange={handleDurationChange}
          disabled={isRunning}
          className={errors.duration ? 'error' : ''}
        />
        {errors.duration && <span className="error-message">{errors.duration}</span>}
        <small>Entre 60 y 1200 segundos (20 minutos)</small>
      </div>

      <div className="control-buttons">
        <button
          onClick={handleStart}
          disabled={isRunning}
          className="btn btn-start"
        >
          {isRunning ? 'En Ejecución...' : 'Iniciar Simulación'}
        </button>
        
        <button
          onClick={onStop}
          disabled={!isRunning}
          className="btn btn-stop"
        >
          Detener Simulación
        </button>
      </div>

      <div className="info-box">
        <h3>Información</h3>
        <ul>
          <li><strong>THINKING:</strong> El filósofo está pensando</li>
          <li><strong>HUNGRY:</strong> El filósofo quiere comer</li>
          <li><strong>EATING:</strong> El filósofo está comiendo</li>
          <li><strong>Priority:</strong> Número de prioridad</li>
        </ul>
      </div>
    </div>
  );
}

export default SimulationControls;
