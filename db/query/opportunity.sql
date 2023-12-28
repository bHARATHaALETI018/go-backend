-- name: CreateOpportunity :one
INSERT INTO opportunity(
    title,
    link,
    status,
    created_by,
    approved_by
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;


-- name: GetOpportunity :one
SELECT * FROM opportunity
WHERE id = $1 LIMIT 1;

-- name: ListOpportunitys :many
SELECT * FROM opportunity
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOpportunityStatus :one
UPDATE opportunity SET status = $2
WHERE id = $1
RETURNING *;

-- name: UpdateOpportunityLink :one
UPDATE opportunity SET link = $2
WHERE id = $1
RETURNING *;
 
-- name: UpdateOpportunityApprovedBy :one
UPDATE opportunity SET approved_by = $2
WHERE id = $1
RETURNING *;
 

-- name: DeleteOpportunity :one
DELETE FROM opportunity WHERE id = $1
RETURNING *;
