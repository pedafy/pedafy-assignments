package apiv1

import "net/http"

func (a *APIv1) statusGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (a *APIv1) statusGetByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (a *APIv1) statusGetByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
