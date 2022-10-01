CREATE TYPE fruit AS ENUM ('apple', 'banana', 'kiwi');

CREATE TABLE choices (
	id int PRIMARY KEY,
	fruits fruit[] NOT NULL
);

-- name: GetChoice :one
SELECT * FROM choices WHERE id = $1;
