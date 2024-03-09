package main

import (
	"roomies/code/app/model"
	"roomies/code/app/util"
	"roomies/code/app/view"
)

func main() {
	mainController := util.NewMainController()

	mainController.AppView = view.NewAppsView(mainController)
	mainController.FeedView = view.NewFeedView(mainController, model.NewFeedModel())
	mainController.PeopleView = view.NewPeopleView(mainController, model.NewPeopleModel())
	mainController.LoginView = view.NewLoginView(mainController, model.NewLoginModel())
	mainController.RegisterView = view.NewRegisterView(mainController, model.NewRegisterModel())


	mainController.Window.SetContent(mainController.LoginView.GetCanvasObject())
	// mainController.SwitchToView(mainController.AppView)
	mainController.ShowAndRun()

}
