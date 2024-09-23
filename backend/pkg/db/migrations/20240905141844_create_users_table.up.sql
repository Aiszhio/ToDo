CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(50) UNIQUE,
                       name VARCHAR(30) UNIQUE,
                       password VARCHAR(30),
                       hash_password VARCHAR(80)
);

CREATE TABLE notes (
                       id SERIAL PRIMARY KEY,
                       user_id INT REFERENCES users(id) ON DELETE CASCADE,
                       title VARCHAR(100) NOT NULL,
                       created_to DATE NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW()
);
