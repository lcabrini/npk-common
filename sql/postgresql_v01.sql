drop type if exists user_status;
create type user_status as enum(
    'new',
    'active',
    'inactive',
    'deleted'
);

drop table if exists users;
create table users(
    id uuid,
    username varchar(50) not null,
    password varchar(50) not null,
    created_at timestamp not null default current_timestamp,
    status user_status not null,
    primary key(id),
    unique(password)
);

insert into users(id, username, password, status) values(
    '6c6cc02e-cca6-451d-9928-1f4d898c838e',
    'sa',
    '%npK-s3Kr3T%',
    'new'
);
