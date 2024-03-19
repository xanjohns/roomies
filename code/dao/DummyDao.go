package dao

import (
	"fmt"
	"roomies/code/shared"
	"time"
)

type DummyAnnouncementsDAO struct {

} 

func (d *DummyAnnouncementsDAO) GetAnnouncements() []*shared.Announcement {
	return []*shared.Announcement{
		&shared.Announcement{
			ItemID: "1",
			AddedByID: "admin",
			AnnouncementText: "Come to our Christmas Party!",
			Timestamp: time.Now(),
		},
		&shared.Announcement{
			ItemID: "2",
			AddedByID: "admin",
			AnnouncementText: "Cleaning checks are next week.",
			Timestamp: time.Now(),
		},
		&shared.Announcement{
			ItemID: "3",
			AddedByID: "admin",
			AnnouncementText: "If you are moving out. . .",
			Timestamp: time.Now(),
		},
		&shared.Announcement{
			ItemID: "4",
			AddedByID: "admin",
			AnnouncementText: "Please look out for . . .",
			Timestamp: time.Now(),
		},
		&shared.Announcement{
			ItemID: "5",
			AddedByID: "admin",
			AnnouncementText: "Don't forget to. . .",
			Timestamp: time.Now(),
		},
		&shared.Announcement{
			ItemID: "6",
			AddedByID: "admin",
			AnnouncementText: "Temporary internet outage will occure tomorrow.",
			Timestamp: time.Now(),
		},
	}
}

type DummyGroceryListDAO struct {
	GroceryListDAO
	dummyList *shared.GroceryList
}

func GetNewDummyGroceryListDAO() *DummyGroceryListDAO {
	l := shared.GetNewGroceryList([]*shared.GroceryListItem{
		{
			ItemID:        "1111",
			GroupID:       "group1",
			ListItem:      "eggs",
			AddedByID:     "user",
			Recurring:     false,
			RecurringDate: time.Time{},
			Timestamp:     time.Time{},
		}, {
			ItemID:        "2222",
			GroupID:       "group1",
			ListItem:      "milk",
			AddedByID:     "user",
			Recurring:     false,
			RecurringDate: time.Time{},
			Timestamp:     time.Time{},
		}, {
			ItemID:        "3333",
			GroupID:       "group1",
			ListItem:      "flour",
			AddedByID:     "user",
			Recurring:     false,
			RecurringDate: time.Time{},
			Timestamp:     time.Time{},
		}})
	return &DummyGroceryListDAO{
		dummyList: l,
	}
}

func (d *DummyGroceryListDAO) GetList(listID string) *shared.GroceryList {
	return d.dummyList
}

func (d *DummyGroceryListDAO) GetListItem(listID, itemID string) *shared.GroceryListItem {
	return d.dummyList.GetItem(itemID)
}

func (d *DummyGroceryListDAO) CreateList() {
	fmt.Println("CreateList not implemented for dummy")
}
func (d *DummyGroceryListDAO) AppendListItem(listID string, item *shared.GroceryListItem) {
	fmt.Println("AppendListItem not implemented for dummy")
}
func (d *DummyGroceryListDAO) DeleteListItem(listID, itemID string) {
	fmt.Println("Don't delete the dummy list")
}
