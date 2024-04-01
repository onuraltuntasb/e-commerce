alter table address 
add constraint fk_address_city foreign key (city_id) references city (id);