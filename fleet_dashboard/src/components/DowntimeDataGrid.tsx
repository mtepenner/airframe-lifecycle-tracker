type Row = {
  model: string;
  operator: string;
  year: number;
  totalDowntimeHours: number;
  maintenanceEvents: number;
};

type Props = {
  rows: Row[];
  loading: boolean;
};

export function DowntimeDataGrid({ rows, loading }: Props) {
  return (
    <section className="panel">
      <div className="panel-header">
        <h2>Downtime Data Grid</h2>
      </div>

      {loading ? <p>Loading matrix...</p> : null}
      {!loading && rows.length === 0 ? <p>No rows match current filters.</p> : null}

      {rows.length > 0 ? (
        <div className="grid-wrap">
          <table>
            <thead>
              <tr>
                <th>Model</th>
                <th>Operator</th>
                <th>Year</th>
                <th>Downtime Hours</th>
                <th>Maintenance Events</th>
              </tr>
            </thead>
            <tbody>
              {rows.map((row, index) => (
                <tr key={`${row.model}-${row.operator}-${row.year}-${index}`}>
                  <td>{row.model}</td>
                  <td>{row.operator}</td>
                  <td>{row.year}</td>
                  <td>{row.totalDowntimeHours}</td>
                  <td>{row.maintenanceEvents}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : null}
    </section>
  );
}
