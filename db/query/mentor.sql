-- name: CreateMentor :one
INSERT INTO mentor(
    email,
    password,
    user_name,
    first_name,
    middle_name,
    last_name,
    phone,
    id_number
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
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

-- name: UpdateMentorMiddleName :one
UPDATE mentor SET middle_name = $2
WHERE id_number = $1
RETURNING *;

-- name: UpdateMentorPhone :one
UPDATE mentor SET phone = $2
WHERE id_number = $1
RETURNING *;

-- name: DeleteMentor :one
DELETE FROM mentor WHERE id_number = $1
RETURNING *;