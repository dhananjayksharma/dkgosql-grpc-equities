-- Table: merchants

-- DROP TABLE merchants;

CREATE TABLE IF NOT EXISTS merchants
(
    id bigint NOT NULL,
    code text COLLATE pg_catalog."default",
    name text COLLATE pg_catalog."default",
    address text COLLATE pg_catalog."default",
    status smallint DEFAULT 1,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT merchants_pkey PRIMARY KEY (id)
);
 
-- Index: Code_UniqueIndex

-- DROP INDEX "Code_UniqueIndex";

CREATE INDEX if not exists "Code_UniqueIndex"
    ON merchants USING btree
    (code COLLATE pg_catalog."default" ASC NULLS LAST);


-- Table: orders

-- DROP TABLE orders;

CREATE TABLE IF NOT EXISTS orders
(
    id bigint NOT NULL,
    status smallint,
    user_id bigint,
    company_id bigint,
    order_id bigint,
    order_type integer,
    quantity bigint,
    created_dt timestamp with time zone,
    updated_dt timestamp with time zone,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
);
  
-- Table: ordersprocessed

-- DROP TABLE ordersprocessed;

CREATE TABLE IF NOT EXISTS ordersprocessed
(
    id bigint NOT NULL,
    status smallint,
    user_id bigint,
    company_id bigint,
    order_id bigint,
    order_type integer,
    quantity bigint,
    quantity_processed bigint,
    created_dt timestamp with time zone,
    updated_dt timestamp with time zone,
    CONSTRAINT ordersprocessed_pkey PRIMARY KEY (id)
);
 
-- Table: users

-- DROP TABLE users;

CREATE TABLE IF NOT EXISTS users
(
    id bigint NOT NULL,
    fk_code text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    first_name text COLLATE pg_catalog."default",
    last_name text COLLATE pg_catalog."default",
    mobile text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    is_active smallint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
 
-- Index: Code_Email_UniqueIndex

-- DROP INDEX "Code_Email_UniqueIndex";

CREATE INDEX if not exists "Code_Email_UniqueIndex"
    ON users USING btree
    (fk_code COLLATE pg_catalog."default" ASC NULLS LAST, email COLLATE pg_catalog."default" ASC NULLS LAST);