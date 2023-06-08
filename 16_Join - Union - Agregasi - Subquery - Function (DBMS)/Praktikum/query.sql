-- Data Manipulation Language (DML)
-- 1a 
--insert 5 operators pada table operators
INSERT INTO operators (name, created_at, updated_at) VALUES
("Operator 1", "2023-05-31", "2023-06-07"),
("Operator 2", "2023-05-31", "2023-06-07"),
("Operator 3", "2023-05-31", "2023-06-07"),
("Operator 4", "2023-05-31", "2023-06-07"),
("Operator 5", "2023-05-31", "2023-06-07");

-- 1b 
-- insert 3 product type.
INSERT INTO product_types (name, created_at, updated_at) VALUES
("Product Type 1", "2023-05-31", "2023-06-07"),
("Product Type 2", "2023-05-31", "2023-06-07"),
("Product Type 3", "2023-05-31", "2023-06-07");

-- 1 c, d, e
-- insert 2 product dengan product type id = 1, dan operators id = 3
-- insert 3 product dengan product type id = 2, dan operators id = 1
-- insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO products (operator_id, name, product_type_id, product_description_id, price, status, created_at, updated_at) VALUES
(3, "Product 1C 1", 1, 1, 150000, 1, "2023-05-31", "2023-06-07"),
(3, "Product 1C 2", 1, 1, 100000, 2, "2023-05-31", "2023-06-07"),
(1, "Product 1D 1", 2, 1, 120000, 1, "2023-05-31", "2023-06-07"),
(1, "Product 1D 2", 2, 1, 60000, 2, "2023-05-31", "2023-06-07"),
(1, "Product 1D 3", 2, 1, 30000, 3, "2023-05-31", "2023-06-07"),
(4, "Product 1E 1", 3, 1, 500000, 1, "2023-05-31", "2023-06-07"),
(4, "Product 1E 2", 3, 1, 250000, 2, "2023-05-31", "2023-06-07"),
(4, "Product 1E 3", 3, 1, 125000, 3, "2023-05-31", "2023-06-07");

-- 1f
-- insert product description pada setiap product
INSERT INTO product_descriptions (id, description, created_at, updated_at) VALUES
(1, "This Descriptions of Product 1", "2023-05-31", "2023-06-07"),
(2, "This Descriptions of Product 2", "2023-05-31", "2023-06-07"),
(3, "This Descriptions of Product 3", "2023-05-31", "2023-06-07"),
(4, "This Descriptions of Product 4", "2023-05-31", "2023-06-07"),
(5, "This Descriptions of Product 5", "2023-05-31", "2023-06-07");

-- 1g
-- insert 3 payment methods
INSERT INTO payment_methods (id, name, status, created_at, updated_at) VALUES
(1, "Payment Method 1", 1, "2023-05-31", "2023-06-07"),
(2, "Payment Method 2", 1, "2023-05-31", "2023-06-07"),
(3, "Payment Method 3", 1, "2023-05-31", "2023-06-07");

-- 1h
-- insert 5 user pada tabel user
INSERT INTO users (name, address, birth_date, status_user, gender, created_at, updated_at) VALUES
("User A", "Alamat A", "2001-01-2", 1, "M", "2023-05-31", "2023-06-07"),
("User B", "Alamat B", "2001-01-2", 1, "F", "2023-05-31", "2023-06-07"),
("User C", "Alamat C", "2001-01-2", 1, "M", "2023-05-31", "2023-06-07"),
("User D", "Alamat D", "2001-01-2", 1, "F", "2023-05-31", "2023-06-07"),
("User E", "Alamat E", "2001-01-2", 1, "M", "2023-05-31", "2023-06-07");

-- 1i
-- insert 3 transaksi di masing-masing user
INSERT INTO transactions (user_id, payment_method_id, total_quantity, total_amount, status, created_at, updated_at) VALUES
(1, 1, 100000, 20, "Paid", "2023-05-31", "2023-06-07"),
(2, 1, 150000, 10, "Unpaid", "2023-05-31", "2023-06-07"),
(3, 2, 200000, 5, "Canceled", "2023-05-31", "2023-06-07");

-- 1j
-- insert 3 product di masing-masing transaksi
INSERT INTO transaction_details (transaction_id, product_id, quantity, total_price, status, created_at, updated_at) VALUES
(1, 1, 1, 100000, "Paid", "2023-05-31", "2023-06-07"),
(1, 3, 1, 150000, "Unpaid", "2023-05-31", "2023-06-07"),
(3, 3, 1, 200000, "Canceled", "2023-05-31", "2023-06-07");

-- 2a
-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT name FROM users WHERE gender = 'M';

-- 2b
-- tampilkan product dengan id = 3
SELECT * FROM products WHERE id = 3;

-- 2c
-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’
SELECT * FROM users WHERE created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';SELECT * FROM `users` WHERE 1

-- 2d
-- Hitung jumlah user / pelanggan dengan status gender Perempuan
SELECT COUNT(*) as female_user_count FROM users WHERE gender = 'F';

-- 2e
-- Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users ORDER BY name ASC;

-- 2f
-- Tampilkan 5 data pada data product
SELECT * FROM products LIMIT 5;

-- 3a
-- Ubah data product id 1 dengan nama ‘product dummy’
UPDATE products SET name = 'product dummy' WHERE id = 1;

-- 3b
-- Update qty = 3 pada transaction detail dengan product id = 1
UPDATE transaction_details SET quantity = 1 WHERE product_id = 1;

-- 4a
-- Delete data pada tabel product dengan id = 1
DELETE FROM products WHERE id = 1;

-- 4b
-- Delete pada pada tabel product dengan product type id 1
DELETE FROM products WHERE product_type_id = 1;

-- Join, Union, Sub query, Function
-- 1. Gabungkan data transaksi dari user id 1 dan user id 
SELECT * FROM transactions WHERE user_id = 1 
UNION 
SELECT * FROM transactions WHERE user_id = 2;

-- 2. Tampilkan jumlah harga transaksi user id 1
SELECT SUM(total_amount) as jumlah_harga FROM transactions t 
JOIN transaction_details td ON t.id = td.transaction_id 
WHERE t.user_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2
SELECT COUNT(*) as total_transaksi FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
WHERE p.product_type_id = 2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan
SELECT p.*, pt.name as product_type_name FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user
SELECT t.*, p.name as product_name, u.name as user_name FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
JOIN users u ON t.user_id = u.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga 
--    dengan transaction id yang dimaksud.

--7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate 
--   berdasarkan qty data transaction id yang dihapus.

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT * FROM products WHERE id NOT IN (
   SELECT product_id
   FROM transaction_details
);