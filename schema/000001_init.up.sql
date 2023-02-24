
CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE projector
(
    id            serial       not null unique,
    name          varchar(255) not null,
    category_id   int references category (id) on delete cascade      not null,
    quantity      int not null ,
    brightness    varchar(255) not null ,
    contrast      varchar(255) not null ,
    price         float not null

);

CREATE TABLE monitor
(
    id            serial       not null unique,
    name          varchar(255) not null,
    category_id   int references category (id) on delete cascade      not null,
    quantity      int not null ,
    brightness    varchar(255) not null ,
    contrast      varchar(255) not null ,
    price         float not null

);

CREATE TABLE video_wall
(
    id            serial       not null unique,
    name          varchar(255) not null,
    category_id   int references category (id) on delete cascade      not null,
    quantity      int not null ,
    brightness    varchar(255) not null ,
    contrast      varchar(255) not null ,
    price         float not null

);

CREATE TABLE mount
(
    id          serial       not null unique,
    name varchar(255) not null ,
    categoryId int references category (id) on delete cascade not null,
    quantity int not null ,
    p_size varchar(255) not null,
    price float not null
);

CREATE TABLE category
(
    id          serial       not null unique,
    name        varchar(255) not null
);

CREATE TABLE commercial_quantity
(
    id          serial not null,
    productId   int not null,
    extraPosId  int not null,
    reciever    varchar(255) not null
);

CREATE TABLE users_comm_quantity
(
    id  serial not null,
    usersId   int references users (id) on delete cascade      not null,
    commQuantityId int references commercial_quantity (id) on delete cascade      not null
);


