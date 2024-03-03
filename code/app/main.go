package main

import (
	"roomies/code/app/model"
	"roomies/code/app/util"
	"roomies/code/app/view"
)

func main() {
	mainController := util.NewMainController()

	mainController.AppView = view.NewAppsView(mainController)
	mainController.FeedView = view.NewFeedView()
	mainController.PeopleView = view.NewPeopleView(mainController, model.NewPeopleModel())


	mainController.SwitchToView(mainController.AppView)
	mainController.ShowAndRun()

}
