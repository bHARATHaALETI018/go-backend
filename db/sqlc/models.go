// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Statuses string

const (
	StatusesApproved           Statuses = "approved"
	StatusesWaitingforapproval Statuses = "waiting for approval"
	StatusesRejected           Statuses = "rejected"
)

func (e *Statuses) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Statuses(s)
	case string:
		*e = Statuses(s)
	default:
		return fmt.Errorf("unsupported scan type for Statuses: %T", src)
	}
	return nil
}

type NullStatuses struct {
	Statuses Statuses `json:"statuses"`
	Valid    bool     `json:"valid"` // Valid is true if Statuses is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatuses) Scan(value interface{}) error {
	if value == nil {
		ns.Statuses, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Statuses.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatuses) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Statuses), nil
}

type Admin struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// teacher id number from the id card
	IDNumber  string       `json:"id_number"`
	Phone     string       `json:"phone"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Mentor struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	// teacher id number from the id card
	IDNumber  string       `json:"id_number"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Opportunity struct {
	ID         int64         `json:"id"`
	Title      string        `json:"title"`
	Link       string        `json:"link"`
	Status     NullStatuses  `json:"status"`
	CreatedBy  int64         `json:"created_by"`
	ApprovedBy sql.NullInt64 `json:"approved_by"`
	CreatedAt  sql.NullTime  `json:"created_at"`
	UpdatedAt  sql.NullTime  `json:"updated_at"`
}

type OpportunityStatusHistory struct {
	ID             int64         `json:"id"`
	OpportunityID  sql.NullInt64 `json:"opportunity_id"`
	PreviousStatus Statuses      `json:"previous_status"`
	Status         Statuses      `json:"status"`
	CreatedAt      sql.NullTime  `json:"created_at"`
	UpdatedAt      sql.NullTime  `json:"updated_at"`
}

type Student struct {
	ID         int64         `json:"id"`
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	UserName   string        `json:"user_name"`
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	RollNumber string        `json:"roll_number"`
	Stream     string        `json:"stream"`
	Section    string        `json:"section"`
	Course     string        `json:"course"`
	Phone      string        `json:"phone"`
	Mentor     sql.NullInt64 `json:"mentor"`
	CreatedAt  sql.NullTime  `json:"created_at"`
	UpdatedAt  sql.NullTime  `json:"updated_at"`
}
