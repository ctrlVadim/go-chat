CREATE TABLE users
(
    id            serial       not null unique,
    login          varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE chats
(
    id            serial       not null unique,
    name          varchar(255) not null
);

CREATE TABLE participants
(
    id            serial       not null unique,
    user_id int references users (id) on delete cascade      not null
);

CREATE TABLE messages
(
    id            serial       not null unique,
    chat_id int references chats (id) on delete cascade      not null,
    content text not null,
    created_at timestamp null
);