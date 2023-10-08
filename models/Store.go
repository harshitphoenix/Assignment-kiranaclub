package models

type Store struct {
	// gorm.Model
	ID        string `gorm:"uniqueIndex:compositeindex;type:text;not null"`
	StoreName string
	AreaCode  int
}
