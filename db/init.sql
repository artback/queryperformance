CREATE
EXTENSION IF NOT EXISTS pg_stat_statements;

CREATE TABLE users
(
    username text primary key,
    password text NOT NULL,
    role     text DEFAULT 'buyer',
    deposit  int  DEFAULT 0
);

CREATE TABLE products
(
    name      text primary key,
    seller_id text,
    CONSTRAINT fk_seller
        FOREIGN KEY (seller_id)
            REFERENCES users (username) ON DELETE CASCADE
);

CREATE TABLE transactions
(
    id           serial primary key,
    product_name text,
    username     text,
    amount       INT default 1,
    price        INT,
    CONSTRAINT fk_product_name
        FOREIGN KEY (product_name)
            REFERENCES products (name),
    CONSTRAINT fk_username
        FOREIGN KEY (username)
            REFERENCES users (username)
);


Select *
from users;
Select *
from users;
Select *
from users;

INSERT INTO users(username, password)
VALUES ('user1', 'password');
INSERT INTO users(username, password)
VALUES ('user2', 'password');
