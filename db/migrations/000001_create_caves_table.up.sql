CREATE TABLE IF NOT EXISTS caves
(
   id serial PRIMARY KEY,
   title VARCHAR(100) UNIQUE NOT NULL,
   country_name VARCHAR(10) NOT NULL,
   region_name VARCHAR(20) NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   longitude FLOAT NOT NULL,
   latitude FLOAT NOT NULL
);