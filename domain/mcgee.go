package domain

type McGee struct {
	Title     string `json:"title" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time"  binding:"required"`
}
