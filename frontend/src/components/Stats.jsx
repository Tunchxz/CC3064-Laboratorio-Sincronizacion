function Stats({ philosophers, stats }) {
  const calculateStateDistribution = () => {
    if (!philosophers || philosophers.length === 0) {
      return { THINKING: 0, HUNGRY: 0, EATING: 0 };
    }

    const distribution = philosophers.reduce((acc, phil) => {
      acc[phil.state] = (acc[phil.state] || 0) + 1;
      return acc;
    }, {});

    return {
      THINKING: distribution.THINKING || 0,
      HUNGRY: distribution.HUNGRY || 0,
      EATING: distribution.EATING || 0
    };
  };

  const calculateAverages = () => {
    if (!philosophers || philosophers.length === 0) {
      return {
        avgEatCount: 0,
        avgPriority: 0,
        totalPhilosophers: 0
      };
    }

    const totalEatCount = philosophers.reduce((sum, phil) => sum + phil.eatCount, 0);
    const totalPriority = philosophers.reduce((sum, phil) => sum + phil.priority, 0);

    return {
      avgEatCount: (totalEatCount / philosophers.length).toFixed(2),
      avgPriority: (totalPriority / philosophers.length).toFixed(0),
      totalPhilosophers: philosophers.length
    };
  };

  const getMaxEater = () => {
    if (!philosophers || philosophers.length === 0) return null;
    
    return philosophers.reduce((max, phil) => 
      phil.eatCount > max.eatCount ? phil : max
    , philosophers[0]);
  };

  const getHighestPriority = () => {
    if (!philosophers || philosophers.length === 0) return null;
    
    // Menor número = mayor prioridad
    return philosophers.reduce((highest, phil) => 
      phil.priority < highest.priority ? phil : highest
    , philosophers[0]);
  };

  const stateDistribution = calculateStateDistribution();
  const averages = calculateAverages();
  const maxEater = getMaxEater();
  const highestPriority = getHighestPriority();

  return (
    <div className="stats-panel">
      <h2>Estadísticas</h2>

      <div className="stats-section">
        <h3>Distribución de Estados</h3>
        <div className="stats-grid">
          <div className="stat-card thinking">
            <div className="stat-icon">T</div>
            <div className="stat-label">THINKING</div>
            <div className="stat-value">{stateDistribution.THINKING}</div>
          </div>
          <div className="stat-card hungry">
            <div className="stat-icon">H</div>
            <div className="stat-label">HUNGRY</div>
            <div className="stat-value">{stateDistribution.HUNGRY}</div>
          </div>
          <div className="stat-card eating">
            <div className="stat-icon">E</div>
            <div className="stat-label">EATING</div>
            <div className="stat-value">{stateDistribution.EATING}</div>
          </div>
        </div>
      </div>

      <div className="stats-section">
        <h3>Promedios</h3>
        <div className="stats-list">
          <div className="stat-item">
            <span className="stat-label">Total de Filósofos:</span>
            <span className="stat-value">{averages.totalPhilosophers}</span>
          </div>
          <div className="stat-item">
            <span className="stat-label">Promedio de Comidas:</span>
            <span className="stat-value">{averages.avgEatCount}</span>
          </div>
          <div className="stat-item">
            <span className="stat-label">Prioridad Promedio:</span>
            <span className="stat-value">{averages.avgPriority}</span>
          </div>
        </div>
      </div>

      {maxEater && (
        <div className="stats-section">
          <h3>Destacados</h3>
          <div className="stats-list">
            <div className="stat-item highlight">
              <span className="stat-label">Más Comidas:</span>
              <span className="stat-value">
                Filósofo #{maxEater.id} ({maxEater.eatCount} veces)
              </span>
            </div>
            {highestPriority && (
              <div className="stat-item highlight">
                <span className="stat-label">Mayor Prioridad:</span>
                <span className="stat-value">
                  Filósofo #{highestPriority.id} (prioridad: {highestPriority.priority})
                </span>
              </div>
            )}
          </div>
        </div>
      )}

      {stats && (
        <div className="stats-section">
          <h3>Tiempos</h3>
          <div className="stats-list">
            {stats.avgThinkTime && (
              <div className="stat-item">
                <span className="stat-label">Tiempo Promedio Pensando:</span>
                <span className="stat-value">{(stats.avgThinkTime / 1000000000).toFixed(2)}s</span>
              </div>
            )}
            {stats.avgEatTime && (
              <div className="stat-item">
                <span className="stat-label">Tiempo Promedio Comiendo:</span>
                <span className="stat-value">{(stats.avgEatTime / 1000000000).toFixed(2)}s</span>
              </div>
            )}
            {stats.avgWaitTime && (
              <div className="stat-item">
                <span className="stat-label">Tiempo Promedio Esperando:</span>
                <span className="stat-value">{(stats.avgWaitTime / 1000000000).toFixed(2)}s</span>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}

export default Stats;
