create extension if not exists "uuid-ossp";

create table if not exists category 
(
  id uuid not null
    constraint df_category_id default uuid_generate_v4()
    constraint pk_category_id primary key,
  name varchar(100) not null
);

create table if not exists link 
(
  id uuid not null 
    constraint df_link_id default uuid_generate_v4()
    constraint pk_link_id primary key,
  content text unique not null,
  registered_at timestamp not null default now()
);

create table if not exists link_category
(
  link_id uuid not null,
  category_id uuid not null,

  constraint pk_link_category primary key (link_id, category_id),
  constraint fk_link_id foreign key (link_id) references link (id),
  constraint fk_category_id foreign key (category_id) references category (id)
);