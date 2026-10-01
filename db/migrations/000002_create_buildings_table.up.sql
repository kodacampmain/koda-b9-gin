CREATE TABLE buildings (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    area INT NOT NULL,
    address VARCHAR(255) NOT NULL,
    person_id INT REFERENCES persons (id)
);