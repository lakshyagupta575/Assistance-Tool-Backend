package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"assistanceTool.com/panelbackend/db"
)

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	db.AppendHeaders(w, req)
	log.Printf("Register Backend API hit!")

	var registerData db.Register
	err := json.NewDecoder(req.Body).Decode(&registerData)
	db.ErrorCheck(err)

	database, err := db.ConnStageLocalDB()
	db.ErrorCheck(err)
	defer database.Close()

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	rows, err := database.Query(fmt.Sprintf(db.CheckUser,string(registerData.Username)))
	db.ErrorCheck(err)
	if(rows.Next()) {
		w.WriteHeader(http.StatusConflict)
		resp["message"] = "Username already exists"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		log.Printf("Error registering user: Username already exists!")
		return
	}

	_, err = database.Exec(fmt.Sprintf(db.RegisterUser, string(registerData.Username)))
	db.ErrorCheck(err)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	db.AppendHeaders(w, req)
	log.Printf("Login Backend API hit!")
}