create table users 
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255),
    phone varchar(255) unique
);

create table bikes 
(
    id serial primary key,
    name varchar(255) not null,
    price int not null,
    description varchar(255)
);

create table rollers 
(
    id serial primary key,
    name varchar(255) not null,
    price int   not null ,
    description varchar(255),
    size numeric(3,1) not null
);

create table rollersReservation
(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    roller_id int references rollers(id) on delete cascade not null
);
create table bikeReservation
(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    bike_id int references bikes(id) on delete cascade not null
);