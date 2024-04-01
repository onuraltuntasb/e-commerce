alter table product_comment 
add constraint fk_product_comment_users foreign key (user_id) references users (id),
add constraint fk_product_comment_product foreign key (product_id) references product (id);