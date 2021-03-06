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

-- Default admin user. For demonstrative purpose.
delete from users where email='admin';
insert into users (email, password, name, age, role, createAt) values ('admin', '607e707428137b535fd1e825ade61d58dbdc7147bd382c1cbd93a26e2e5ab7b4', 'admin', 1, 'admin', now())
on conflict (email) do nothing ;

create table if not exists posts (
--     uid uuid default uuid_generate_v4(),
    uid varchar(255) primary key ,
    userEmail varchar (40) not null ,
    text text not null ,
    createAt timestamp
);
