alter table card 
add constraint fk_card_users foreign key (user_id) references users (id),
add constraint fk_card_product foreign key (product_id) references product(id);