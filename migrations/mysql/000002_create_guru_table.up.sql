CREATE TABLE guru (
   id uuid primary key,
   user_id uuid not null UNIQUE,
   nama TEXT not null UNIQUE,
   is_active bool not null DEFAULT true,
   
   CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
   CONSTRAINT unique_guru_id UNIQUE(id)
);