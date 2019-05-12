package apiv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-assignments/src/database"
)

func (a *APIv1) assignmentsGetAllHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var assignments []database.Assignments
	order, ok := r.URL.Query()["sort"]

	if ok && len(order[0]) > 0 {
		assignments, err = a.dbHandler.GetAllOrderAssignments(order[0])
	} else {
		assignments, err = a.dbHandler.GetAllAssignments()
	}
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if len(assignments) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonAssignments, err := json.Marshal(assignments)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonAssignments)
		}
	}
}

func (a *APIv1) assignmentsGetFilterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assignments, err := a.dbHandler.GetAllByFilterAssignments(vars["id_type"], vars["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if len(assignments) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonAssignments, err := json.Marshal(assignments)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonAssignments)
		}
	}
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
