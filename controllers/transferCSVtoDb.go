package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"server/models"
	"strconv"
)

func (h *UtilHandler) TransferCSVtoDb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TransferCSVtoDb")
	// filePath, err := filepath.Abs("StoreMasterAssignment.csv")
	path, err := os.Getwd()
	println("filePath", path)
	csvFile, err := os.Open(path + "/controllers/StoreMasterAssignment.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	// reader.Comma = ';'
	// reader.FieldsPerRecord = -1
	// stm, err := h.db.Createdb.Prepare("INSERT INTO stores(store_id,store_name, area_code) VALUES($1,$2,$3)")
	for _, record := range records {
		areaCode, _ := strconv.Atoi(record[0])
		store := models.Store{
			ID:        record[2],
			StoreName: record[1],
			AreaCode:  areaCode,
		}
		result := h.db.Create(&store)
		if result.Error != nil {
			fmt.Println("result.Error", result.Error)
		}
	}

}
