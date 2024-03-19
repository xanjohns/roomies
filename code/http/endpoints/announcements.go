package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"roomies/code/dao"
	"roomies/code/shared"
)

type AnnouncementsHandler struct {
	responseWriter http.ResponseWriter
	request *http.Request
	dao dao.AnnouncementsDAO
}

func NewAnnouncementsHandler(announcementsDAO dao.AnnouncementsDAO) *AnnouncementsHandler {
	return &AnnouncementsHandler{
		dao: announcementsDAO,
	}
}

func (h *AnnouncementsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	h.reset()
	h.setup(w,r)

	var res []byte
	var err error

	switch h.request.Method {
	case "GET":
		res, err = json.Marshal(h.getList())
	default:
		fmt.Fprintf(w, "Method not supported\n!")
	}
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	w.Write(res)
}

func (h *AnnouncementsHandler) reset() {
	h.request = nil
	h.responseWriter = nil
}

func (h *AnnouncementsHandler) getList() []*shared.Announcement {
	return h.dao.GetAnnouncements()
}

func (h *AnnouncementsHandler) setup (w http.ResponseWriter, r *http.Request) {
	h.responseWriter = w
	h.request = r
}