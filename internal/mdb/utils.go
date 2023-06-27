package mdb

import (
	"database/sql"
	"errors"
	"time"
)

func scanRows(row *sql.Rows) (Email, error) {
	var email Email
	var confirmedAt int64

	err := row.Scan(&email.ID, &email.Address, &confirmedAt, &email.OptOut)
	if err != nil {
		return Email{}, err
	}

	confirmedTime := time.Unix(confirmedAt, 0)
	email.ConfirmedAt = confirmedTime

	return email, nil
}

func scanRow(row *sql.Row) (Email, error) {
	var email Email
	var confirmedAt int64

	err := row.Scan(&email.ID, &email.Address, &confirmedAt, &email.OptOut)
	if err != nil {
		return Email{}, err
	}

	confirmedTime := time.Unix(confirmedAt, 0)
	email.ConfirmedAt = confirmedTime

	return email, nil
}

func checkParams(params BatchParams) error {
	if params.Offset <= 0 || params.Limit <= 0 {
		return errors.New("Invalid param range")
	}
	return  nil
}
