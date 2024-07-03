--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
-- create salesorder address table

CREATE TABLE IF NOT EXISTS sales_order.sales_order_shipping_address(
    id SERIAL NOT NULL,    
    partner_id VARCHAR(128) NOT NULL,
    sales_order_id UUID NOT NULL PRIMARY KEY,
    company_id UUID NOT NULL,
    contact_id UUID,
    address_id UUID,
    site_id UUID  NOT NULL,
    deleted boolean NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    created_by VARCHAR(128) NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_by VARCHAR(128) NOT NULL
);
