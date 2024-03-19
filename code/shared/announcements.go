package shared

import "time"


type Announcement struct {
	ItemID string
	AddedByID string
	AnnouncementText string
	Timestamp time.Time
}