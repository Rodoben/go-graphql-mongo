--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
-- create extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
