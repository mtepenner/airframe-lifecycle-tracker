import { TimelineScrubber } from './components/TimelineScrubber';
import { DowntimeDataGrid } from './components/DowntimeDataGrid';
import { MaintenanceHeatmap } from './components/MaintenanceHeatmap';
import { useFleetAnalytics } from './hooks/useFleetAnalytics';
import { useAirframeHistory } from './hooks/useAirframeHistory';
import { useDowntimeAnalytics } from './hooks/useDowntimeAnalytics';
import { useFilterStore } from './stores/useFilterStore';

export default function App() {
  const model = useFilterStore((state) => state.model);
  const setModel = useFilterStore((state) => state.setModel);
  const operator = useFilterStore((state) => state.operator);
  const setOperator = useFilterStore((state) => state.setOperator);
  const tailNumber = useFilterStore((state) => state.tailNumber);
  const setTailNumber = useFilterStore((state) => state.setTailNumber);

  const fleet = useFleetAnalytics();
  const history = useAirframeHistory();
  const downtime = useDowntimeAnalytics();

  return (
    <main className="page">
      <header className="hero">
        <h1>Airframe Lifecycle Tracker</h1>
        <p>Time-travel analytics for configuration history, downtime trends, and maintenance cadence.</p>
      </header>

      <section className="panel filters">
        <div className="panel-header">
          <h2>Global Filters</h2>
        </div>
        <div className="filter-grid">
          <label>
            Tail Number
            <input value={tailNumber} onChange={(event) => setTailNumber(event.target.value.toUpperCase())} />
          </label>
          <label>
            Model
            <input value={model} onChange={(event) => setModel(event.target.value)} placeholder="MD-11" />
          </label>
          <label>
            Operator
            <input value={operator} onChange={(event) => setOperator(event.target.value)} placeholder="Atlas Air" />
          </label>
        </div>
      </section>

      <TimelineScrubber minYear={2000} maxYear={2030} />

      <section className="panel">
        <div className="panel-header">
          <h2>Airframe SCD Timeline</h2>
        </div>
        {history.isLoading ? <p>Loading history...</p> : null}
        {history.error ? <p>History unavailable.</p> : null}
        {history.data?.length ? (
          <ul className="timeline-list">
            {history.data.map((entry, index) => (
              <li key={`${entry.tailNumber}-${entry.validFrom}-${index}`}>
                <strong>{entry.livery}</strong>
                <span>{entry.operator}</span>
                <span>{entry.seatingCapacity} seats</span>
                <span>{new Date(entry.validFrom).toLocaleDateString()} to {new Date(entry.validTo).toLocaleDateString()}</span>
              </li>
            ))}
          </ul>
        ) : null}
      </section>

      <DowntimeDataGrid rows={fleet.data?.rows || []} loading={fleet.isLoading} />

      <MaintenanceHeatmap points={downtime.data?.points || []} />

      <section className="panel stats">
        <h3>Fleet Snapshot</h3>
        <p>Total maintenance events: {fleet.data?.totalEvents || 0}</p>
        <p>Total downtime hours: {downtime.data?.totalHours || 0}</p>
      </section>
    </main>
  );
}
