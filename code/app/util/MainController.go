package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MainController struct {
	Content    *fyne.Container
	Window     fyne.Window
	App        fyne.App
	AppView    RoomiesView
	FeedView   RoomiesView
	PeopleView RoomiesView
}

func NewMainController() *MainController {
	mainController := new(MainController)
	mainController.App = app.New()
	mainController.Window = mainController.App.NewWindow("Roomies")
	mainController.Window.Resize(fyne.NewSize(400, 650))

	return mainController
}

func (this *MainController) ShowAndRun() {
	if this.Content != nil {
		this.Window.SetContent(this.Content)
	}
	this.Window.ShowAndRun()
}

func (this *MainController) SwitchToView(roomiesView RoomiesView) {
	if roomiesView != nil {
		roomiesView.RefreshContent()
		newContent := this.CreateBorder(roomiesView.GetCanvasObject())
		this.Window.SetContent(newContent)
	}
}

func (this *MainController) CreateBorder(centerView fyne.CanvasObject) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}),
	)

	tabs := container.NewHBox(
		widget.NewButton("Feed", func() {
			this.SwitchToView(this.FeedView)
		}),
		widget.NewButton("Apps", func() {
			this.SwitchToView(this.AppView)
		}),
		widget.NewButton("People", func() {
			this.SwitchToView(this.PeopleView)
		}),
	)

	centerTabs := container.NewCenter(tabs)

	return container.NewBorder(toolbar, centerTabs, nil, nil, centerView)

}
