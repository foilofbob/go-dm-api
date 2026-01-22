create table if not exists spell
(
    `id`              int           not null auto_increment,
    `name`            varchar(50)   not null,
    `source`          varchar(10)   not null,
    `page`            varchar(10)   not null,
    `level`           varchar(10)   not null,
    `casting_time`    varchar(20)   not null,
    `duration`        varchar(50)   not null,
    `school`          varchar(25)   not null,
    `range`           varchar(50)   not null,
    `components`      varchar(500)  not null,
    `classes`         varchar(500)  not null,
    `variant_classes` varchar(500)  not null,
    `subclasses`      varchar(500)  not null,
    `description`     varchar(5000) not null,
    `higher_casting`  varchar(1000) not null,
    PRIMARY KEY (`id`)
);
