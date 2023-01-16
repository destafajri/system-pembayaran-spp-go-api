CREATE TABLE bayar (
   id uuid primary key,
   spp_id uuid not null,
   tanggal_bayar date not null,

   CONSTRAINT fk_spp_id FOREIGN KEY(spp_id) REFERENCES spp(id),
   CONSTRAINT unique_bayar_id UNIQUE(id)
);