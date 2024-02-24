package endpoints

import "net/http"

type Endpoint interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
