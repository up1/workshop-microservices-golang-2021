CREATE TABLE users(
   id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   age INT
);

INSERT INTO users VALUES (1, 'name1', 20);
INSERT INTO users VALUES (2, 'name2', 30);
INSERT INTO users VALUES (3, 'name3', 40);