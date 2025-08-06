create table campaign_setting
(
    id     int          not null auto_increment,
    `name` varchar(100) not null,
    primary key (id)
);

create table campaign
(
    id                  int          not null auto_increment,
    `name`              varchar(128) not null,
    current_player_xp   int          not null,
    campaign_setting_id int          not null,
    primary key (id),
    foreign key (campaign_setting_id) references campaign_setting (id)
);

create table if not exists note
(
    id             int          not null auto_increment,
    campaign_id    int          not null,
    reference_type varchar(30)  not null,
    reference_id   int          not null,
    category       varchar(20),
    title          varchar(100) not null,
    content        varchar(10000),
    primary key (id),
    foreign key (campaign_id) references campaign (id) on delete cascade
);

create table if not exists game_day
(
    id          int not null auto_increment,
    campaign_id int not null,
    in_game_day int not null default 1,
    `day`       int not null default 1,
    `month`     int not null default 1,
    `year`      int not null default 1,
    primary key (id),
    foreign key (campaign_id) references campaign (id) on delete cascade
);

create table if not exists `month`
(
    id                  int         not null auto_increment,
    campaign_setting_id int         not null,
    `name`              varchar(50) not null,
    num_days            int         not null default 30,
    `order`             int         not null,
    primary key (id),
    foreign key (campaign_setting_id) references campaign_setting (id)
);

create table if not exists week_day
(
    id                  int         not null auto_increment,
    campaign_setting_id int         not null,
    `name`              varchar(50) not null,
    `order`             int         not null,
    primary key (id),
    foreign key (campaign_setting_id) references campaign_setting (id)
);

create table if not exists calendar_event
(
    id                  int          not null auto_increment,
    campaign_setting_id int          not null,
    month_id            int          not null,
    `day`               int          not null,
    `name`              varchar(100) not null,
    `description`       varchar(2000),
    primary key (id),
    foreign key (campaign_setting_id) references campaign_setting (id),
    foreign key (month_id) references `month` (id)
);

create table if not exists calendar_cycle
(
    id                  int          not null auto_increment,
    campaign_setting_id int          not null,
    `name`              varchar(100) not null,
    `period`            int not null,
    primary key (id),
    foreign key (campaign_setting_id) references campaign_setting (id)
);

create table if not exists campaign_calendar_cycle_offset
(
    id                int not null auto_increment,
    campaign_id       int not null,
    calendar_cycle_id int not null,
    `offset`          int not null,
    primary key (id),
    foreign key (campaign_id) references campaign (id),
    foreign key (calendar_cycle_id) references calendar_cycle (id)
);
