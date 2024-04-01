alter table address 
add constraint fk_address_district foreign key (district_id) references district (id);