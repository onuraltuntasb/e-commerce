create table card (
id serial primary key unique not null,
product_id int not null,
user_id int not null,
created_at timestamp not null,
updated_at timestamp not null
);