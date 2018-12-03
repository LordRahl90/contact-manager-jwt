package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

//Message function to return the json messages
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond function to encode the data in json.
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//HandleError function to handle generic errors.
func HandleError(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Fatal(err)
	}
}
