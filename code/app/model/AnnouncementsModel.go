package model

import "roomies/code/shared"
import "net/http"
import "log"
import "io"
import "encoding/json"


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

func (this *AnnouncementsModel) GetItemsHTTP() []shared.Announcement {

	response, err := http.Get("http://localhost:8081/announcements")
	if err != nil {
		log.Println(err)
		return nil
	}
	data, _ := io.ReadAll(response.Body)

	var items []shared.Announcement
	if err := json.Unmarshal(data, &items); err != nil {
		log.Println(err)
	}
	log.Println(items)

	// Dummy data

	// items := []string {
	// 	"Eggs",
	// 	"Milk",
	// 	"Flour",
	// }

	return items
}