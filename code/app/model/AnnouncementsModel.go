package model


type AnnouncementsModel struct {

}

type Announcement struct {
	Date string `json:"date"`
	Message string `json:"message"`
	From string `json:"from"`
}

func NewAnnouncementsModel() *AnnouncementsModel {
	return new(AnnouncementsModel)
}

func (this *AnnouncementsModel) GetItemsHTTP() []Announcement {

	// Mock data
	announcements := []Announcement {
		Announcement{
			Date: "Jan 21",
			Message: "Reminder about the free concert tonight!",
			From: "Admin",
		},
		Announcement{
			Date: "Feb 01",
			Message: "Enter our giveaway to win prizes!",
			From: "Admin",
		},
		Announcement{
			Date: "Mar 29",
			Message: "Rent is due soon.",
			From: "Admin",
		},
	}

	return announcements
}