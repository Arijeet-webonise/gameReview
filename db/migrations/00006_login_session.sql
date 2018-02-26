-- +goose Up
-- SQL in this section is executed when the migration is applied.

create table session(
  uuid varchar(255) unique not null,
  userid int references users(id) not null primary key
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table session;
