package controllers

import (
	"fmt"
	"image"
	"os"
	"server/models"
	"server/utils"
	"time"

	"gorm.io/gorm"
)

func ProcessJob(db *gorm.DB, jobId int, visit Visit) {
	fmt.Println("ProcessJob")
	var job models.Job
	db.Where("id = ?", jobId).First(&job)
	fmt.Println(job)
	if job.ID != 0 {
		// var store models.Store
		store := db.Where("id = ?", job.StoreId).First(&models.Store{})
		fmt.Println(store)
		if store == nil {
			job.Status = "failed"
			db.Save(&job)
		} else {

			for _, url := range visit.ImageURLs {
				fmt.Println("url", url)
				var metaData models.Metadata
				metaData.URL = url
				metaData.JobId = job.ID

				file, err := utils.GetImage(url)

				if err != nil {
					fmt.Println("error", err)
					panic(err)
					job.Status = "failed"
				}
				defer file.Close()

				reader, err := os.Open(file.Name())
				if err != nil {
					fmt.Println("error", err)
					panic(err)
					job.Status = "failed"
				}
				img, _, err := image.DecodeConfig(reader)
				if err != nil {
					fmt.Println("error", err)
					panic(err)
					job.Status = "failed"
				}
				metaData.Width = img.Width
				metaData.Height = img.Height
				db.Create(&metaData) //create a new record in the database for a processed image
				os.Remove(file.Name())
				utils.Sleeper()
			}

			if job.Status != "failed" {
				job.Status = "completed"

			}
			job.CompletedAt = time.Now()
			db.Save(&job)
		}
	}
}
