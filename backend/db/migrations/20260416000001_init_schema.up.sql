-- Enable pgcrypto for UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create ENUM types
CREATE TYPE user_role AS ENUM ('admin', 'kasir');
CREATE TYPE payment_method AS ENUM ('cash', 'hutang');
CREATE TYPE debt_status AS ENUM ('lunas', 'belum');

-- Table: Tokos
CREATE TABLE tokos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nama_toko TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Table: Pelanggans (Buyers)
CREATE TABLE pelanggans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_toko UUID NOT NULL REFERENCES tokos(id) ON DELETE RESTRICT,
    nama TEXT NOT NULL,
    no_telp TEXT,
    alamat TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE UNIQUE INDEX idx_pelanggan_telp_toko_active ON pelanggans(no_telp, id_toko) WHERE deleted_at IS NULL;

-- Table: Users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_toko UUID NOT NULL REFERENCES tokos(id) ON DELETE RESTRICT,
    nama TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role user_role NOT NULL DEFAULT 'kasir',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Table: Produks
CREATE TABLE produks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_toko UUID NOT NULL REFERENCES tokos(id) ON DELETE RESTRICT,
    barcode TEXT NOT NULL,
    nama_produk TEXT NOT NULL,
    harga DECIMAL(15, 2) NOT NULL DEFAULT 0,
    stok INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE UNIQUE INDEX idx_produk_barcode_toko_active ON produks(barcode, id_toko) WHERE deleted_at IS NULL;

-- Table: Transaksis
CREATE TABLE transaksis (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_toko UUID NOT NULL REFERENCES tokos(id) ON DELETE RESTRICT,
    id_kasir UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    id_pelanggan UUID REFERENCES pelanggans(id) ON DELETE SET NULL,
    tanggal TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    total DECIMAL(15, 2) NOT NULL DEFAULT 0,
    metode_pembayaran payment_method NOT NULL DEFAULT 'cash',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Table: Detail Transaksis
CREATE TABLE detail_transaksis (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_transaksi UUID NOT NULL REFERENCES transaksis(id) ON DELETE CASCADE,
    id_produk UUID NOT NULL REFERENCES produks(id) ON DELETE RESTRICT,
    qty INT NOT NULL DEFAULT 1,
    harga_satuan DECIMAL(15, 2) NOT NULL,
    subtotal DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Table: Hutangs
CREATE TABLE hutangs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_transaksi UUID NOT NULL REFERENCES transaksis(id) ON DELETE CASCADE,
    total_hutang DECIMAL(15, 2) NOT NULL,
    sisa_hutang DECIMAL(15, 2) NOT NULL,
    status debt_status NOT NULL DEFAULT 'belum',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Table: Detail Hutangs
CREATE TABLE detail_hutangs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
    id_hutang UUID NOT NULL REFERENCES hutangs(id) ON DELETE CASCADE,
    id_produk UUID NOT NULL REFERENCES produks(id) ON DELETE RESTRICT,
    qty INT NOT NULL,
    subtotal DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Table: Pembayaran Hutangs
CREATE TABLE pembayaran_hutangs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_hutang UUID NOT NULL REFERENCES hutangs(id) ON DELETE CASCADE,
    tanggal_bayar TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    jumlah_bayar DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_users_toko ON users(id_toko);
CREATE INDEX idx_produks_toko ON produks(id_toko);
CREATE INDEX idx_transaksis_toko ON transaksis(id_toko);
CREATE INDEX idx_transaksis_kasir ON transaksis(id_kasir);
CREATE INDEX idx_transaksis_pelanggan ON transaksis(id_pelanggan);
CREATE INDEX idx_hutangs_transaksi ON hutangs(id_transaksi);
CREATE INDEX idx_pembayaran_hutang ON pembayaran_hutangs(id_hutang);
