CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4(),
    username varchar(64) NOT NULL,
    email VARCHAR(320) NOT NULL,
    password varchar(128) NOT NULL,
    PRIMARY KEY (id)
);
