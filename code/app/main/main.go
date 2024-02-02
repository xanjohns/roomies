package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const appsButtonSizeHeight = 100
const appsButtonSizeWidth = 100
	

func main() {
	a := app.New()
	w := a.NewWindow("Roomies")
	w.Resize(fyne.NewSize(600, 600))


	groceryButton := widget.NewButton("Grocery List", nil)
	groceryButton.Resize(fyne.NewSize(appsButtonSizeWidth, appsButtonSizeHeight))

	calendarButton := widget.NewButton("Calendar", nil)
	calendarButton.Resize(fyne.NewSize(appsButtonSizeWidth, appsButtonSizeHeight))

	announcementsButton := widget.NewButton("Announcements", nil)
	announcementsButton.Resize(fyne.NewSize(appsButtonSizeWidth, appsButtonSizeHeight))

	rentButton := widget.NewButton("Pay Rent", nil)
	rentButton.Resize(fyne.NewSize(appsButtonSizeWidth, appsButtonSizeHeight))

	serviceRequestButton := widget.NewButton("Service Request", nil)
	serviceRequestButton.Resize(fyne.NewSize(appsButtonSizeWidth, appsButtonSizeHeight))

	propertyInfoButton := widget.NewButton("Property Info", nil)
	propertyInfoButton.Resize(fyne.NewSize(appsButtonSizeWidth, appsButtonSizeHeight))

	todosButton := widget.NewButton("Todo's", nil)
	todosButton.Resize(fyne.NewSize(appsButtonSizeWidth, appsButtonSizeHeight))

	apps := container.NewVBox(
		groceryButton,
		layout.NewSpacer(),
		calendarButton,
		layout.NewSpacer(),
		announcementsButton,
		layout.NewSpacer(),
		rentButton,
		layout.NewSpacer(),
		serviceRequestButton,
		layout.NewSpacer(),
		propertyInfoButton,
		layout.NewSpacer(),
		todosButton,
	)

	scroll := container.NewVScroll(apps)
	scrollPadded := container.NewPadded(scroll)


	tabs := container.NewAppTabs(
		container.NewTabItem("Feed", widget.NewLabel("Hellow")),
		container.NewTabItem("Apps", scrollPadded), 
		container.NewTabItem("People", widget.NewLabel("Peoples")),
	)


	tabs.SetTabLocation(container.TabLocationBottom)
	tabs.Resize(fyne.NewSize(60, 300))

	w.SetContent(tabs)
	w.ShowAndRun()

}
