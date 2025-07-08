-- +goose Up
CREATE TYPE gender_type AS ENUM ('male', 'female', 'unknown');


CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    gender gender_type,
    age INTEGER,
    country_id VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX idx_people_name ON people(name);
CREATE INDEX idx_people_gender ON people(gender);
CREATE INDEX idx_people_country ON people(country_id);

-- +goose Down
DROP TABLE IF EXISTS people;
DROP TYPE IF EXISTS gender_type;