create extension if not exists "uuid-ossp";

create table if not exists account
(
    id                        uuid        not null
        constraint df_account_id         default uuid_generate_v4()
        constraint pk_account            primary key,
    email           varchar(64)   not null
        constraint uk_account_email unique,
    username        varchar(64)   not null
        constraint uk_account_username unique,
    full_name       varchar(64)   not null,
    password        text          not null,
    hash            varchar(1024) not null,
    deleted         bool          not null
        constraint df_account_deleted default false,
    created_at timestamp
        constraint df_account_created_at default now(),
    updated_at timestamp
);

create or replace rule account_soft_deletion as on delete to account do instead
(
    update account set deleted = true where id = old.id and not deleted
);

create table if not exists role
(
    id   uuid        not null
        constraint df_role_id default uuid_generate_v4()
        constraint pk_role primary key,
    name varchar(64) not null
        constraint uk_role_name unique
);

create table if not exists account_role
(
  account_id uuid not null,
  role_id    uuid not null,
    constraint pk_account_role primary key (account_id, role_id),
    constraint fk_account_role_account_id foreign key (account_id)
        references account (id),
    constraint fk_account_role_role_id foreign key (role_id)
        references role (id)
);

create table if not exists password_reset
(
    email      varchar(50) not null
        constraint uk_password_reset unique,
    token      text        not null,
    created_at timestamp
        constraint df_password_reset_created_at default now()
);