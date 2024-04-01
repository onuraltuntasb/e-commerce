create table neighborhood (
    id serial primary key,
    district_id int,
    postal_code varchar(5) not null,
    name varchar(127) not null
);