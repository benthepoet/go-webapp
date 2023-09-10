create table address (
  id integer primary key,
  address_line varchar(255),
  address_line2 varchar(255),
  city varchar(255),
  state_or_province varchar(255),
  postal_code varchar(255),
  country_code varchar(2)
)

create table contact (
  id integer primary key,
  first_name varchar(255) not null,
  last_name varchar(255) not null,
  email varchar(255) not null
)

create table company (
  id integer primary key,
  name varchar(255),
  physical_address_id integer,
  mailing_address_id integer
)

create table contact_company (
  id integer primary key,
  contact_id integer,
  company_id integer,
  is_default bool
)

create table tenant (
  id integer primary key,
  name varchar(255)
)

create table store (
  id integer primary key,
  tenant_id integer,
  domain_name varchar(255),
  title varchar(255),
  is_enabled bool default false
)

create table product (
  id integer primary key,
  store_id integer,
  name varchar(255)
)

create table product_variant (
  id integer primary key,
  product_id integer,
  name varchar(255),
  price decimal(10, 2),
  is_enabled bool default false
)

create table store_login (
  id integer primary key,
  store_id integer,
  contact_id integer,
  email varchar(255),
  password_hash varchar(255),
  is_enabled bool default false
  is_archived bool default false
)

create table store_cart (
  id integer primary key,
  store_login_id integer
)

create table store_cart_item (
  id integer primary key,
  product_variant_id integer,
  quantity integer
)

create table sales_order_status (
  id integer primary key,
  name varchar(255)
)

create table sales_order (
  id integer primary key,
  store_login_id integer,
  order_status_id integer,
  shipping_address_id integer,
  billing_address_id integer,
  special_instructions varchar(512)
)

create table sales_order_item (
  id integer primary key,
  product_variant_id integer,
  price decimal(10, 2),
  quantity integer
)