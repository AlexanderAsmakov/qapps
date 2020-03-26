package models

import "github.com/jinzhu/gorm"

type Build struct {
	gorm.Model
	App               App `gorm:"foreignkey:AppID"`
	AppID             uint
	Branch            string `gorm:"type:varchar(255);not null"`
	Comment           string `gorm:"type:text;not null"`
	BundleVersion     string `gorm:"type:varchar(255);not null"`
	BundleIdentifier  string `gorm:"type:varchar(255);not null"`
	BundleName        string `gorm:"type:varchar(255);not null"`
	BundleDisplayName string `gorm:"type:varchar(255);not null"`
}
