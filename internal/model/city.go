package model

import (
	"context"
	"database/sql"
)

// City struct
type City struct {
	ID   int32
	Name string
}

// List of City
func (u *City) List(ctx context.Context, db *sql.DB) ([]City, error) {
	var list []City
	var err error
	query := `SELECT id, name FROM cities`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var city City
		err = rows.Scan(&city.ID, &city.Name)
		if err != nil {
			return list, err
		}

		list = append(list, city)
	}

	return list, rows.Err()
}

// Get City
func (u *City) Get(ctx context.Context, db *sql.DB) error {
	return db.QueryRowContext(ctx, `SELECT id, name FROM cities WHERE id = $1`, u.ID).Scan(&u.ID, &u.Name)
}

// Create New City
func (u *City) Create(ctx context.Context, db *sql.DB) error {

	stmt, err := db.PrepareContext(ctx, `INSERT INTO cities (name) VALUES ($1) RETURNING id`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, u.Name).Scan(&u.ID)
}
