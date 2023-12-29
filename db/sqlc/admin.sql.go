// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: admin.sql

package db

import (
	"context"
)

const createAdmin = `-- name: CreateAdmin :one
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
RETURNING id, email, password, user_name, first_name, last_name, id_number, phone, created_at, updated_at
`

type CreateAdminParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IDNumber  string `json:"id_number"`
	Phone     string `json:"phone"`
}

func (q *Queries) CreateAdmin(ctx context.Context, arg CreateAdminParams) (Admin, error) {
	row := q.db.QueryRowContext(ctx, createAdmin,
		arg.Email,
		arg.Password,
		arg.UserName,
		arg.FirstName,
		arg.LastName,
		arg.IDNumber,
		arg.Phone,
	)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.UserName,
		&i.FirstName,
		&i.LastName,
		&i.IDNumber,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAdmin = `-- name: DeleteAdmin :one
DELETE FROM admin WHERE id_number = $1
RETURNING id, email, password, user_name, first_name, last_name, id_number, phone, created_at, updated_at
`

func (q *Queries) DeleteAdmin(ctx context.Context, idNumber string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, deleteAdmin, idNumber)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.UserName,
		&i.FirstName,
		&i.LastName,
		&i.IDNumber,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAdmin = `-- name: GetAdmin :one
SELECT id, email, password, user_name, first_name, last_name, id_number, phone, created_at, updated_at FROM admin
WHERE id_number = $1 LIMIT 1
`

func (q *Queries) GetAdmin(ctx context.Context, idNumber string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdmin, idNumber)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.UserName,
		&i.FirstName,
		&i.LastName,
		&i.IDNumber,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAdmins = `-- name: ListAdmins :many
SELECT id, email, password, user_name, first_name, last_name, id_number, phone, created_at, updated_at FROM admin
ORDER BY id_number
LIMIT $1
OFFSET $2
`

type ListAdminsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAdmins(ctx context.Context, arg ListAdminsParams) ([]Admin, error) {
	rows, err := q.db.QueryContext(ctx, listAdmins, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Admin
	for rows.Next() {
		var i Admin
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.UserName,
			&i.FirstName,
			&i.LastName,
			&i.IDNumber,
			&i.Phone,
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

const updateAdminPhone = `-- name: UpdateAdminPhone :one
UPDATE admin SET phone = $2, updated_at = now()
WHERE id_number = $1
RETURNING id, email, password, user_name, first_name, last_name, id_number, phone, created_at, updated_at
`

type UpdateAdminPhoneParams struct {
	IDNumber string `json:"id_number"`
	Phone    string `json:"phone"`
}

func (q *Queries) UpdateAdminPhone(ctx context.Context, arg UpdateAdminPhoneParams) (Admin, error) {
	row := q.db.QueryRowContext(ctx, updateAdminPhone, arg.IDNumber, arg.Phone)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.UserName,
		&i.FirstName,
		&i.LastName,
		&i.IDNumber,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
