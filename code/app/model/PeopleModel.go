package model

import "strconv"

type PeopleModel struct {

}

type Contact struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func NewPeopleModel() *PeopleModel {
	return new(PeopleModel)
}

func (this *PeopleModel) GetItemsHTTP() []Contact {

	// Mock data
	contacts := []Contact {
		Contact{
			Name: "Bob John",
		},
		Contact{
			Name: "Jane Sapphire",
		},
		Contact{
			Name: "Franz Ire",
		},
	}

	for i := 0; i < 30; i++ {
		contacts = append(contacts, Contact{Name: "User " + strconv.Itoa(i)})
	}

	return contacts
}