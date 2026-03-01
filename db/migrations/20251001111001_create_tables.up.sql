CREATE TABLE profil_desa (
    id BIGSERIAL PRIMARY KEY,
    alamat TEXT,
    kecamatan VARCHAR(100),
    kabupaten VARCHAR(100),
    provinsi VARCHAR(100),
    kode_pos VARCHAR(10),
    jumlah_penduduk VARCHAR(10),
    jumlah_laki VARCHAR(10),
    jumlah_perempuan VARCHAR(10),
    jumlah_kk VARCHAR(10),
    tahun_pembentukan VARCHAR(4),
    telepon VARCHAR(20),
    email VARCHAR(150),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE sejarah (
    id BIGSERIAL PRIMARY KEY,
    profil_desa_id BIGINT REFERENCES profil_desa(id) ON DELETE CASCADE,
    body TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE demografis (
    id BIGSERIAL PRIMARY KEY,
    profil_desa_id BIGINT REFERENCES profil_desa(id) ON DELETE CASCADE,
    balita VARCHAR(50),
    anak VARCHAR(50),
    dewasa VARCHAR(50),
    lansia VARCHAR(50),
    pertanian VARCHAR(50),
    perdagangan VARCHAR(50),
    jasa VARCHAR(50),
    industri VARCHAR(50),
    sekolah VARCHAR(50),
    puskesmas VARCHAR(50),
    masjid VARCHAR(50),
    pasar_tradisional VARCHAR(50),
    pos_keamanan VARCHAR(50),
    balai_desa VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);


CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(100) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);
