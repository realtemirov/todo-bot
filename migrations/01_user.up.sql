CREATE TABLE IF NOT EXISTS "users" (
    "id" bigint PRIMARY KEY,
    "first_name" varchar(64) NOT NULL,
    "last_name" varchar(64),
    "username" varchar(32),
    "is_admin" boolean DEFAULT false,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp ,
    "photo_url" text
);