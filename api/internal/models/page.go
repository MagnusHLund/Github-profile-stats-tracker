package models

type Page struct {
	PageId  uint   `gorm:"primaryKey"`
	PageUrl string `gorm:"type:varchar(255);unique; not null"`
}
