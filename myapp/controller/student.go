package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Helper to convert string to int64
func getUserId(userIdParam string) (int64, error) {
	return strconv.ParseInt(userIdParam, 10, 64)
}

// Create student (POST)
func AddStudent(w http.ResponseWriter, r *http.Request) {
	var stud model.Student

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Debug: Print what we received
	println("Received student - ID:", stud.StdId, "First:", stud.FirstName, "Last:", stud.LastName, "Email:", stud.Email)

	if err := stud.Create(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Return the created student (so frontend can display it)
	httpResp.RespondWithJSON(w, http.StatusCreated, stud)
}

// Get one student (GET)
func GetStud(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["sid"]
	stdId, err := getUserId(sid)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	s := model.Student{StdId: stdId}
	if err := s.Read(); err != nil {
		if err == sql.ErrNoRows {
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, s)
}

// Update student (PUT)
func UpdateStud(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oldSid := vars["sid"]
	oldStdId, err := getUserId(oldSid)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	if err := stud.Update(oldStdId); err != nil {
		if err == sql.ErrNoRows {
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, stud)
}

// Delete student (DELETE)
func DeleteStud(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["sid"]
	stdId, err := getUserId(sid)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	s := model.Student{StdId: stdId}
	if err := s.Delete(); err != nil {
		if err == sql.ErrNoRows {
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// Get all students (GET)
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := model.GetAllStudents()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, students)
}
