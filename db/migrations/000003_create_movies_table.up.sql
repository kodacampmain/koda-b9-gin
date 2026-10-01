CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(100) NOT NULL,
    rating NUMERIC(3, 1) NOT NULL CHECK (
        rating >= 0.0
        AND rating <= 10.0
    )
);