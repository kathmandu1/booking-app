package models

import (
	"gorm.io/gorm"
)

type Facility struct {
	ID               int    `gorm:"AUTO_INCREMENT"`
	FacilityName     string `gorm:"not null"`
	Status           bool   `gorm:"default:false"`
	FacilityCategory int
	gorm.Model
}
