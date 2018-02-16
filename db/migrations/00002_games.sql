-- +goose Up
-- SQL in this section is executed when the migration is applied.

create table Genre(
  id Serial primary key,
  name varchar(255) not null
);

create table Games(
    id Serial primary key,
    title varchar(255) not null,
    developer varchar(255) null,
    summary varchar(255) null,
    rating varchar(4) not null,
    image_name varchar(255) null
);

create table GenreToGameRelation(
  id Serial primary key,
  game int references Games(id),
  genre int references Genre(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table GenreToGameRelation;
drop table Games;
drop table Genre;
