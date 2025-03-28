package models

type Visitor struct {
	VisitorId       uint   `gorm:"primaryKey"`
	PageId          uint   `gorm:"not null;uniqueIndex:idx_hashedip_pageid"`
	TotalVisits     uint32 `gorm:"not null"`
	FirstVisit      uint   `gorm:"not null"`
	LatestVisit     uint   `gorm:"not null"`
	HashedIpAddress string `gorm:"not null;uniqueIndex:idx_hashedip_pageid"` // TODO: Modify max length
}
