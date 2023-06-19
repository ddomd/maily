package api

import (
	"time"

	"github.com/ddomd/maily/internal/mdb"
)

type Email struct {
	ID          int64      `json:"id"`
	Address       string   `json:"email"`
	ConfirmedAt time.Time  `json:"confirmed_at"`
	OptOut      bool       `json:"opt_out"`
}

type Emails struct {
	Emails []Email `json:"emails"`
}

func dbEmailToJson(dbEmail mdb.Email) Email{
	return Email{
		dbEmail.ID,
		dbEmail.Address,
		dbEmail.ConfirmedAt,
		dbEmail.OptOut,
	}
}

func dbEmailsToJson(dbEmails []mdb.Email) Emails {
	emails := []Email{}

	for _, dbEmail := range dbEmails {
		emails = append(emails, dbEmailToJson(dbEmail))
	}

	return Emails{emails}
}
