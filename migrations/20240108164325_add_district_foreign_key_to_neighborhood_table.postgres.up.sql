alter table neighborhood
add constraint fk_neighborhood_district foreign key (district_id) references district (id);
