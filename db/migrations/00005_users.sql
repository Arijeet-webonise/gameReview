-- +goose Up
-- SQL in this section is executed when the migration is applied.

create table users(
  id SERIAL Primary key,
  firstname varchar(255) not null,
  lastname varchar(255) null,
  email varchar(255) null,
  username varchar(255) not null,
  password varchar(255) not null,
  roles varchar(255) null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table users;
