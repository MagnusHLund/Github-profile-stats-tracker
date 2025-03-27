package models

type Page struct {
	PageId uint `gorm:"primaryKey"`
	PageUrl string `gorm:"unique; not null"`
}
