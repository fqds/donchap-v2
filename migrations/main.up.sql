CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null unique,
    encrypted_password varchar not null
);

CREATE TABLE lobbies (
    id bigserial not null primary key,
    name varchar not null unique,
    master_id bigserial not null
);

CREATE TABLE lobby_parameters (
    id bigserial not null primary key,
    lobby_id bigserial not null,
    name varchar not null,
    formula varchar,
    dropdown_list varchar,
    is_visible boolean
);

CREATE TABLE lobby_players (
    id bigserial not null primary key,
    user_id bigserial not null,
    lobby_id bigserial not null,
    position bigserial not null
);

CREATE TABLE player_parameters (
    id bigserial not null primary key,
    player_id bigserial not null,
    parameter_id bigserial not null,
    value varchar
);