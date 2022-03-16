-- Database sale_db

-- DROP TABLE public.orders;

CREATE TABLE public.orders (
	order_id serial NOT NULL,
	customer_name varchar(255) NOT NULL,
	ordered_at date NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY (order_id)
);

-- DROP TABLE public.items;
CREATE TABLE public.items (
	item_id serial NOT NULL,
	item_code varchar(6) NOT NULL,
	description varchar(255) NULL,
	quantity int NULL,
	order_id int4 NOT NULL,
	CONSTRAINT items_pk PRIMARY KEY (item_id),
  CONSTRAINT fk_orders_items FOREIGN KEY (order_id) REFERENCES public.orders(order_id) ON DELETE CASCADE ON UPDATE CASCADE
);

