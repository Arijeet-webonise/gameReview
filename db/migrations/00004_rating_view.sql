-- +goose Up
-- SQL in this section is executed when the migration is applied.

create view rating_view as
  select game, sum(rating) as total_rating, count(rating) from comment group by game;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop view rating_view;
