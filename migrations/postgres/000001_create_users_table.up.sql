CREATE TABLE users (
   id uuid primary key,
   email TEXT not null UNIQUE,
   username TEXT not null UNIQUE,
   password TEXT not null,
   role TEXT not null,
   is_active bool not null DEFAULT true,
   created_at timestamp(6) not null,
   updated_at timestamp(6) not null,
   deleted_at timestamp(6),
   
   CONSTRAINT unique_user_id UNIQUE(id)
);