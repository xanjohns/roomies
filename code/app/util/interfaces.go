package util

import "fyne.io/fyne/v2"

type RoomiesView interface {
	RefreshContent()
	GetCanvasObject() fyne.CanvasObject
	GetViewName() string
}
