package models

import (
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

// var tableName string = "roles"
func (rol *RolModel) List() ([]Rol, error) {
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

// func (rol *RolModel) Get(id int) (Model, error) {
// 	var myRol Rol
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", tableName, id)

// 	row := Rol{}

// 	err = dbConnection.Get(&row, query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer row.Close()

// 	fmt.Printf("\n\nGet rol: %v\n\n", row)

// 	fmt.Println("RolModel Get")
// 	return myRol, nil
// }
