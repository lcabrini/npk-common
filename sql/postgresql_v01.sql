drop table if exists users;
create table users(
    id uuid,
    username varchar(50) not null,
    password varchar(50) not null,
    primary key(id),
    unique(password)
);

insert into users(id, username, password) values(
    '6c6cc02e-cca6-451d-9928-1f4d898c838e',
    'sa',
    '%npK-s3r3T%'
);

