package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

// Create course (POST)
func AddCourse(w http.ResponseWriter, r *http.Request) {
	var course model.Course

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := course.Create(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusCreated, course)
}

// Get one course (GET)
func GetCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["cid"]

	course := model.Course{Cid: cid}
	if err := course.Read(); err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, course)
}

// Update course (PUT)
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oldCid := vars["cid"]

	var course model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	if err := course.Update(oldCid); err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, course)
}

// Delete course (DELETE)
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["cid"]

	course := model.Course{Cid: cid}
	if err := course.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// Get all courses (GET)
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := model.GetAllCourses()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}
