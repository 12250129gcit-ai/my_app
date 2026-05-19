// Enrollment Management JavaScript

// Load students and courses when page loads
document.addEventListener('DOMContentLoaded', function() {
    console.log("DOM loaded, initializing enrollment page...");
    loadStudents();
    loadCourses();
    loadEnrollments();
});

// Load students into dropdown
function loadStudents() {
    console.log("Loading students...");
    fetch('/students')
        .then(response => response.json())
        .then(students => {
            console.log("Students loaded:", students.length);
            var select = document.getElementById("sid");
            if (!select) return;
            
            select.innerHTML = '<option value="">Select Student</option>';
            students.forEach(student => {
                var option = document.createElement("option");
                option.value = student.stdid;
                option.text = student.stdid + " - " + (student.fname || student.firstname) + " " + (student.lname || student.lastname);
                select.appendChild(option);
            });
        })
        .catch(error => {
            console.error("Error loading students:", error);
        });
}

// Load courses into dropdown
function loadCourses() {
    console.log("Loading courses...");
    fetch('/courses')
        .then(response => response.json())
        .then(courses => {
            console.log("Courses loaded:", courses.length);
            var select = document.getElementById("cid");
            if (!select) return;
            
            select.innerHTML = '<option value="">Select Course</option>';
            courses.forEach(course => {
                var option = document.createElement("option");
                option.value = course.cid;
                option.text = course.cid + " - " + course.coursename;
                select.appendChild(option);
            });
        })
        .catch(error => {
            console.error("Error loading courses:", error);
        });
}

// Add enrollment
function addEnrollment() {
    console.log("addEnrollment called");
    
    var sidSelect = document.getElementById("sid");
    var cidSelect = document.getElementById("cid");
    
    if (!sidSelect || !cidSelect) {
        alert("Form elements not found");
        return;
    }
    
    var stdid = sidSelect.value;
    var cid = cidSelect.value;
    
    if (stdid === "") {
        alert("Please select a student");
        return;
    }
    if (cid === "") {
        alert("Please select a course");
        return;
    }
    
    var enrollmentData = {
        stdid: parseInt(stdid),
        cid: cid
    };
    
    console.log("Sending enrollment:", enrollmentData);
    
    fetch('/enroll', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(enrollmentData)
    })
    .then(response => {
        console.log("Response status:", response.status);
        if (response.ok) {
            return response.json();
        }
        return response.json().then(err => {
            throw new Error(err.error || "Enrollment failed");
        });
    })
    .then(data => {
        console.log("Enrollment success:", data);
        alert("Student enrolled successfully!");
        loadEnrollments();  // Refresh the table
        resetForm();
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Failed to enroll: " + error.message);
    });
}

// Load all enrollments
function loadEnrollments() {
    console.log("Loading enrollments...");
    fetch('/enrollments')
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch enrollments");
            }
            return response.json();
        })
        .then(enrollments => {
            console.log("Enrollments received:", enrollments);
            
            var table = document.getElementById("myTable");
            if (!table) {
                console.error("Table not found");
                return;
            }
            
            // Clear rows except header
            while (table.rows.length > 1) {
                table.deleteRow(1);
            }
            
            if (enrollments.length === 0) {
                var row = table.insertRow();
                row.insertCell(0).innerHTML = "No enrollments yet";
                row.insertCell(1).innerHTML = "";
                row.insertCell(2).innerHTML = "";
                row.insertCell(3).innerHTML = "";
                row.insertCell(4).innerHTML = "";
                row.insertCell(5).innerHTML = "";
                return;
            }
            
            // Add each enrollment to table
            enrollments.forEach(enrollment => {
                var row = table.insertRow();
                row.insertCell(0).innerHTML = enrollment.stdid;
                row.insertCell(1).innerHTML = enrollment.student_name || "Unknown";
                row.insertCell(2).innerHTML = enrollment.cid;
                row.insertCell(3).innerHTML = enrollment.course_name || "Unknown";
                row.insertCell(4).innerHTML = enrollment.date_enrolled || "";
                // Use same button styling as student page
                row.insertCell(5).innerHTML = '<button id="button-1" onclick="deleteEnrollment(' + enrollment.stdid + ', \'' + enrollment.cid + '\')">Delete</button>';
            });
            
            console.log("Table now has", enrollments.length, "enrollments");
        })
        .catch(error => {
            console.error("Error loading enrollments:", error);
        });
}

// Delete enrollment
function deleteEnrollment(stdid, cid) {
    if (confirm('Remove this student from the course?')) {
        fetch('/enroll/' + stdid + '/' + cid, { 
            method: 'DELETE' 
        })
        .then(response => {
            if (response.ok) {
                alert("Enrollment removed successfully!");
                loadEnrollments();  // Refresh the table
            } else {
                throw new Error("Delete failed");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("Error removing enrollment: " + error.message);
        });
    }
}

// Reset form
function resetForm() {
    var sidSelect = document.getElementById("sid");
    var cidSelect = document.getElementById("cid");
    
    if (sidSelect) sidSelect.value = "";
    if (cidSelect) cidSelect.value = "";
}

// Make functions globally available
window.addEnrollment = addEnrollment;
window.deleteEnrollment = deleteEnrollment;