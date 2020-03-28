package models

type App struct {
	ID               int64           `json:"id" db:"id"`
	Name             string          `json:"name" db:"name"`
	Title            string          `json:"title" db:"title"`
	BundleIdentifier string          `json:"bundleIdentifier" db:"bundle_identifier"`
	OS               operationSystem `json:"os" db:"os"`
}

type operationSystem string

func (os operationSystem) String() string {
	return string(os)
}

const (
	OperationSystemIos     operationSystem = "ios"
	OperationSystemAndroid operationSystem = "android"
)
