create table product_comment (
id serial primary key unique not null,
description varchar(255) not null,
star int not null,
user_id int not null,
product_id int not null,
created_at timestamp not null,
updated_at timestamp not null);