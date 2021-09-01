-- +goose Up
-- +goose StatementBegin

CREATE SEQUENCE IF NOT EXISTS conference_id_seq AS BIGINT;

CREATE TABLE IF NOT EXISTS public.conferences
(
    id bigint NOT NULL DEFAULT nextval('conference_id_seq'),
    name text COLLATE pg_catalog."default" NOT NULL,
    event_time timestamp without time zone,
    participant_count integer NOT NULL DEFAULT 0,
    speaker_count integer NOT NULL DEFAULT 0,
    CONSTRAINT conferences_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS conferences_ix_user_id
    ON public.conferences USING btree
    (user_id ASC NULLS LAST);

-- +goose StatementEnd
