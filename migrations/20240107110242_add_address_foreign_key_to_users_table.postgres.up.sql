alter table users 
add constraint fk_users_address foreign key (address_id) references address (id);