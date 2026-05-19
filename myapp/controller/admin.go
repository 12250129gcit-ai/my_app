package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"time"
)

// Signup handler
func Signup(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	if err := admin.Create(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "admin added successfully"})
}

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	defer r.Body.Close()

	if err := admin.Get(); err != nil {
		httpResp.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Set session cookie
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    admin.Email, // Use email as session identifier
		HttpOnly: true,
		Expires:  time.Now().Add(30 * time.Minute),
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Login successful"})
}

// Logout handler
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Logged out successfully"})
}

// VerifyCookie checks if user is authenticated
func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			httpResp.RespondWithError(w, http.StatusUnauthorized, "Please login first")
			return false
		}
		httpResp.RespondWithError(w, http.StatusInternalServerError, "Server error")
		return false
	}

	if cookie.Value == "" {
		httpResp.RespondWithError(w, http.StatusUnauthorized, "Invalid session")
		return false
	}

	return true
}
