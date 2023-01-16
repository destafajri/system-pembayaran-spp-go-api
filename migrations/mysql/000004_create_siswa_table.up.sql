CREATE TABLE siswa (
   id uuid primary key,
   user_id uuid not null UNIQUE,
   kelas_id uuid not null,
   nis int not null UNIQUE,
   nama TEXT not null,
   angkatan TEXT not null,
   is_active bool not null DEFAULT true,

   CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
   CONSTRAINT fk_kelas_id FOREIGN KEY(kelas_id) REFERENCES kelas(id),
   CONSTRAINT unique_siswa_id UNIQUE(id)
);