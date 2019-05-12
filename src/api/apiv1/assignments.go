package apiv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (a *APIv1) assignmentsNewHandler(w http.ResponseWriter, r *http.Request) {
	var assignment database.Assignments
	err := json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"error":"malormated data"}`))
		return
	}

	newAssignment, err := a.dbHandler.NewAssignment(assignment)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else {
		jsonAssignments, err := json.Marshal(newAssignment)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, `{"data":%s}`, jsonAssignments)
		}
	}
}

func (a *APIv1) assignmentsModifyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IDStr := vars["id"]
	ID, _ := strconv.Atoi(IDStr)
	var assignment database.Assignments
	err := json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"error":"malormated data"}`))
		return
	}

	modifiedAssignment, err := a.dbHandler.ModifyAssignment(assignment, ID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if modifiedAssignment.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"error":"unknown data"}`))
	} else {
		jsonAssignments, err := json.Marshal(modifiedAssignment)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"data":%s}`, jsonAssignments)
		}
	}
}

func (a *APIv1) assignmentsArchiveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IDStr := vars["id"]
	ID, _ := strconv.Atoi(IDStr)

	assignment, err := a.dbHandler.ArchiveAssignment(ID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if assignment.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"error":"unknown data"}`))
	} else {
		jsonAssignments, err := json.Marshal(assignment)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"data":%s}`, jsonAssignments)
		}
	}
}
