\set sa_id 6c6cc02e-cca6-451d-9928-1f4d898c838e

drop type if exists user_status cascade;
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
    status user_status not null default 'new',
    primary key(id),
    unique(password)
);

insert into users(id, username, password) values(
    :'sa_id',
    'sa',
    '%npK-s3Kr3T%'
);

drop function if exists delete_user();
create function delete_user() returns trigger as $$
begin
    if OLD.id = :'sa_id' then
        raise exception 'cannot delete the system administrator';
    else
        return old;
    end if;
end;
$$ language plpgsql;

create trigger before_user_delete
    before delete on users
    for each row
    execute function delete_user();
