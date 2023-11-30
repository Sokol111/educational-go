package model

import "time"

type User struct {
	ID               int64
	Version          int32
	Name             string
	Enabled          bool
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
