create table refresh_token (
id serial primary key unique not null,
refresh_token varchar(511) not null,
user_id int not null,
created_at timestamp not null,
updated_at timestamp not null
);