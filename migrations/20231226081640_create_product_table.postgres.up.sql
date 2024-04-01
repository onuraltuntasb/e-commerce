create table product (
id serial primary key unique not null,
user_id int not null,
product jsonb not null,
created_at timestamp not null,
updated_at timestamp not null
);