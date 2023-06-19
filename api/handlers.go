package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ddomd/maily/internal/mdb"
	"github.com/ddomd/maily/utils"
	"github.com/go-chi/chi/v5"
)

func (s *Server) HandleCreateEmail(rw http.ResponseWriter, req *http.Request) {
	type Params struct {
		EmailAddress string `json:"email_address"`
	}

	params := Params{}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, "Couldn't decode params")
		return
	}

	email, err := s.DB.CreateEmail(params.EmailAddress)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, "Couldn't create Email")
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Created email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	utils.RespondWithJson(rw, http.StatusOK, res)
}

func (s *Server) HandleGetBatchEmail(rw http.ResponseWriter, req *http.Request) {
	limitParam := chi.URLParam(req, "limit")
	offsetParam := chi.URLParam(req, "offset")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		utils.RespondWithError(rw, http.StatusBadRequest, "Limit url param is not an int")
		return
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		utils.RespondWithError(rw, http.StatusBadRequest, "Offset url param is not an int")
		return
	}

	emails, err := s.DB.GetBatchEmails(mdb.BatchParams{
		Offset: offset, 
		Limit: limit,
	})
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, "Couldn't create Email")
		return
	}

	res := dbEmailsToJson(emails)

	log.Printf("REST: Retrieved email batch(id:%d-id:%d), client Addr: %s\n",
		res.Emails[0].ID,
		res.Emails[len(res.Emails)-1].ID,
		req.RemoteAddr,
	)
	utils.RespondWithJson(rw, http.StatusOK, res)
}

func (s *Server) HandleGetEmail(rw http.ResponseWriter, req *http.Request) {
	type Params struct {
		EmailAddress string `json:"email_address"`
	}

	params := Params{}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, "Couldn't decode params")
		return
	}

	email, err := s.DB.GetEmail(params.EmailAddress)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, "Entry not found")
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Retrieved email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	utils.RespondWithJson(rw, http.StatusOK, res)
}

func (s *Server) HandleUpdateEmail(rw http.ResponseWriter, req *http.Request) {
	type Params struct {
		EmailAddress       string     `json:"email_address"`
		OptOut      bool       `json:"opt_out"`
	}

	params := Params{}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, "Couldn't decode params")
		return
	}

	email, err := s.DB.UpdateEmail(params.EmailAddress, params.OptOut)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Updated email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	utils.RespondWithJson(rw, http.StatusOK, res)
}

func (s *Server) HandleDeleteEmail(rw http.ResponseWriter, req *http.Request) {
	type Params struct {
		EmailAddress string `json:"email_address"`
	}

	params := Params{}

	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, "Couldn't decode params")
		return
	}

	email, err := s.DB.DeleteEmail(params.EmailAddress)
	if err != nil {
		utils.RespondWithError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	res := dbEmailToJson(email)

	log.Printf("REST: Deleted email entry(id:%d), client Addr: %s\n", res.ID, req.RemoteAddr)
	utils.RespondWithJson(rw, http.StatusOK, res)
}
