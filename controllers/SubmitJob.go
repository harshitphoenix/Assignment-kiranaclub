package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "image/jpeg"
	"server/models"
)

type Visit struct {
	StoreID   string    `json:"store_id"`
	ImageURLs []string  `json:"image_url"`
	VisitTime time.Time `json:"visit_time"`
}
type JobRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

type SubmitJobResponse struct {
	JobsIds []int `json:"jobs_ids"`
}

func (h *UtilHandler) SubmitJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubmitJob")
	decoder := json.NewDecoder(r.Body)
	var t JobRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	if t.Count == 0 || len(t.Visits) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "count or visit length is 0"}`))
		return
	}
	if t.Count != len(t.Visits) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "count and visit length not match"}`))
		return
	}
	var jobsIds []int
	for _, v := range t.Visits {
		job := models.Job{
			StoreId:   v.StoreID,
			CreatedAt: time.Now(),
			Status:    "pending",
		}
		h.db.Create(&job)
		jobsIds = append(jobsIds, job.ID)

		go ProcessJob(h.db, job.ID, v)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&SubmitJobResponse{jobsIds})
}
