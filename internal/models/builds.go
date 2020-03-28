package models

import "time"

type Build struct {
	ID                int64     `json:"id" db:"id"`
	AppID             int64     `json:"appId" db:"app_id"`
	Branch            string    `json:"branch" db:"branch"`
	Comment           string    `json:"comment" db:"comment"`
	BundleVersion     string    `json:"bundleVersion" db:"bundle_version"`
	BundleIdentifier  string    `json:"bundleIdentifier" db:"bundle_identifier"`
	BundleName        string    `json:"bundleName" db:"bundle_name"`
	BundleDisplayName string    `json:"bundleDisplayName" db:"bundle_display_name"`
	FileName          string    `json:"fileName" db:"file_name"`
	BuildKey          string    `json:"buildKey" db:"build_key"`
	IsRelease         bool      `json:"isRelease" db:"is_release"`
	CreatedAt         time.Time `json:"createdAt" db:"created_at"`
}
