package response

import (
	"encoding/json"
	"net/http"
)

type Success struct{
	Status 		string       	`json:"status"`
	Data 		interface{}  	`json:"data"`
}

type Error struct{
	Status    	string     		`json:"status"`
	Message		string  		`json:"message"`
}


func JSON(w http.ResponseWriter, status int, data interface{}){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Success{
		Status: "success",
		Data: data,
	})

}

func JSONError(w http.ResponseWriter, status int, message string){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Error{
		Status: "error",
		Message: message,
	})
}