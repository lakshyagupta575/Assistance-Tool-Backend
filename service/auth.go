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

	// Checking if username is already taken or not
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
	_, err = database.Exec(fmt.Sprintf(db.StoreUserDetails,string(registerData.Username),string(registerData.Password)))
	db.ErrorCheck(err)

	w.WriteHeader(http.StatusOK)
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

	var loginData db.Login
	err := json.NewDecoder(req.Body).Decode(&loginData)
	db.ErrorCheck(err)

	database, err := db.ConnStageLocalDB()
	db.ErrorCheck(err)
	defer database.Close()

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	// checking if the user exists or not
	rows, err := database.Query(fmt.Sprintf(db.CheckUser,string(loginData.Username)))
	db.ErrorCheck(err)
	if(!rows.Next()) {
		w.WriteHeader(http.StatusUnauthorized)
		resp["message"] = "User does not exist"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	} 

	// if the user exists, verifying the entered id/pass combination
	rows, err = database.Query(fmt.Sprintf(db.VerifyUserCredentials,string(loginData.Username),string(loginData.Password)))
	db.ErrorCheck(err)
	if(!rows.Next()) {
			w.WriteHeader(http.StatusUnauthorized)
			resp["message"] = "Incorrect password"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
			return
	} 

	w.WriteHeader(http.StatusOK)
	resp["message"] = "User logged in successfully"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)	
}