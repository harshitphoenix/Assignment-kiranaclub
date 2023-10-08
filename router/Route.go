package router

import (
	"io"
	"net/http"

	"gorm.io/gorm"

	"server/controllers"

	"github.com/gorilla/mux"
)

func Handlers(Db *gorm.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	h := controllers.NewUtilHandler(Db)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Health check done!!!")
	}).Methods("GET")
	router.HandleFunc("/createJob", h.SubmitJob).Methods("POST")
	router.HandleFunc("/getJobStatus", h.GetJobStatus).Methods("GET")
	router.HandleFunc("/getVisitInfo", h.GetVisitInfo).Methods("GET")
	router.HandleFunc("/transferData", h.TransferCSVtoDb).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "404 Route Not Found")
	})

	return router
}
