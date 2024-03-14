package endpoints

import (
	"encoding/json"
	"net/http"
	"roomies/code/dao"
	"roomies/code/shared"
)

type LoginHandler struct {
	responseWriter http.ResponseWriter
	request        *http.Request
	dao            dao.LoginDAO
}

func NewLoginHandler(listDAO dao.GroceryListDAO) *LoginHandler {
	return nil
}

func (l *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	loginBody := shared.GetNewLoginRequest()
	err := json.NewDecoder(r.Body).Decode(loginBody)
	if err != nil {
		return
	}
	if l.dao.ValidateLogin(loginBody.Username, loginBody.Password) {
		w.WriteHeader(403)
	} else {
		w.WriteHeader(200)
	}
}
