CREATE TABLE IF NOT EXISTS users (
    user_id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);