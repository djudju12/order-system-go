-- name: CreateProduct :one
INSERT INTO products (
   name,
   price, 
   description
) VALUES(
  $1, $2, $3 
) RETURNING *;

-- name: GetProduct :one 
SELECT * FROM products 
WHERE id = $1; 

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2; 