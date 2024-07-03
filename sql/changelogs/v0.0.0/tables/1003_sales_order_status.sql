--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
-- create salesorder status table



CREATE TABLE IF NOT EXISTS sales_order.sales_order_status(
    id SERIAL NOT NULL,
    status_id UUID NOT NULL DEFAULT uuid_generate_v4(),
    partner_id VARCHAR(128) NOT NULL,
    status_name varchar(50) NOT NULL

);


/*  below query are to insert manual data to status table a per BSO-160 */


INSERT INTO sales_order.sales_order_status (partner_id, status_name) VALUES ('', 'New') returning id;
INSERT INTO sales_order.sales_order_status (partner_id, status_name) VALUES ('', 'In Progress') returning id;
INSERT INTO sales_order.sales_order_status (partner_id, status_name) VALUES ('', 'Approved') returning id;
INSERT INTO sales_order.sales_order_status (partner_id, status_name) VALUES ('', 'Waiting on Deposit') returning id;
INSERT INTO sales_order.sales_order_status (partner_id, status_name) VALUES ('', 'Pending Management Approval') returning id;
INSERT INTO sales_order.sales_order_status (partner_id, status_name) VALUES ('', 'Back Ordered') returning id;
INSERT INTO sales_order.sales_order_status (partner_id, status_name) VALUES ('', 'Complete') returning id;
