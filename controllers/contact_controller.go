package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LordRahl90/contacts_manager/utils"

	"github.com/LordRahl90/contacts_manager/models"
)

//CreateContact Handler/
var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	user := r.Context().Value("user").(uint)

	fmt.Println(user)

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Contact Format"))
		return
	}

	contact.UserID = user
	response := contact.Create()

	utils.Respond(w, response)
}

//GetUserContact function to Retrieve the contacts for a user.
var GetUserContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	data := models.GetContacts(user)
	response := utils.Message(true, "Contacts Retrieved Successfully")
	response["data"] = data
	utils.Respond(w, response)
}
