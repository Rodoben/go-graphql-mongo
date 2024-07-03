--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
-- create salesorder company table
CREATE TABLE IF NOT EXISTS sales_order.sales_order_company(
    id SERIAL NOT NULL,
    company_id UUID NOT NULL,
    partner_id VARCHAR(128) NOT NULL,
    sales_order_id UUID NOT NULL PRIMARY KEY,
    contact_id UUID,
    address_id UUID,
    is_same_address boolean NOT NULL,
    site_id UUID,
    deleted boolean NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    created_by VARCHAR(128) NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_by VARCHAR(128) NOT NULL
);



