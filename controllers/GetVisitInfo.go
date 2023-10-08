package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"time"
)

type UrlProcessInfo struct {
	Date      time.Time `json:"date"`
	Perimeter int       `json:"perimeter"`
}
type VisitInfo struct {
	StoreId   string           `json:"store_id"`
	area      string           `json:"area"`
	StoreName string           `json:"store_name"`
	Data      []UrlProcessInfo `json:"data"`
}

func (h *UtilHandler) GetVisitInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Visit Info")
	area := r.URL.Query().Get("area")
	storeId := r.URL.Query().Get("storeid")
	startDate := r.URL.Query().Get("startdate")
	endDate := r.URL.Query().Get("enddate")

	// startDate, err := time.Parse(stDate, stDate)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// endDate, err := time.Parse(eDate, eDate)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	if len(area) == 0 || len(storeId) == 0 || len(startDate) == 0 || len(endDate) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "area or storeid or startdate or enddate is missing"}`))
		return
	}
	var jobs []models.Job
	result := h.db.Where("store_id = ? AND created_at BETWEEN ? AND ?", storeId, startDate, endDate).Find(&jobs)
	// result := h.db.Where("store_id = ?", storeId).Find(&jobs)

	if result.Error != nil {
		fmt.Println("result.Error", result.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "job_id not found"}`))
		return
	} else {
		var store models.Store

		res := h.db.Where("id = ?", storeId).Find(&store)
		if res.Error != nil {
			fmt.Println("result.Error", result.Error)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "store not found"}`))
			return
		}

		var visitInfo VisitInfo
		visitInfo.StoreName = store.StoreName
		visitInfo.StoreId = storeId
		visitInfo.area = area
		if len(jobs) > 0 {
			for _, job := range jobs {
				var metaDatas []models.Metadata
				result := h.db.Where("job_id = ?", job.ID).Find(&metaDatas)
				if result.Error != nil {
					fmt.Println("result.Error", result.Error)
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{"message": "job_id not found"}`))
					return
				}

				var urlProcessInfo UrlProcessInfo
				urlProcessInfo.Date = job.CreatedAt

				for _, v := range metaDatas {
					urlProcessInfo.Perimeter = 2*v.Width + 2*v.Height
				}
				visitInfo.Data = append(visitInfo.Data, urlProcessInfo)
			}
		} else {
			visitInfo.Data = []UrlProcessInfo{}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(visitInfo)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
