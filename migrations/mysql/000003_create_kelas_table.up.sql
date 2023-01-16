CREATE TABLE kelas (
   id uuid primary key,
   guru_id uuid not null,
   kelas TEXT not null UNIQUE,

   CONSTRAINT fk_guru_id FOREIGN KEY(guru_id) REFERENCES guru(id),
   CONSTRAINT unique_kelas_id UNIQUE(id)
);