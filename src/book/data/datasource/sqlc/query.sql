-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBook :many
SELECT * FROM books
ORDER BY name;

-- name: CreateBook :one
INSERT INTO books (
  name, author
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: UpdateBook :one
UPDATE books
  set name = $2,
  author = $3
WHERE id = $1
RETURNING *;
