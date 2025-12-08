create table if not exists category
(
    id          int          not null auto_increment,
    campaign_id int          not null,
    `name`      varchar(100) not null,
    primary key (id),
    foreign key (campaign_id) references campaign(id) on delete cascade
);