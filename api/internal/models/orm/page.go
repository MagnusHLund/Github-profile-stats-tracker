package orm

type Page struct {
	PageId               uint   `gorm:"primaryKey"`
	PageOwnerGitUsername string `gorm:"type:varchar(39);unique; not null"`
}
