alter table address 
add constraint fk_address_neighborhood foreign key (neighborhood_id) references neighborhood (id);