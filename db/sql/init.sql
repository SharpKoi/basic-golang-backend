create extension if not exists "uuid-ossp"; -- you can use uuid-ossp to auto generate default uuid for posts

create table if not exists users (
    id serial primary key ,
    email varchar (40) unique not null ,
    password varchar (64) not null ,
    name varchar (16) ,
    age int check (age > 0),
    role varchar (10) not null ,
    createAt timestamp
);

create table if not exists posts (
--     uid uuid default uuid_generate_v4(),
    uid varchar(255) primary key ,
    userEmail varchar (40) not null ,
    text text not null ,
    createAt timestamp
);
