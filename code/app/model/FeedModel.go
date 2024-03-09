package model

import "time"

type FeedModel struct {

}

type FeedItem struct {
	ItemID string
	ItemText string
	AddedByID string
	Timestamp time.Time
	OriginViewName string
}

func NewFeedModel() *FeedModel {
	return new(FeedModel)
}

func (this *FeedModel) GetItemsHTTP() []FeedItem {
	return nil
}