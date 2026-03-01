CREATE TABLE visi_misi (
    id BIGSERIAL PRIMARY KEY,
    profil_desa_id BIGINT REFERENCES profil_desa(id) ON DELETE CASCADE,
    visi TEXT,
    misi TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);