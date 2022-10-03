-- +goose Up

-- +goose StatementBegin
create table events
(
    uuid             UUID primary key not null ,
    title            varchar,
    datetime         timestamp,
    start_datetime   timestamptz,
    end_datetime     timestamptz,
    description      varchar,
    userid           UUID,
    remind_time_before varchar
);
-- +goose StatementEnd

-- +goose StatementBegin
create table users
(
    uuid           UUID primary key not null,
    datetime       timestamptz

);
-- +goose StatementEnd

-- +goose Down
drop table users;
drop table events;