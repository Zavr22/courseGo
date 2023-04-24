
CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          int          not null
);

CREATE TABLE category
(
    id          serial       not null unique,
    name        varchar(255) not null
);

CREATE TABLE projector
(
    id            serial       not null unique,
    name          varchar(255) not null,
    category_id   int references category (id) on delete cascade      not null,
    quantity      int not null ,
    brightness    varchar(255) not null ,
    contrast      varchar(255) not null ,
    price         float not null,
    weight        float not null,
    focal_distance float not null


);

CREATE TABLE monitor
(
    id            serial       not null unique,
    name          varchar(255) not null,
    category_id   int references category (id) on delete cascade      not null,
    quantity      int not null ,
    brightness    varchar(255) not null ,
    contrast      varchar(255) not null ,
    price         float not null,
    weight        float not null
);

CREATE TABLE video_wall
(
    id            serial       not null unique,
    name          varchar(255) not null,
    category_id   int references category (id) on delete cascade      not null,
    quantity      int not null ,
    brightness    varchar(255) not null ,
    contrast      varchar(255) not null ,
    price         float not null,
    weight        float not null

);

CREATE TABLE mount
(
    id          serial       not null unique,
    name varchar(255) not null ,
    categoryId int references category (id) on delete cascade not null,
    quantity int not null ,
    max_weight float not null,
    price float not null,
    maxWeight float not null,
    roi float not null
);



CREATE TABLE commercial_quantity
(
    id         serial  unique,
    products    json not null,
    status      varchar(40) not null
);

CREATE TABLE users_comm_quantity
(
    id  serial not null,
    usersId   int not null,
    commQuantityId int not null
);

CREATE TABLE settings
(
    id int,
    roi float not null
);

CREATE TABLE products
(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    category_id int REFERENCES category(id) ON DELETE CASCADE NOT NULL
);

