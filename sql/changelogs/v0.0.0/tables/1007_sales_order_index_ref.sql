--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel
-- create salesorder indexes and references key

CREATE INDEX "FK1" ON sales_order.sales_order ("sales_order_id");
CREATE INDEX "FK2" ON sales_order.sales_order_company ("sales_order_id");
CREATE INDEX "FK3" ON sales_order.sales_order_shipping_address ("sales_order_id");
CREATE INDEX "FK4" ON sales_order.sales_order_note ("sales_order_id");
CREATE INDEX "FK5" ON sales_order.sales_order_product ("sales_order_id");


ALTER TABLE sales_order.sales_order_note ADD CONSTRAINT "FK_sales_order_note.sales_order_id" FOREIGN KEY ("sales_order_id") REFERENCES sales_order.sales_order ("sales_order_id");
ALTER TABLE sales_order.sales_order_company ADD CONSTRAINT "FK_sales_order_company.sales_order_id" FOREIGN KEY ("sales_order_id") REFERENCES sales_order.sales_order ("sales_order_id");
ALTER TABLE sales_order.sales_order_shipping_address ADD CONSTRAINT "FK_sales_order_shipping_address.sales_order_id" FOREIGN KEY ("sales_order_id") REFERENCES sales_order.sales_order ("sales_order_id");
ALTER TABLE sales_order.sales_order_product ADD CONSTRAINT "FK_sales_order_product.sales_order_id" FOREIGN KEY ("sales_order_id") REFERENCES sales_order.sales_order ("sales_order_id");
