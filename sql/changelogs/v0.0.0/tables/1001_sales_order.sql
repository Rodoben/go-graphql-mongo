--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
-- create salesorder table
CREATE TABLE IF NOT EXISTS sales_order.sales_order(
    id SERIAL NOT NULL,
    partner_id VARCHAR(128) NOT NULL,
    sales_order_id UUID NOT NULL  DEFAULT uuid_generate_v4() PRIMARY KEY,
    sales_order_no integer NOT NULL,
    company_id UUID NOT NULL,
    sales_order_status_id UUID NOT NULL,
    sales_rep_id integer NOT NULL,
    user_id VARCHAR(128) NOT NULL,
    order_date VARCHAR(50) ,
    due_date VARCHAR(50) ,
    sub_total numeric (18,2) ,
    sales_tax numeric (18,2) ,
    total numeric (18,2) ,
    opportunity_id integer,
    customer_po VARCHAR(50) ,
    location_id integer NOT NULL,
    department_id integer NOT NULL,
    deleted boolean NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    created_by VARCHAR(128),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    updated_by VARCHAR(128) NOT NULL
);

COMMENT ON TABLE sales_order.sales_order IS 'timer can be associated with any task/ticket/etc';
COMMENT ON COLUMN sales_order.sales_order.sales_order_id IS 'sales_order_no represents UUID of entity for which timer is created';


