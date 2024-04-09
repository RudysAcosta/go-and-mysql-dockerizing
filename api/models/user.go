package models

import "time"

// User model
type User struct {
	ID              uint       `db:"id"`
	RoleID          uint       `db:"role_id"`
	FirstName       string     `db:"first_name"`
	LastName        string     `db:"last_name"`
	Email           string     `db:"email"`
	EmailConfirmed  bool       `db:"email_confirmed"`
	Password        string     `db:"password"`
	LastLogin       *time.Time `db:"last_login"`
	Created         time.Time  `db:"created"`
	CreatedBy       string     `db:"created_by"`
	Updated         time.Time  `db:"updated"`
	UpdatedBy       string     `db:"updated_by"`
	Telephone       string     `db:"telephone"`
	Status          string     `db:"status"`
	StatusUpdated   *time.Time `db:"status_updated"`
	StatusUpdatedBy string     `db:"status_updated_by"`
}
