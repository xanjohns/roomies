package util

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MainController struct {
	Content      *fyne.Container
	Window       fyne.Window
	App          fyne.App
	AppView      RoomiesView
	FeedView     RoomiesView
	PeopleView   RoomiesView
	LoginView    RoomiesView
	RegisterView RoomiesView
	PaymentView  RoomiesView
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
		fmt.Println(roomiesView.GetViewName())
		newContent := this.CreateBorder(roomiesView)
		this.Window.SetContent(newContent)
	}
}

func (this *MainController) CreateBorder(centerView RoomiesView) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() { centerView.RefreshContent() }),
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

	logoutButton := widget.NewButton("Logout", func() { this.Window.SetContent(this.LoginView.GetCanvasObject()) })
	topBar := container.NewBorder(nil, nil, toolbar, logoutButton)

	return container.NewBorder(topBar, centerTabs, nil, nil, centerView.GetCanvasObject())

}
