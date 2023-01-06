CREATE TABLE IF NOT EXISTS todos (
    id text PRIMARY KEY,
    created_at timestamp NOT NULL,
    deleted_at timestamp,
    user_id bigint NOT NULL,
    title   text NOT NULL,
    description text ,
    photo_url text,
    file_url text,
    deadline timestamp,
    is_set bool DEFAULT false,
    notification timestamp
);