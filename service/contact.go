package service

import (
	"log"
	"net/http"

	"assistanceTool.com/panelbackend/db"
)

func ContactHandler(w http.ResponseWriter, req *http.Request) {
	db.AppendHeaders(w,req)
	log.Printf("Contact Backend API Hit!")
}