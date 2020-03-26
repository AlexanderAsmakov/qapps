package models

import "github.com/jinzhu/gorm"

type App struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255);not null"`
	Title            string `gorm:"type:varchar(255);not null"`
	BundleIdentifier string `gorm:"type:varchar(255);not null"`
	OS               string `gorm:"type:varchar(32);not null;default:'ios'"`
}
