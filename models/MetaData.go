package models

type Metadata struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	JobId  int    `json:"job_id"`
}
