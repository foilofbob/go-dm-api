alter table player
    add column strength int          not null default 10,
    add column dexterity int          not null default 10,
    add column constitution int          not null default 10,
    add column intelligence int          not null default 10,
    add column wisdom int          not null default 10,
    add column charisma int          not null default 10,
    add column proficiencies          varchar(256) not null default '',
    add column saves         varchar(256) not null default '';
