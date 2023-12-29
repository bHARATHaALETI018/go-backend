-- name: CreateStudent :one
INSERT INTO student(
    email,
    password,
    user_name,
    first_name,
    last_name,
    roll_number,
    stream,
    section,
    course,
    phone,
    mentor
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 
)
RETURNING *;


-- name: GetStudent :one
SELECT * FROM student
WHERE roll_number = $1 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM student
ORDER BY roll_number
LIMIT $1
OFFSET $2; 

-- name: UpdateStudentMentor :one
UPDATE student SET mentor = $2, updated_at = now()
WHERE roll_number = $1
RETURNING *;

-- name: UpdateStudentCourse :one
UPDATE student SET course = $2, updated_at = now()
WHERE roll_number = $1
RETURNING *;

-- name: UpdateStudentStream :one
UPDATE student SET stream = $2, updated_at = now()
WHERE roll_number = $1
RETURNING *;

-- name: UpdateStudentSection :one
UPDATE student SET section = $2, updated_at = now()
WHERE roll_number = $1
RETURNING *;

-- name: UpdateStudentPhone :one
UPDATE student SET phone = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteStudent :one
DELETE FROM student WHERE roll_number = $1
RETURNING *;