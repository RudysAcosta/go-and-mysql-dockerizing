package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Rol model
type Rol struct {
	ID          uint      `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

const tableName = "roles"

type RolModel struct {
}

func (rolModel *RolModel) List() ([]Rol, error) {
	var roles []Rol
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := dbConnection.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var rol Rol
		if err := rows.Scan(&rol.ID, &rol.Name, &rol.Description); err != nil {
			return nil, err
		}

		roles = append(roles, rol)
	}
	return roles, nil
}

func (rolModel *RolModel) Get(id int) (Model, error) {
	var rol Rol
	// Make the sql with id and table name
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", tableName, id)

	row := dbConnection.QueryRow(query)

	// Scan the result of the query into the Rol struct
	err := row.Scan(&rol.ID, &rol.Name, &rol.Description)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("Rol not found")
		}

		return nil, err
	}

	return rol, nil
}
