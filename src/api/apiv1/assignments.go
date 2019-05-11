package apiv1

import "net/http"

func (a *APIv1) assignmentsGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"valid":true}`))
}

func (a *APIv1) assignmentsGetOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"valid":true}`))
}

func (a *APIv1) assignmentsGetFilterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"valid":true}`))
}

func (a *APIv1) assignmentsGetByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"valid":true}`))
}

func (a *APIv1) assignmentsNewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"valid":true}`))
}

func (a *APIv1) assignmentsModifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"valid":true}`))
}

func (a *APIv1) assignmentsArchiveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"valid":true}`))
}
