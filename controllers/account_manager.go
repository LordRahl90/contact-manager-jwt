package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LordRahl90/contacts_manager/models"
	"github.com/LordRahl90/contacts_manager/utils"
)

//CreateAccount function to create a new account detail.
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	response := account.Create()
	utils.Respond(w, response)
}

//Authenticate Process
func Authenticate(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	response := models.Login(account.Email, account.Password)
	utils.Respond(w, response)
}
