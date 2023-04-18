CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null unique,
    encrypted_password varchar not null
);

CREATE TABLE transactions (
    id bigserial not null primary key,
    name varchar not null unique,
    master bigserial not null
);

CREATE TABLE lobby_parameters (
    id bigserial not null primary key,
    lobby_id bigserial not null,
    name varchar not null,
    formula varchar,
    dropdown_list varchar,
    is_visible boolean
);