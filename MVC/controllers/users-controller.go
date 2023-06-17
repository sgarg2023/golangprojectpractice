package controllers

import (
	"encoding/json"
	"net/http"
	"services"
	"strconv"
)

/*func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query().Get("user_id")
	log.Printf("About to process user ID %v", userId)
}*/

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.Write([]byte("user_id must be a a number"))
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
