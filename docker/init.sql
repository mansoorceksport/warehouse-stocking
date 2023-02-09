create database stocking;
GRANT ALL PRIVILEGES ON DATABASE stocking TO postgres;

\connect stocking;
create extension if not exists "uuid-ossp";


create table if not exists item(
    item_id serial,
    item_code uuid default uuid_generate_v4(),
    name varchar(50),
    price double precision,
    primary key (item_id)
);


INSERT INTO item(name, price)
values
    ('apple', 0.99),
    ('orange', 0.95),
    ('Dragon Fruite', 0.90);


create table if not exists depot(
    depot_id serial,
    depot_code uuid default uuid_generate_v4(),
    name varchar(50),
    primary key (depot_id)
);

create table if not exists store(
    store_id serial,
    store_code uuid default uuid_generate_v4(),
    name varchar(50),
    primary key (store_id)
);

insert into depot(name) values ('depot cikarang Pertama');

create table if not exists warehouse_inventory(
    depot_id serial not null,
    item_id serial not null,
    quantity integer default 0
);

insert into warehouse_inventory(depot_id, item_id, quantity)
values
    (1,1,100),
    (1,2,100),
    (1,3,50);



create table if not exists store_inventory(
  store_id serial not null,
  item_id serial not null,
  quantity integer default 0
);

