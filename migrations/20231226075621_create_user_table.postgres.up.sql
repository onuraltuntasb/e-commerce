create table users (
	id serial primary key,
	first_name varchar(255) not null,
	last_name varchar(255) not null,
	password varchar(255) not null,
	email varchar(255) unique not null,
	status int not null,
	is_seller bool not null,
	auth_type int not null,
	address_id int,
	created_at timestamp not null,
	updated_at timestamp not null
);