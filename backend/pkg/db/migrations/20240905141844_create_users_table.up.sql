CREATE TABLE users (
                       email VARCHAR(50) PRIMARY KEY,
                       name VARCHAR(30),
                       password VARCHAR(30),
                       hash_password VARCHAR(80)
);
