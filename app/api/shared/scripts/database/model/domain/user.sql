-- user.sql
-- Initialies User Table

CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    email TEXT UNIQUE,
    username TEXT,
    password VARCHAR);