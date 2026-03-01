-- Ubah tipe kolom tahun dari INT ke VARCHAR(10)
ALTER TABLE apbd_desa
ALTER COLUMN tahun TYPE VARCHAR(4)
USING tahun::VARCHAR(4);
