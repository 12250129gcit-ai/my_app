package routes

import (
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()

	// ========== STUDENT ROUTES ==========
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")
	router.HandleFunc("/students", controller.GetAllStudents).Methods("GET")

	// ========== COURSE ROUTES ==========
	router.HandleFunc("/course", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/courses", controller.GetAllCourses).Methods("GET")

	// ========== ENROLLMENT ROUTES (ADD THESE) ==========
	router.HandleFunc("/enroll", controller.AddEnrollment).Methods("POST")
	router.HandleFunc("/enrollments", controller.GetAllEnrollments).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}", controller.DeleteEnrollment).Methods("DELETE")

	// ========== ADMIN ROUTES ==========
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")

	// ========== STATIC FILES ==========
	fs := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fs)

	println("✅ Server starting on http://localhost:8080")
	println("📋 Registered routes:")
	println("   POST   /student")
	println("   GET    /student/{sid}")
	println("   PUT    /student/{sid}")
	println("   DELETE /student/{sid}")
	println("   GET    /students")
	println("   POST   /course")
	println("   GET    /course/{cid}")
	println("   PUT    /course/{cid}")
	println("   DELETE /course/{cid}")
	println("   GET    /courses")
	println("   POST   /enroll")
	println("   GET    /enrollments")
	println("   DELETE /enroll/{sid}/{cid}")

	http.ListenAndServe(":8080", router)
}
