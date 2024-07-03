-- SET UP FUCNTION FOR UPDATE sales_order_no coloumn


CREATE OR REPLACE FUNCTION sales_order_no_update()
  RETURNS trigger AS
$BODY$
BEGIN
    new.sales_order_no := (SELECT (case when max(sales_order_no) IS NULL then 1 else max(sales_order_no)+1 end) as unique_item_id  FROM sales_order.sales_order WHERE partner_id = new.partner_id);
    RETURN new;
END;
$BODY$
LANGUAGE plpgsql; 


-- SET up trigger for calling function


CREATE TRIGGER sales_order_no_insert
  BEFORE INSERT
  ON sales_order.sales_order
  FOR EACH ROW
  EXECUTE PROCEDURE sales_order_no_update();