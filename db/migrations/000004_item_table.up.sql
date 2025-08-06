create table if not exists item
(
    id            int          not null auto_increment,
    campaign_id   int          not null,
    `name`        varchar(100) not null,
    `description` varchar(5000),
    link          varchar(500),
    rarity        varchar(10), # common, uncommon, rare, very rare, legendary
    cost          varchar(20),
    requirements  varchar(64),
    is_container  boolean,     # denotes if other items can belong to this
    carried_by    varchar(20), # denotes a table - player, item
    carried_by_id int,         # id from corresponding table
    primary key (id)
);
