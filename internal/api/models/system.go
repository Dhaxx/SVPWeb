package models

import "database/sql"

type System struct {
	ID   int           `json:"id"`
	Name string        `json:"name"`
	Obs  sql.NullString `json:"obs"`
}