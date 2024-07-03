--liquibase formatted sql

--changeset suraj.bobade:1 context:dev and qa and rel
-- give view(s) ownership to special privileges user
-- ALTER VIEW time_entry.v_time_log OWNER TO ${DB_USER};
