package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrEditConflict        = errors.New("edit conflict")
	ErrDuplicateEmail      = errors.New("duplicate email")
	ErrDuplicateCarCatalog = errors.New("exact duplicate row: This car catalog already exists")
	// Shop represents a collection of constants related to the company's shop.
	Shop = struct {
		HEADQUARTERS string
	}{
		HEADQUARTERS: "HEADQUARTERS",
	}
)

// Models represents a collection of models for interacting with the database.
type Models struct {
	UserModel        *UserModel
	ClientModel      *ClientModel
	CarProductModel  *CarProductModel
	CarCatalogModel  *CarCatalogModel
	TransactionModel *TransactionModel
	SearchModel      *SearchModel
}

// NewModels creates a new instance of Models with the provided database connection.
//
// Parameters:
//
//	db - The SQL database connection
//
// Returns:
//
//	Models - A new instance of Models
func NewModels(db *sql.DB) Models {
	return Models{
		UserModel:        &UserModel{db},
		ClientModel:      &ClientModel{db},
		CarProductModel:  &CarProductModel{db},
		CarCatalogModel:  &CarCatalogModel{db},
		TransactionModel: &TransactionModel{db},
		SearchModel:      &SearchModel{db},
	}
}
