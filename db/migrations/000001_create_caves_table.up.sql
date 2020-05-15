CREATE TABLE IF NOT EXISTS caves
(
   id serial PRIMARY KEY,
   title VARCHAR(100) UNIQUE NOT NULL,
   countryName VARCHAR(10) NOT NULL,
   regionName VARCHAR(20) NOT NULL,
   createdAt TIMESTAMP NOT NULL,
   updatedAt TIMESTAMP NOT NULL,
   longitude FLOAT NOT NULL,
   latitude FLOAT NOT NULL
);