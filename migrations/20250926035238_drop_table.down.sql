
CREATE TABLE batas_wilayah (
    id SERIAL PRIMARY KEY,
    profil_desa_id INT,
    arah VARCHAR(20),
    desa_batas VARCHAR(100)
);

CREATE TABLE komposisi_penduduk (
    id SERIAL PRIMARY KEY,
    profil_desa_id INT,
    kategori VARCHAR(20),
    jumlah INT
);

CREATE TABLE mata_pencaharian (
    id SERIAL PRIMARY KEY,
    profil_desa_id INT,
    jenis VARCHAR(20),
    persentase DECIMAL(5,2)
);

CREATE TABLE fasilitas_umum (
    id SERIAL PRIMARY KEY,
    profil_desa_id INT,
    jenis VARCHAR(20),
    jumlah INT
);
