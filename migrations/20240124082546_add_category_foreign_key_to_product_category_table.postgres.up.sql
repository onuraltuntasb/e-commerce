alter table product_category 
add constraint fk_category_product foreign key (category_id) references category (id);