package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ddomd/maily/internal/mdb"
	"github.com/go-chi/chi/v5"
)

//Creates a single email entry accepts a json payload containing an email address
func (s *Server) HandleCreateEmail(rw http.ResponseWriter, req *http.Request) {
	type Params struct {
		EmailAddress string `json:"email_address"`
	}

	params := Params{}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't decode params")
		return
	}

	email, err := s.DB.CreateEmail(params.EmailAddress)
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't create Email")
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Created email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	RespondWithJson(rw, http.StatusOK, res)
}

//Retrieves a single email entry by id passed via url params
func (s *Server) HandleGetEmail(rw http.ResponseWriter, req *http.Request) {
	idParam := chi.URLParam(req, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		RespondWithError(rw, http.StatusBadRequest, "id is not a number")
		return
	}

	email, err := s.DB.GetEmail(int64(id))
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Entry not found")
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Retrieved email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	RespondWithJson(rw, http.StatusOK, res)
}

//Retrieves all email entries in a single call
func (s *Server) HandleGetAllEmails(rw http.ResponseWriter, req *http.Request) {

	emails, err := s.DB.GetAllEmails()
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't create Email")
		return
	}

	res := dbEmailsToJson(emails)

	log.Printf("REST: Retrieved all emails, client Addr: %s\n", req.RemoteAddr)
	RespondWithJson(rw, http.StatusOK, res)
}

func (s *Server) HandleGetAllSubscribed(rw http.ResponseWriter, req *http.Request) {

	emails, err := s.DB.GetAllSubscribed()
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't create Email")
		return
	}

	res := dbEmailsToJson(emails)

	log.Printf("REST: Retrieved all emails, client Addr: %s\n", req.RemoteAddr)

	RespondWithJson(rw, http.StatusOK, res)
}

//Retrieves a batch of email entries using offset pagination limit and offset are specified via url param
func (s *Server) HandleGetBatchEmail(rw http.ResponseWriter, req *http.Request) {
	limitParam := chi.URLParam(req, "limit")
	offsetParam := chi.URLParam(req, "offset")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		RespondWithError(rw, http.StatusBadRequest, "Limit url param is not an int")
		return
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		RespondWithError(rw, http.StatusBadRequest, "Offset url param is not an int")
		return
	}

	emails, err := s.DB.GetBatchEmails(mdb.BatchParams{
		Offset: offset, 
		Limit: limit,
	})
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't create Email")
		return
	}

	res := dbEmailsToJson(emails)

	log.Printf("REST: Retrieved email batch(id:%d-id:%d), client Addr: %s\n",
		res.Emails[0].ID,
		res.Emails[len(res.Emails)-1].ID,
		req.RemoteAddr,
	)
	RespondWithJson(rw, http.StatusOK, res)
}

//Retrieves a batch of subscribed email entries using offset pagination limit and offset are passed via url params
func (s *Server) HandleGetBatchSubscribed(rw http.ResponseWriter, req *http.Request) {
	limitParam := chi.URLParam(req, "limit")
	offsetParam := chi.URLParam(req, "offset")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		RespondWithError(rw, http.StatusBadRequest, "Limit url param is not an int")
		return
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		RespondWithError(rw, http.StatusBadRequest, "Offset url param is not an int")
		return
	}

	emails, err := s.DB.GetBatchSubscribed(mdb.BatchParams{
		Offset: offset, 
		Limit: limit,
	})
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't retrieve emails")
		return
	}

	res := dbEmailsToJson(emails)

	log.Printf("REST: Retrieved subscribed email batch(id:%d-id:%d), client Addr: %s\n",
		res.Emails[0].ID,
		res.Emails[len(res.Emails)-1].ID,
		req.RemoteAddr,
	)

	RespondWithJson(rw, http.StatusOK, res)
}

//Updates a single email entry by id(passed via url param) and sets opt_out status(passed via json payload)
func (s *Server) HandleUpdateEmail(rw http.ResponseWriter, req *http.Request) {
	type Params struct {
		OptOut bool `json:"opt_out"`
	}

	idParam := chi.URLParam(req, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		RespondWithError(rw, http.StatusBadRequest, "id is not a number")
		return
	}

	params := Params{}

	err = json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't decode params")
		return
	}

	email, err := s.DB.UpdateEmail(int64(id), params.OptOut)
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Updated email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	RespondWithJson(rw, http.StatusOK, res)
}

//Deletes a single email entry by id(passed via url param)
func (s *Server) HandleDeleteEmail(rw http.ResponseWriter, req *http.Request) {
	idParam := chi.URLParam(req, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		RespondWithError(rw, http.StatusBadRequest, "id is not an int")
		return
	}

	email, err := s.DB.DeleteEmail(int64(id))
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Deleted email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	RespondWithJson(rw, http.StatusOK, res)
}

//Deletes all unsubscribed email entries
func (s *Server) HandleDeleteUnsubscribed(rw http.ResponseWriter, req *http.Request) {
	emails, err := s.DB.DeleteUnsubscribed()
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't retrieve emails")
		return
	}

	res := dbEmailsToJson(emails)

	log.Printf("REST: deleted all unsubscribed email entries, client Addr: %s\n",req.RemoteAddr)

	RespondWithJson(rw, http.StatusOK, res)
}

//Deletes all email entries unsubscribed before a date specified via json payload
func (s *Server) HandleDeleteUnsubscribedBefore(rw http.ResponseWriter, req *http.Request) {
	type Params struct {
		Date int64 `json:"date"`
	}

	params := Params{}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't decode params")
	}


	emails, err := s.DB.DeleteUnsubscribedBefore(params.Date)
	if err != nil {
		RespondWithError(rw, http.StatusInternalServerError, "Couldn't retrieve emails")
		return
	}

	res := dbEmailsToJson(emails)

	log.Printf("REST: deleted all unsubscribed email entries, client Addr: %s\n",req.RemoteAddr)

	RespondWithJson(rw, http.StatusOK, res)
}
