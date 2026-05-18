CREATE TABLE IF NOT EXISTS maintenance_logs (
    id BIGSERIAL PRIMARY KEY,
    tail_number TEXT NOT NULL REFERENCES airframes(tail_number) ON DELETE CASCADE,
    model TEXT NOT NULL,
    check_type TEXT NOT NULL,
    downtime_hours INTEGER NOT NULL CHECK (downtime_hours >= 0),
    performed_at TIMESTAMPTZ NOT NULL,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_maintenance_performed_at ON maintenance_logs (performed_at);
CREATE INDEX IF NOT EXISTS idx_maintenance_model ON maintenance_logs (model);

INSERT INTO maintenance_logs (tail_number, model, check_type, downtime_hours, performed_at, notes)
VALUES
    ('N711MD', 'MD-11', 'D-CHECK', 190, '2023-02-14T08:00:00Z', 'Full structural inspection'),
    ('N711MD', 'MD-11', 'A-CHECK', 8, '2024-03-08T09:00:00Z', 'Routine A-check'),
    ('N711MD', 'MD-11', 'C-CHECK', 72, '2025-01-15T10:30:00Z', 'Avionics update'),
    ('N737AL', 'Boeing 737', 'A-CHECK', 6, '2024-04-11T13:30:00Z', 'Scheduled inspection'),
    ('N737AL', 'Boeing 737', 'C-CHECK', 56, '2024-10-02T06:00:00Z', 'Hydraulics maintenance'),
    ('N737AL', 'Boeing 737', 'D-CHECK', 162, '2025-02-20T11:00:00Z', 'Major overhaul'),
    ('N8DC', 'DC-8', 'B-CHECK', 24, '2023-09-17T07:45:00Z', 'Engine borescope'),
    ('N8DC', 'DC-8', 'C-CHECK', 64, '2024-12-01T05:15:00Z', 'Cabin systems retrofit'),
    ('N8DC', 'DC-8', 'D-CHECK', 210, '2025-03-03T16:20:00Z', 'Aging aircraft structural program')
ON CONFLICT DO NOTHING;
