create table address (
id serial primary key,
user_id serial not null,
header varchar(127) unique not null,
name varchar(127) not null,
country varchar(127) not null,
city_id serial not null,
district_id serial not null,
neighborhood_id serial not null,
description varchar(127) not null,
phone varchar(10) not null,
bill_type bool not null,
ssn varchar(11) not null,
is_primary bool not null,
created_at timestamp not null,
updated_at timestamp not null
);


