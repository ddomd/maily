package mdb

import "time"

type Email struct {
	ID          int64
	Address       string
	ConfirmedAt time.Time
	OptOut      bool
}

type BatchParams struct {
	Offset  int
	Limit   int
}
