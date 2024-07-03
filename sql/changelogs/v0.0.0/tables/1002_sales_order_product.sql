
--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
-- create salesorder product table

CREATE TABLE IF NOT EXISTS sales_order.sales_order_product(
    id SERIAL NOT NULL,
    partner_id VARCHAR(128) NOT NULL,
    sales_order_id UUID NOT NULL PRIMARY KEY,
    product_id UUID NOT NULL,
    quantity integer NOT NULL,
    unit_price numeric (18,2) NOT NULL,
    bill_status boolean NOT NULL,
    product_class TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    created_by VARCHAR(128),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_by VARCHAR(128) NOT NULL
);