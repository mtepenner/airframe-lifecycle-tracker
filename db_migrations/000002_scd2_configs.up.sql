CREATE TABLE IF NOT EXISTS airframe_configurations (
    id BIGSERIAL PRIMARY KEY,
    tail_number TEXT NOT NULL REFERENCES airframes(tail_number) ON DELETE CASCADE,
    model TEXT NOT NULL,
    operator TEXT NOT NULL,
    livery TEXT NOT NULL,
    seating_capacity INTEGER NOT NULL CHECK (seating_capacity > 0),
    valid_from TIMESTAMPTZ NOT NULL,
    valid_to TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CHECK (valid_from < valid_to)
);

CREATE INDEX IF NOT EXISTS idx_airframe_configs_tail_valid
    ON airframe_configurations (tail_number, valid_from, valid_to);

CREATE INDEX IF NOT EXISTS idx_airframe_configs_operator
    ON airframe_configurations (operator);

INSERT INTO airframe_configurations (tail_number, model, operator, livery, seating_capacity, valid_from, valid_to)
VALUES
    ('N711MD', 'MD-11', 'Atlas Air', 'Atlas Legacy', 285, '2010-01-01', '2018-06-30'),
    ('N711MD', 'MD-11', 'Atlas Air', 'Atlas Modern', 293, '2018-07-01', '9999-12-31'),
    ('N737AL', 'Boeing 737', 'Alaska Airlines', 'Classic Blue', 162, '2012-01-01', '2021-12-31'),
    ('N737AL', 'Boeing 737', 'Alaska Airlines', 'West Coast Refresh', 170, '2022-01-01', '9999-12-31'),
    ('N8DC', 'DC-8', 'Charter One', 'Retro Cargo', 110, '2000-01-01', '2015-12-31'),
    ('N8DC', 'DC-8', 'Charter One', 'Retro Cargo Refit', 118, '2016-01-01', '9999-12-31')
ON CONFLICT DO NOTHING;
