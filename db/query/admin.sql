-- name: CreateAdmin :one
INSERT INTO admin (
  email,
  password,
  user_name,
  first_name,  
  last_name,
  id_number,
  phone
) VALUES (
  $1, $2, $3, $4, $5, $6, $7 
)
RETURNING *;

-- name: GetAdmin :one
SELECT * FROM admin
WHERE id_number = $1 LIMIT 1;

-- name: ListAdmins :many
SELECT * FROM admin
ORDER BY id_number
LIMIT $1
OFFSET $2;

-- name: UpdateAdminPhone :one
UPDATE admin SET phone = $2, updated_at = now()
WHERE id_number = $1
RETURNING *;

-- name: DeleteAdmin :one
DELETE FROM admin WHERE id_number = $1
RETURNING *;
