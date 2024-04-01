alter table product 
add constraint fk_product_users foreign key (user_id) references users (id);