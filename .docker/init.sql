CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4(),
    username varchar(64) NOT NULL,
    email VARCHAR(320) NOT NULL,
    password varchar(128) NOT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS shorteners (
    id uuid DEFAULT uuid_generate_v4(),
    url_shortened varchar(8) NOT NULL,  
    url_original varchar(128) NOT NULL,
    user_id uuid DEFAULT uuid_generate_v4(),
    visits INT DEFAULT 0 NOT NULL,
    PRIMARY KEY (id)
);
