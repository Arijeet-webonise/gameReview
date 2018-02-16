-- +goose Up
-- SQL in this section is executed when the migration is applied.

create table comment(
  id SERIAL primary key,
  game int references Games(id),
  comment varchar(1000) not null,
  rating int not null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table comment;
