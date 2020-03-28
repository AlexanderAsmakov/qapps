package models

import "time"

type AppMetadata struct {
	ID           int64     `json:"id" db:"id"`
	AppID        int64     `json:"appId" db:"app_id"`
	Name         string    `json:"name" db:"name"`
	Keywords     string    `json:"keywords" db:"keywords"`
	Description  string    `json:"description" db:"description"`
	ReleaseNotes string    `json:"releaseNotes" db:"release_notes"`
	MarketingUrl string    `json:"marketingUrl" db:"marketing_url"`
	PrivacyUrl   string    `json:"privacyUrl" db:"privacy_url"`
	SupportUrl   string    `json:"supportUrl" db:"support_url"`
	AuthorID     int64     `json:"authorId" db:"author_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}
