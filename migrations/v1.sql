CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_credentials (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    email varchar(100) UNIQUE NOT NULL,
    password varchar(100) NOT NULL
);