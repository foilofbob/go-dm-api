create table if not exists experience
(
    id            int not null auto_increment,
    campaign_id   int not null,
    `description` varchar(1000), -- e.g. 1x Wereboar, 2x Werewolves, and 4x Wererats
    xp            int not null,  -- raw calculated amount
    finalized     boolean,       -- planned, or has it happened?
    primary key (id),
    foreign key (campaign_id) references campaign (id) on delete cascade
);
