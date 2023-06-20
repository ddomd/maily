package mdb

import (
	"database/sql"
	"log"
	"time"
)

func (mdb *MDB) CreateEmail(address string) (Email, error){
	currentTime := time.Now().Unix()

	result := mdb.DB.QueryRow(`
		INSERT INTO 
		emails(email, confirmed_at, opt_out)
		VALUES(?, ?, false)
		RETURNING *;
	`, address, currentTime)
	
	createdEmail, err := scanRow(result)
	if err != nil {
		return Email{}, err
	}

	return createdEmail, nil
}

func (mdb *MDB) GetEmail(id int64) (Email, error) {
	emailRow := mdb.DB.QueryRow(`
		SELECT *
		FROM emails
		WHERE id = ?;
	`, id)

	foundEmail, err := scanRow(emailRow)
	if err != nil {
		return Email{}, err
	}

	return foundEmail, nil
}

func (mdb *MDB) GetBatchEmails(params BatchParams) ([]Email, error) {
	emails := make([]Email, 0)

	emailRows, err := mdb.DB.Query(`
		SELECT * FROM emails
		WHERE opt_out=false
		ORDER BY id ASC
		LIMIT ? OFFSET ?;
	`, params.Limit, (params.Offset-1)*params.Limit)
	if err != nil {
		return nil, err
	}

	defer emailRows.Close()

	for emailRows.Next() {
		email, err := scanRows(emailRows)
		if err != nil {
			return nil, err
		}

		emails = append(emails, email)
	}

	return emails, nil
}

func (mdb *MDB) UpdateEmail(id int64, optOut bool) (Email, error) {
	currentTime := time.Now().Unix()

	result := mdb.DB.QueryRow(`
		UPDATE emails
		SET confirmed_at=?,
				opt_out=?
		WHERE id=?
		RETURNING *;
	`, currentTime, optOut, id)
	
	updatedEmail, err := scanRow(result)
	if err != nil {
		return Email{}, err
	}

	return updatedEmail, nil
}

func (mdb *MDB) DeleteEmail(id int64) (Email, error) {
	result := mdb.DB.QueryRow(`
		DELETE FROM emails
		WHERE id=?
		RETURNING *;
	`, id)
	
	deletedEmail, err := scanRow(result)
	if err != nil {
		return Email{}, err
	}

	return deletedEmail, nil
}

func (mdb *MDB) CreateEmailTable() {
	_, err := mdb.DB.Exec(`
		CREATE TABLE IF NOT EXISTS emails (
			id INTEGER   PRIMARY KEY,
			email TEXT   UNIQUE NOT NULL,
			confirmed_at INTEGER NOT NULL,
			opt_out      INTEGER NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

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
