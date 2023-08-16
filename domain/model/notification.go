package model

import "time"

type NotificationType string

type Notification struct {
	ID        string
	Type      NotificationType
	Timestamp time.Time
	Recipient string
}
