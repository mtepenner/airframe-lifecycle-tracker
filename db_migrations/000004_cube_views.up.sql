CREATE MATERIALIZED VIEW IF NOT EXISTS mv_fleet_cube AS
SELECT
    COALESCE(model, 'ALL_MODELS') AS model,
    COALESCE(operator, 'ALL_OPERATORS') AS operator,
    year,
    SUM(total_downtime_hours)::int AS total_downtime_hours,
    SUM(maintenance_events)::int AS maintenance_events
FROM (
    SELECT
        ml.model,
        ac.operator,
        EXTRACT(YEAR FROM ml.performed_at)::int AS year,
        SUM(ml.downtime_hours)::int AS total_downtime_hours,
        COUNT(*)::int AS maintenance_events
    FROM maintenance_logs ml
    LEFT JOIN LATERAL (
        SELECT operator
        FROM airframe_configurations c
        WHERE c.tail_number = ml.tail_number
          AND ml.performed_at >= c.valid_from
          AND ml.performed_at < c.valid_to
        ORDER BY c.valid_from DESC
        LIMIT 1
    ) ac ON TRUE
    GROUP BY CUBE (ml.model, ac.operator, EXTRACT(YEAR FROM ml.performed_at))
) cube_data
WHERE year IS NOT NULL
GROUP BY model, operator, year;

CREATE UNIQUE INDEX IF NOT EXISTS idx_mv_fleet_cube_unique
    ON mv_fleet_cube (model, operator, year);

CREATE OR REPLACE FUNCTION refresh_mv_fleet_cube() RETURNS void AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_fleet_cube;
END;
$$ LANGUAGE plpgsql;
