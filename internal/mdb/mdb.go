package mdb

import (
	"database/sql"
)

type MDB struct {
	DB *sql.DB
}

func NewMdb(db *sql.DB) *MDB {
	return &MDB{db}
}
