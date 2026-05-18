type DowntimePoint = {
  date: string;
  hours: number;
};

type Props = {
  points: DowntimePoint[];
};

function heatColor(hours: number) {
  if (hours === 0) {
    return '#0f172a';
  }
  if (hours < 12) {
    return '#22d3ee';
  }
  if (hours < 72) {
    return '#38bdf8';
  }
  if (hours < 150) {
    return '#818cf8';
  }
  return '#f43f5e';
}

export function MaintenanceHeatmap({ points }: Props) {
  const monthlyTotals = new Map<number, number>();

  for (const point of points) {
    const month = new Date(point.date).getUTCMonth();
    monthlyTotals.set(month, (monthlyTotals.get(month) || 0) + point.hours);
  }

  return (
    <section className="panel">
      <div className="panel-header">
        <h2>Maintenance Heatmap</h2>
      </div>
      <div className="heatmap-grid">
        {Array.from({ length: 12 }).map((_, month) => {
          const total = monthlyTotals.get(month) || 0;
          return (
            <div key={month} className="heat-cell" style={{ backgroundColor: heatColor(total) }}>
              <span className="heat-month">{new Date(Date.UTC(2025, month, 1)).toLocaleString(undefined, { month: 'short' })}</span>
              <strong>{total}h</strong>
            </div>
          );
        })}
      </div>
    </section>
  );
}
