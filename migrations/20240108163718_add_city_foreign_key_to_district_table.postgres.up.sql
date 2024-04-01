alter table district
add constraint fk_district_city foreign key (city_id) references city (id);