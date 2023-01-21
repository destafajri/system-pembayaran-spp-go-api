CREATE TABLE kelas (
   id uuid primary key,
   guru_id uuid not null,
   kelas TEXT not null UNIQUE,
   created_at timestamp(6) not null,
   updated_at timestamp(6) not null,
   deleted_at timestamp(6),

   CONSTRAINT fk_guru_id FOREIGN KEY(guru_id) REFERENCES guru(id),
   CONSTRAINT unique_kelas_id UNIQUE(id)
);