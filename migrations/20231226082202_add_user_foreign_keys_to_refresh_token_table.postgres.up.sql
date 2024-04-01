alter table refresh_token 
add constraint fk_refresh_token_users foreign key (user_id) references users (id);