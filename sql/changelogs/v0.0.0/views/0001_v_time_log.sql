-- --liquibase formatted sql

-- --changeset suraj.bobade:1 context:dev and qa and rel splitStatements:false
-- -- Create time_entry.v_time_log view
-- DROP VIEW IF EXISTS time_entry.v_time_log ;

-- CREATE OR REPLACE VIEW time_entry.v_time_log AS
--   SELECT
--     logs.timer_id,
--     logs.assignee_id,
--     logs.is_active,
--     logs.start_time,
--     logs.end_time,
--     logs.deduction,
--     logs.actual_time,
--     logs.is_billable
--   FROM time_entry.time_log logs
--   INNER JOIN time_entry.timer t ON logs.timer_id = t.id;
