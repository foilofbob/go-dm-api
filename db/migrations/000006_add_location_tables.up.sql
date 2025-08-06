# Top level, e.g. "Ruins of Shattengrod" - linked to from menu
create table if not exists location
(
    id          int          not null auto_increment,
    campaign_id int          not null,
    `name`      varchar(100) not null,
    primary key (id),
    foreign key (campaign_id) references campaign(id) on delete cascade
);

# Logical divisions, e.g. "Floor Level 1" or "Vigil Borough"
create table if not exists sublocation
(
    id            int          not null auto_increment,
    campaign_id   int          not null,
    location_id   int          not null,
    `name`        varchar(100) not null,
    `description` varchar(1000),
    primary key (id),
    foreign key (campaign_id) references campaign(id) on delete cascade,
    foreign key (location_id) references location(id) on delete cascade
);

# Specific locations, e.g. "Smith Sharpe Weapons" or "Hillside Inn"
create table if not exists point_of_interest
(
    id             int          not null auto_increment,
    campaign_id    int          not null,
    sublocation_id int          not null,
    `name`         varchar(100) not null,
    primary key (id),
    foreign key (campaign_id) references campaign(id) on delete cascade,
    foreign key (sublocation_id) references sublocation(id) on delete cascade
);
