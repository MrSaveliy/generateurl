-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS urls
(
    id    BIGSERIAL,
    url   text,
    short varchar(10),
    CONSTRAINT shorts_pkey PRIMARY KEY (id)
);
select nextval('urls_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shorts;
-- +goose StatementEnd
