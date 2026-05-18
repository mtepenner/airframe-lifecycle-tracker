CREATE TABLE IF NOT EXISTS airframes (
    id BIGSERIAL PRIMARY KEY,
    tail_number TEXT NOT NULL UNIQUE,
    model TEXT NOT NULL,
    manufacturer TEXT NOT NULL,
    first_flight_date DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO airframes (tail_number, model, manufacturer, first_flight_date)
VALUES
    ('N711MD', 'MD-11', 'McDonnell Douglas', '1992-03-14'),
    ('N737AL', 'Boeing 737', 'Boeing', '2003-09-10'),
    ('N8DC', 'DC-8', 'Douglas', '1972-06-21')
ON CONFLICT (tail_number) DO NOTHING;
