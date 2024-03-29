// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: opportunity.sql

package db

import (
	"context"
	"database/sql"
)

const createOpportunity = `-- name: CreateOpportunity :one
INSERT INTO opportunity(
    title,
    link, 
    created_by,
    approved_by
) VALUES (
    $1, $2, $3, $4 
)
RETURNING id, title, link, status, created_by, approved_by, created_at, updated_at
`

type CreateOpportunityParams struct {
	Title      string        `json:"title"`
	Link       string        `json:"link"`
	CreatedBy  int64         `json:"created_by"`
	ApprovedBy sql.NullInt64 `json:"approved_by"`
}

func (q *Queries) CreateOpportunity(ctx context.Context, arg CreateOpportunityParams) (Opportunity, error) {
	row := q.db.QueryRowContext(ctx, createOpportunity,
		arg.Title,
		arg.Link,
		arg.CreatedBy,
		arg.ApprovedBy,
	)
	var i Opportunity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Link,
		&i.Status,
		&i.CreatedBy,
		&i.ApprovedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOpportunity = `-- name: DeleteOpportunity :one
DELETE FROM opportunity WHERE id = $1
RETURNING id, title, link, status, created_by, approved_by, created_at, updated_at
`

func (q *Queries) DeleteOpportunity(ctx context.Context, id int64) (Opportunity, error) {
	row := q.db.QueryRowContext(ctx, deleteOpportunity, id)
	var i Opportunity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Link,
		&i.Status,
		&i.CreatedBy,
		&i.ApprovedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOpportunity = `-- name: GetOpportunity :one
SELECT id, title, link, status, created_by, approved_by, created_at, updated_at FROM opportunity
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOpportunity(ctx context.Context, id int64) (Opportunity, error) {
	row := q.db.QueryRowContext(ctx, getOpportunity, id)
	var i Opportunity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Link,
		&i.Status,
		&i.CreatedBy,
		&i.ApprovedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOpportunitys = `-- name: ListOpportunitys :many
SELECT id, title, link, status, created_by, approved_by, created_at, updated_at FROM opportunity
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListOpportunitysParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOpportunitys(ctx context.Context, arg ListOpportunitysParams) ([]Opportunity, error) {
	rows, err := q.db.QueryContext(ctx, listOpportunitys, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Opportunity
	for rows.Next() {
		var i Opportunity
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Link,
			&i.Status,
			&i.CreatedBy,
			&i.ApprovedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOpportunityApprovedBy = `-- name: UpdateOpportunityApprovedBy :one
UPDATE opportunity SET approved_by = $2, updated_at = now()
WHERE id = $1
RETURNING id, title, link, status, created_by, approved_by, created_at, updated_at
`

type UpdateOpportunityApprovedByParams struct {
	ID         int64         `json:"id"`
	ApprovedBy sql.NullInt64 `json:"approved_by"`
}

func (q *Queries) UpdateOpportunityApprovedBy(ctx context.Context, arg UpdateOpportunityApprovedByParams) (Opportunity, error) {
	row := q.db.QueryRowContext(ctx, updateOpportunityApprovedBy, arg.ID, arg.ApprovedBy)
	var i Opportunity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Link,
		&i.Status,
		&i.CreatedBy,
		&i.ApprovedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateOpportunityLink = `-- name: UpdateOpportunityLink :one
UPDATE opportunity SET link = $2, updated_at = now()
WHERE id = $1
RETURNING id, title, link, status, created_by, approved_by, created_at, updated_at
`

type UpdateOpportunityLinkParams struct {
	ID   int64  `json:"id"`
	Link string `json:"link"`
}

func (q *Queries) UpdateOpportunityLink(ctx context.Context, arg UpdateOpportunityLinkParams) (Opportunity, error) {
	row := q.db.QueryRowContext(ctx, updateOpportunityLink, arg.ID, arg.Link)
	var i Opportunity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Link,
		&i.Status,
		&i.CreatedBy,
		&i.ApprovedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateOpportunityStatus = `-- name: UpdateOpportunityStatus :one
UPDATE opportunity SET status = $2, updated_at = now()
WHERE id = $1
RETURNING id, title, link, status, created_by, approved_by, created_at, updated_at
`

type UpdateOpportunityStatusParams struct {
	ID     int64        `json:"id"`
	Status NullStatuses `json:"status"`
}

func (q *Queries) UpdateOpportunityStatus(ctx context.Context, arg UpdateOpportunityStatusParams) (Opportunity, error) {
	row := q.db.QueryRowContext(ctx, updateOpportunityStatus, arg.ID, arg.Status)
	var i Opportunity
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Link,
		&i.Status,
		&i.CreatedBy,
		&i.ApprovedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
