package apiv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *APIv1) statusGetAllHandler(w http.ResponseWriter, r *http.Request) {
	status, err := a.dbHandler.GetAllStatus()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if len(status) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"data":""}`))
	} else {
		jsonStatus, err := json.Marshal(status)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonStatus)
		}
	}
}

func (a *APIv1) statusGetByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, _ := strconv.Atoi(idstr)
	status, err := a.dbHandler.GetStatusByID(id)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if status.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonStatus, err := json.Marshal(status)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonStatus)
		}
	}
}

func (a *APIv1) statusGetByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	status, err := a.dbHandler.GetStatusByName(name)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if status.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonStatus, err := json.Marshal(status)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonStatus)
		}
	}
}
