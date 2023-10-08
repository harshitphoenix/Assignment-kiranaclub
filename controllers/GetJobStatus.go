package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
)

type ErrorMessage struct {
	StoreId string `json:"store_id"`
	Message string `json:"message"`
}
type Response struct {
	Status string       `json:"status"`
	JobId  int          `json:"job_id"`
	Error  ErrorMessage `json:"error"`
}

func (h *UtilHandler) GetJobStatus(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Job Status")
	queryParams := r.URL.Query().Get("job_id")

	if len(queryParams) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "job_id is missing"}`))
		return
	}
	res := Response{}
	var job models.Job

	result := h.db.Where("id = ?", queryParams).Find(&job)

	if result.Error != nil {
		fmt.Println("result.Error", result.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "job_id not found"}`))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		if job.ID == 0 {
			res.Status = ""
			res.JobId = 0
			res.Error = ErrorMessage{
				StoreId: "",
				Message: "job_id not found",
			}
		} else {
			res.Status = job.Status
			res.JobId = job.ID
			if job.Status == "failed" {
				if job.StoreId == "" {
					res.Error = ErrorMessage{
						StoreId: "",
						Message: "Store not found",
					}

				} else {
					res.Error = ErrorMessage{
						StoreId: job.StoreId,
						Message: "Image Processing Failed",
					}

				}

			}
		}

	}
	json.NewEncoder(w).Encode(res)
}
