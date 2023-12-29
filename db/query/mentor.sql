-- name: CreateMentor :one
INSERT INTO mentor(
    email,
    password,
    user_name,
    first_name,  
    last_name,
    phone,
    id_number
) VALUES (
    $1, $2, $3, $4, $5, $6, $7 
)
RETURNING *;


-- name: GetMentor :one
SELECT * FROM mentor
WHERE id_number = $1 LIMIT 1;

-- name: ListMentors :many
SELECT * FROM mentor
ORDER BY id_number
LIMIT $1
OFFSET $2; 

-- name: UpdateMentorPhone :one
UPDATE mentor SET phone = $2, updated_at = now()
WHERE id_number = $1
RETURNING *;

-- name: DeleteMentor :one
DELETE FROM mentor WHERE id_number = $1
RETURNING *;