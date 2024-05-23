-- Upewnij się, że schemat 'account' istnieje
CREATE SCHEMA IF NOT EXISTS account;


CREATE TABLE IF NOT EXISTS account.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    stream_guid UUID DEFAULT NULL,
    settings_id SERIAL UNIQUE
);

CREATE TABLE IF NOT EXISTS account.settings (
    id SERIAL PRIMARY KEY,
    user_id SERIAL UNIQUE NOT NULL,
    -- Definicja klucza obcego odnoszącego się do tabeli 'users'
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES account.users(id) ON DELETE CASCADE
);