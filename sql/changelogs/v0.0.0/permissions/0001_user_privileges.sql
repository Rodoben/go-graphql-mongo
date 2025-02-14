--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
GRANT USAGE ON SCHEMA sales_order to ${DB_USER};
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA sales_order to ${DB_USER};
ALTER DEFAULT PRIVILEGES IN SCHEMA sales_order GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO ${DB_USER};
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA sales_order to ${DB_USER};
ALTER DEFAULT PRIVILEGES IN SCHEMA sales_order GRANT ALL PRIVILEGES ON SEQUENCES TO ${DB_USER};
