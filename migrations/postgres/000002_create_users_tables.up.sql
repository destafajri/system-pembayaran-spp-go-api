CREATE TABLE users (
   id TEXT primary key,
   name TEXT not null,
   phone TEXT not null UNIQUE,
   role TEXT,
   password TEXT
);