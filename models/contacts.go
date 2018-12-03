package models

import (
	"strings"

	"github.com/LordRahl90/contacts_manager/utils"
	"github.com/jinzhu/gorm"
)

//Contact Struct
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	UserID uint   `json:"user_id"` //relationship to the accounts table.
}

//Validate the contact
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	db := GetDb()

	if contact.Name == "" {
		return utils.Message(false, "Please provide a name"), false
	}

	if contact.Phone == "" {
		return utils.Message(false, "Provide a Phone Number"), false
	}

	if contact.Email == "" {
		return utils.Message(false, "Provide a valid email"), false
	}

	if !strings.Contains(contact.Email, "@") {
		return utils.Message(false, "Invaid Email Address Provided"), false
	}

	if contact.UserID < 0 {
		return utils.Message(false, "Invalid User"), false
	}

	//make sure the user doesnt save the contact twice.
	var temp = Contact{}
	db.Where("user_id=? and email=?", contact.UserID, contact.Email).First(&temp)
	if temp.Email != "" {
		return utils.Message(false, "This Contact have been saved before."), false
	}

	return utils.Message(true, "success"), true
}

//Create new contact
func (contact Contact) Create() map[string]interface{} {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	//we proceed to create account.
	db.Create(&contact)
	if contact.ID <= 0 {
		return utils.Message(false, "An error occurred while creating the contact")
	}
	response := utils.Message(true, "Contact Created Successfully.")
	response["contact"] = contact
	return response
}

//GetContact - function to get a single contact.
func GetContact(id uint) (contact *Contact) {
	db.Where("id=?", id).First(&contact)
	return
}

//GetContacts Get all the contacts for a user.
func GetContacts(user uint) (contacts []*Contact) {
	db := GetDb()
	db.Where("user_id=?", user).Find(&contacts)

	return
}
