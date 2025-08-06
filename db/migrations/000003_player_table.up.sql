create table if not exists player
(
    id                 int          not null auto_increment,
    campaign_id        int          not null,
    `name`             varchar(128) not null,
    race               varchar(128) not null,
    class              varchar(128) not null,
    armor_class        int          not null default 10,
    hit_points         int          not null default 1,
    passive_perception int          not null default 10,
    languages          varchar(256) not null default 'common',
    movement           int          not null default 30,
    primary key (id),
    foreign key (campaign_id) references campaign (id) on delete cascade
);
