create table district (
    id serial primary key,
    city_id int,
    name varchar(127) not null
);