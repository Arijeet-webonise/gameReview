-- +goose Up
-- SQL in this section is executed when the migration is applied.

alter table games add column video varchar(255);
alter table games add column video_type varchar(255);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table games drop column video;
alter table games drop column video_type;
