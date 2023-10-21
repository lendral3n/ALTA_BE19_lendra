CREATE DATABASE alta_online_shop

USE alta_online_shop;

-- Membuat tabel
CREATE TABLE product_types (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE operators (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_descriptions (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    product_type_id INT(11),
    operator_id INT(11),
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    price INT,
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 	 update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    FOREIGN KEY (operator_id) REFERENCES operators(id)
);

CREATE TABLE payment_methods (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE users (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  status SMALLINT,
  dob DATE,
  gender CHAR(1),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  id INT(11) PRIMARY KEY AUTO_INCREMENT,
  user_id INT(11),
  payment_method_id INT(11),
  status VARCHAR(10),
  total_qty INT(11),
  total_price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
 transaction_id INT(11) PRIMARY KEY AUTO_INCREMENT, 
 product_id INT(11), 
 status VARCHAR(10), 
 qty INT(11), 
 price NUMERIC(25,2), 
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 FOREIGN KEY (transaction_id) REFERENCES transactions(id), 
 FOREIGN KEY (product_id) REFERENCES products(id)
);


INSERT INTO product_types (name) 
VALUES ('Smartphone'), ('Tablet'), ('Laptop');

INSERT INTO operators (name) 
VALUES ('Operator1'),
		 ('Operator2'),
	 	 ('Operator3'),
	 	 ('Operator4'),
	 	 ('Operator5'),
	 	 ('Operator6'),
	 	 ('Operator7'),
		 ('Operator8');

INSERT INTO products (product_type_id, operator_id, code, name, status, price)
VALUES (1, 1, 'HP001', 'Samsung Galaxy S21', 1, 1000),
       (1, 2, 'HP002', 'iPhone 13', 1, 1200);

INSERT INTO products (product_type_id, operator_id, code, name, status, price)
VALUES (2, 3, 'Tab001', 'iPad Pro', 1, 800),
       (2, 4, 'Tab002', 'Samsung Galaxy Tab S7', 1, 700),
       (2, 5, 'Tab003', 'Microsoft Surface Pro 7', 1, 900);

INSERT INTO products (product_type_id, operator_id, code, name, status, price)
VALUES (3, 6, 'Lap001', 'MacBook Pro', 1, 1500),
       (3, 7, 'Lap002', 'Dell XPS 13', 1, 1200),
       (3, 8, 'Lap003', 'Lenovo ThinkPad X1 Carbon', 1, 1300);

INSERT INTO product_descriptions (description)
VALUES
  ('Samsung Galaxy S21 adalah smartphone flagship dengan fitur kamera yang sangat baik.'),
  ('iPhone 13 adalah iPhone terbaru dengan chip A15 Bionic yang powerful.'),
  ('iPad Pro adalah tablet andalan dengan layar Retina yang indah dan dukungan Apple Pencil.'),
  ('Samsung Galaxy Tab S7 adalah tablet Android dengan layar Super AMOLED dan S Pen.'),
  ('Microsoft Surface Pro 7 adalah laptop 2-in-1 dengan layar sentuh dan kemampuan tablet.'),
  ('MacBook Pro adalah laptop kelas atas dengan Retina Display dan performa yang luar biasa.'),
  ('Dell XPS 13 adalah laptop ultrabook dengan desain premium dan layar InfinityEdge.'),
  ('Lenovo ThinkPad X1 Carbon adalah laptop bisnis yang ringan dan tahan lama.');

INSERT INTO payment_methods (name, status)
VALUES ('Cash Or Duel', 1),
       ('PayLater', 1),
       ('Bank Transfer', 1);

INSERT INTO users (status, dob, gender)
VALUES (1, '1990-05-15', 'M'),
       (1, '1988-12-10', 'F'),
       (1, '1995-07-22', 'M'),
       (1, '1982-09-30', 'F'),
       (1, '1998-03-05', 'M');
       
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES
  (1, 1, 'Completed', 2, 2200.00),
  (1, 2, 'Completed', 1, 1200.00),
  (1, 3, 'In Progress', 3, 3300.00);

-- User 2
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES
  (2, 1, 'Completed', 3, 3300.00),
  (2, 2, 'Completed', 2, 2400.00),
  (2, 3, 'In Progress', 1, 900.00);

-- User 3
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES
  (3, 1, 'In Progress', 1, 1100.00),
  (3, 2, 'In Progress', 2, 2200.00),
  (3, 3, 'Completed', 3, 3300.00);
  
  
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (1, 1, 'Completed', 1, 1000.00),
  (1, 2, 'Completed', 1, 1200.00);

-- Transaksi 2 User 1
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (2, 1, 'Completed', 2, 2000.00);

-- Transaksi 3 User 1
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (3, 2, 'In Progress', 3, 3600.00);

-- Transaksi 1 User 2
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (4, 3, 'Completed', 1, 800.00),
  (4, 4, 'Completed', 2, 1400.00);

-- Transaksi 2 User 2
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (5, 4, 'Completed', 1, 700.00),
  (5, 5, 'In Progress', 2, 1800.00);

-- Transaksi 3 User 2
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (6, 5, 'In Progress', 1, 900.00),
  (6, 6, 'In Progress', 2, 2400.00);

-- Transaksi 1 User 3
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (7, 6, 'Completed', 3, 4500.00);

-- Transaksi 2 User 3
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (8, 7, 'In Progress', 1, 1200.00),
  (8, 8, 'In Progress', 2, 2600.00);

-- Transaksi 3 User 3
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
  (9, 8, 'Completed', 2, 2600.00),
  (9, 1, 'In Progress', 1, 1000.00);
  '
  
  SELECT * FROM users WHERE gender = 'M';
  
  SELECT * FROM products WHERE id = 3;

  SELECT * FROM users WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND name LIKE '%a%';
  
  SELECT COUNT(*) FROM users WHERE gender = 'F';
  
  SELECT * FROM users ORDER BY name;
  
  SELECT * FROM transactions WHERE id IN (SELECT transaction_id FROM transaction_details WHERE product_id = 3) LIMIT 5;
  
  UPDATE products SET name = 'product dummy' WHERE id = 1;
  
  UPDATE transaction_details SET qty = 3 WHERE product_id = 1;
  
  DELETE FROM products WHERE id = 1;
  
  DELETE FROM products WHERE product_type_id = 1;


