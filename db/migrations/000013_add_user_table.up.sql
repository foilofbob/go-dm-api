create table if not exists user
(
    id int not null auto_increment,
    username varchar(30) not null unique,
    password varchar(200) not null,
    primary key (id)
)
