package models

import "github.com/rudysacosta/go-and-mysql-dockerizing/db"

type Model interface {
}

var dbConnection *db.Database

func SetDB(database *db.Database) {
	dbConnection = database
}

// Model defines the methods for CRUD (Create, Read, Update, Delete)
// operations on a data model.
type ReadableModel interface {
	List() ([]Model, error)
	Get(id int) (Model, error)
}

// WriteModel defines methods for writing operations on a data model.
type WritableModel interface {
	// Create adds a new item to the data model.
	Create(item Model) error

	// Update updates an existing item in the data model.
	Update(id int, item Model) error

	// Delete removes an item from the data model by its ID.
	Delete(id int) error
}
