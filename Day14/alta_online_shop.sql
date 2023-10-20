-- create database

CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE user(
	id INT PRIMARY KEY AUTO_INCREMENT,
	nama VARCHAR(255),
	status_user BOOL DEFAULT true,
	jenis_kelamin ENUM('L', 'P'),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);

ALTER TABLE user ADD COLUMN tanggal_lahir DATE;

INSERT INTO user (nama, status_user, tanggal_lahir, jenis_kelamin)
VALUES
('lele', true, '2023-11-23', 'L'),
('lala', false, '1985-09-10', 'P'),
('lolo', true, '1997-08-20', 'P');

CREATE TABLE product (
	id INT PRIMARY KEY AUTO_INCREMENT,
	nama_produk VARCHAR(255),
	jenis_produk VARCHAR(255),
	deskripsi_produk TEXT,
	stok_harga INT,
	harga_produk DECIMAL(10, 2)
);

INSERT INTO product (nama_produk, jenis_produk, deskripsi_produk, stok_harga, harga_produk)
VALUES
('susu', 'Kategori 1', 'susu beruang', 100, 50.00),
('sasa', 'Kategori 2', 'sasa bukan rasa dari si dia 2', 75, 75.99),
('sisa', 'Kategori 1', 'hanya sisaan kenangan dari dia', 120, 30.50);

CREATE TABLE product_type (
	id INT PRIMARY KEY AUTO_INCREMENT,
	nama_tipe_produk VARCHAR(255),
	merk VARCHAR(255)
);

INSERT INTO product_type (nama_tipe_produk, merk)
VALUES
('Tipe Produk 1', 'Merk 1'),
('Tipe Produk 2', 'Merk 2'),
('Tipe Produk 3', 'Merk 1');


CREATE TABLE product_description (
	id INT PRIMARY KEY AUTO_INCREMENT,
	deskripsi TEXT
);

INSERT INTO product_description (deskripsi)
VALUES
('Produk Best Seller'),
('Produk Diskonan'),
('Produk Mahal');

CREATE TABLE payment_method (
	id INT PRIMARY KEY AUTO_INCREMENT,
	metode_pembayaran VARCHAR(255)
);

INSERT INTO payment_method (metode_pembayaran)
VALUES
('Kartu Kredit'),
('Transfer Bank'),
('e-Wallet'),
('COD (Cash Or Duel)');

CREATE TABLE transaction (
	id INT PRIMARY KEY AUTO_INCREMENT,
	id_pelanggan INT,
	tanggal_transaksi DATE,
	jumlah_total DECIMAL(10, 2),
	FOREIGN KEY (id_pelanggan) REFERENCES user(id)
);

INSERT INTO transaction (id_pelanggan, tanggal_transaksi, jumlah_total)
VALUES
(1, '2023-01-15', 100.50),
(2, '2023-01-20', 75.99),
(3, '2023-01-22', 30.50);

CREATE TABLE transaction_detail (
	id INT PRIMARY KEY AUTO_INCREMENT,
    id_transaksi INT,
    id_produk INT,
    kuantitas INT,
    harga_satuan DECIMAL(10, 2),
    FOREIGN KEY (id_transaksi) REFERENCES transaction(id),
    FOREIGN KEY (id_produk) REFERENCES product(id)
);

INSERT INTO transaction_detail (id_transaksi, id_produk, kuantitas, harga_satuan)
VALUES
(1, 1, 2, 50.00),
(2, 2, 1, 75.99),
(3, 3, 3, 30.50);

CREATE TABLE kurir (
	id INT PRIMARY KEY AUTO_INCREMENT,
	nama VARCHAR (255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);

ALTER TABLE kurir ADD COLUMN ongkos_dasar DECIMAL(10, 2);

INSERT INTO kurir (nama, ongkos_dasar)
VALUES
('Kurir 1', 10.00),
('Kurir 2', 12.50),
('Kurir 3', 8.75);

RENAME TABLE kurir TO shipping;

DROP TABLE shipping;

CREATE TABLE description (
	id INT PRIMARY KEY AUTO_INCREMENT,
	deskripsi TEXT
);

ALTER TABLE payment_method ADD COLUMN id_deskripsi INT, ADD CONSTRAINT fk_payment_method_description FOREIGN KEY(id_deskripsi) REFERENCES description(id);


CREATE TABLE alamat (
    id INT PRIMARY KEY AUTO_INCREMENT,
    id_pengguna INT,
    alamat TEXT,
    FOREIGN KEY (id_pengguna) REFERENCES user(id)
);

CREATE TABLE user_payment_method_detail (
    id INT PRIMARY KEY AUTO_INCREMENT,
    id_pengguna INT,
    id_metode_pembayaran INT,
    FOREIGN KEY (id_pengguna) REFERENCES user(id),
    FOREIGN KEY (id_metode_pembayaran) REFERENCES payment_method(id)
);
