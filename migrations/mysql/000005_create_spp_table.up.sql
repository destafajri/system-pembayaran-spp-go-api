CREATE TABLE spp (
   id uuid primary key,
   siswa_id uuid not null,
   no_spp TEXT not null UNIQUE,
   jatuh_tempo date not null,
   jumlah TEXT not null,
   is_active bool not null DEFAULT true,
   created_at timestamp(6) not null,
   updated_at timestamp(6) not null,
   deleted_at timestamp(6),

   CONSTRAINT fk_siswa_id FOREIGN KEY(siswa_id) REFERENCES siswa(id),
   CONSTRAINT unique_spp_id UNIQUE(id)
);