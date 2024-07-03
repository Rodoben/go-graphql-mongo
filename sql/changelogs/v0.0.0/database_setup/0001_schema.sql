--liquibase formatted sql

--changeset ronald.benjamin:1context:dev and qa and rel
CREATE SCHEMA IF NOT EXISTS sales_order;
