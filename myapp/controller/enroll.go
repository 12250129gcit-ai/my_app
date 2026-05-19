package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Add enrollment (POST)
func AddEnrollment(w http.ResponseWriter, r *http.Request) {
	var enroll model.Enroll

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&enroll); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	// Set current date
	enroll.DateEnrolled = time.Now().Format("2006-01-02 15:04:05")

	// Save to database
	if err := enroll.EnrollStud(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusCreated, enroll)
}

// Get all enrollments (GET)
func GetAllEnrollments(w http.ResponseWriter, r *http.Request) {
	enrollments, err := model.GetAllEnrollments()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusOK, enrollments)
}

// Delete enrollment (DELETE)
func DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sidStr := vars["sid"]
	cid := vars["cid"]

	// Convert sid from string to int64
	stdId, err := strconv.ParseInt(sidStr, 10, 64)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	enroll := model.Enroll{
		StdId:    stdId,
		CourseID: cid,
	}

	if err := enroll.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "Enrollment not found")
		return
	}

	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
