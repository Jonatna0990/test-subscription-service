package entity

import "time"

type Filter struct {
	UserID      string
	ServiceName string
	From        time.Time
	To          time.Time
}
