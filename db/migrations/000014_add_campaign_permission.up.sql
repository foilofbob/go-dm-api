CREATE TABLE IF NOT EXISTS user_permission
(
    id int not null auto_increment,
    user_id int not null,
    campaign_id int not null,
    primary key (id),
    foreign key (user_id) references `user` (id) on delete cascade,
    foreign key (campaign_id) references `campaign` (id) on delete cascade
);
