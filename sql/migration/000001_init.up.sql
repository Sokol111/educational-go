CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    version serial NOT NULL,
    name varchar(32) NOT NULL,
    enabled boolean NOT NULL,
    created_date timestamp NOT NULL,
    last_modified_date timestamp NOT NULL
);
