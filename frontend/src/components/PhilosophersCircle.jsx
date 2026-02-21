import { useEffect, useState } from 'react';
import EAT_IMG from '../assets/EAT.png';
import HUNGRY_IMG from '../assets/HUNGRY.png';
import THINK_IMG from '../assets/THINK.png';

function PhilosophersCircle({ philosophers }) {
  const [positions, setPositions] = useState([]);

  useEffect(() => {
    if (philosophers && philosophers.length > 0) {
      const radius = 200; // Radio del círculo
      const centerX = 300; // Centro X del contenedor
      const centerY = 300; // Centro Y del contenedor
      
      const newPositions = philosophers.map((_, index) => {
        const angle = (index * 2 * Math.PI) / philosophers.length - Math.PI / 2;
        return {
          x: centerX + radius * Math.cos(angle),
          y: centerY + radius * Math.sin(angle),
          angle: angle
        };
      });
      
      setPositions(newPositions);
    }
  }, [philosophers]);

  const getStateImage = (state) => {
    switch (state) {
      case 'EATING':
        return EAT_IMG;
      case 'HUNGRY':
        return HUNGRY_IMG;
      case 'THINKING':
      default:
        return THINK_IMG;
    }
  };

  const getStateColor = (state) => {
    switch (state) {
      case 'EATING':
        return '#4CAF50'; // Verde
      case 'HUNGRY':
        return '#FFC107'; // Amarillo
      case 'THINKING':
      default:
        return '#2196F3'; // Azul
    }
  };

  if (!philosophers || philosophers.length === 0) {
    return (
      <div className="philosophers-circle-empty">
        <p>La simulación no está activa. Inicia una simulación para ver a los filósofos.</p>
      </div>
    );
  }

  return (
    <div className="philosophers-circle">
      <svg width="600" height="600" className="circle-svg">
        {/* Mesa circular en el centro */}
        <circle cx="300" cy="300" r="100" fill="#8B7355" stroke="#654321" strokeWidth="3" />
        <text x="300" y="305" textAnchor="middle" fill="white" fontSize="16" fontWeight="bold">
          MESA
        </text>
        
        {/* Líneas conectando a los filósofos con la mesa */}
        {positions.map((pos, index) => (
          <line 
            key={`line-${index}`}
            x1="300" 
            y1="300" 
            x2={pos.x} 
            y2={pos.y}
            stroke="#ccc"
            strokeWidth="1"
            strokeDasharray="5,5"
          />
        ))}
      </svg>

      {/* Filósofos con sus imágenes y datos */}
      {philosophers.map((philosopher, index) => {
        if (!positions[index]) return null;
        
        const pos = positions[index];
        
        return (
          <div
            key={philosopher.id}
            className="philosopher"
            style={{
              position: 'absolute',
              left: `${pos.x - 50}px`,
              top: `${pos.y - 50}px`,
              transition: 'all 0.3s ease'
            }}
          >
            <div 
              className={`philosopher-card state-${philosopher.state.toLowerCase()}`}
              style={{ borderColor: getStateColor(philosopher.state) }}
            >
              <img 
                src={getStateImage(philosopher.state)} 
                alt={philosopher.state}
                className="philosopher-image"
              />
              <div className="philosopher-info">
                <div className="philosopher-id">#{philosopher.id}</div>
                <div className="philosopher-state" style={{ color: getStateColor(philosopher.state) }}>
                  {philosopher.state}
                </div>
                <div className="philosopher-stats">
                  <div className="stat">🍽️ {philosopher.eatCount}</div>
                  <div className="stat">⏱️ {philosopher.priority}</div>
                </div>
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
}

export default PhilosophersCircle;
