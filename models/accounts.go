package models

import (
	"os"
	"strings"

	"github.com/LordRahl90/contacts_manager/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Token struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//Account struct
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Token    string `json:"token";sql:"-"`
}

//Validate function to validate model.
func (account *Account) Validate() (response map[string]interface{}, status bool) {
	db := GetDb()
	if !strings.Contains(account.Email, "@") {
		//email doesnt exist in a string
		return utils.Message(false, "Email Address is required."), false
	}

	if account.Password == "" {
		return utils.Message(false, "Password must be provided"), false
	}

	if len(account.Password) < 6 {
		return utils.Message(false, "Password Must be greater than 6"), false
	}

	var temp = &Account{}
	db.Where("email=?", account.Email).First(temp)
	if temp.Email != "" {
		//means email exists already
		return utils.Message(false, "Account has been registered before."), false
	}

	return utils.Message(false, "Requirement passed"), true
}

//Create function to pop up a new account.
func (account *Account) Create() map[string]interface{} {
	db := GetDb()
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	db.Create(&account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account. connection error")
	}

	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = "" //delete password.
	response := utils.Message(true, "Account Created Successfully!")
	response["account"] = account
	return response
}

//Login and generate a token
func Login(email, password string) (response map[string]interface{}) {
	println(email)
	db := GetDb()
	account := Account{}
	temp := []Account{}

	db.Where("email=?", email).Find(&temp)
	if len(temp) <= 0 {
		return utils.Message(false, "Email Not found")
	}
	account = temp[0]

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //password does not match
		return utils.Message(false, "Invalid Login Credentials")
	}

	account.Password = "" //delete password.
	//create a token.
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := utils.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

//GetUser function to Get a user's account
func GetUser(user uint) (account *Account) {
	db := GetDb()
	db.Table("accounts").Where("id=?", user).First(account)
	if account.Email == "" {
		return nil
	}

	account.Password = "" //password deleted.
	return account
}
