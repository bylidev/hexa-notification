package domain

import "time"

type Notification struct {
	SendedAt time.Time
	Details  []byte
}
