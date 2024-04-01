alter table product_category 
add constraint fk_product_category foreign key (product_id) references product (id);