package controllers

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type ControllerInterface interface {
	SubmitJob(w http.ResponseWriter, r *http.Request)
	GetJobStatus(w http.ResponseWriter, r *http.Request)
	GetVisitInfo(w http.ResponseWriter, r *http.Request)
	TransferCSVtoDb(w http.ResponseWriter, r *http.Request)
}
type UtilHandler struct {
	db *gorm.DB
}

func NewUtilHandler(db *gorm.DB) ControllerInterface {
	fmt.Println("New Connection")
	return &UtilHandler{
		db: db,
	}
}
