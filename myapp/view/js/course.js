// Course Management JavaScript
var selectedRow = null;

// Wait for DOM to be fully loaded before running code
document.addEventListener('DOMContentLoaded', function() {
    console.log("DOM loaded, initializing course page...");
    
    // Set up the add button
    var addBtn = document.getElementById("button-add");
    if (addBtn) {
        addBtn.onclick = function() { addCourse(); };
    }
    
    // Load all courses
    loadAllCourses();
});

function addCourse() {
    console.log("addCourse function called!");
    
    // Get form elements
    var cidInput = document.getElementById("cid");
    var cnameInput = document.getElementById("cname");
    
    // Check if elements exist
    if (!cidInput) {
        console.error("Could not find element with id 'cid'");
        alert("Error: Course ID input not found");
        return;
    }
    if (!cnameInput) {
        console.error("Could not find element with id 'cname'");
        alert("Error: Course Name input not found");
        return;
    }
    
    var courseData = {
        cid: cidInput.value,
        coursename: cnameInput.value
    };
    
    console.log("Course data:", courseData);
    
    if (courseData.cid === "") {
        alert("Please enter Course ID");
        return;
    }
    if (courseData.coursename === "") {
        alert("Please enter Course Name");
        return;
    }
    
    fetch('/course', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(courseData)
    })
    .then(response => {
        console.log("Response status:", response.status);
        if (response.ok) {
            return response.json();
        }
        throw new Error('HTTP error! status: ' + response.status);
    })
    .then(data => {
        console.log("Success:", data);
        alert("Course added successfully!");
        loadAllCourses();  // Refresh the table
        resetForm();
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Failed to add course: " + error.message);
    });
}

function resetForm() {
    var cidInput = document.getElementById("cid");
    var cnameInput = document.getElementById("cname");
    
    if (cidInput) cidInput.value = "";
    if (cnameInput) cnameInput.value = "";
}

function loadAllCourses() {
    console.log("Loading courses...");
    fetch('/courses')
        .then(response => {
            if (!response.ok) {
                throw new Error("HTTP error! status: " + response.status);
            }
            return response.json();
        })
        .then(courses => {
            console.log("Received courses:", courses);
            
            var table = document.getElementById("myTable");
            if (!table) {
                console.error("Table with id 'myTable' not found");
                return;
            }
            
            // Clear rows except header
            while (table.rows.length > 1) {
                table.deleteRow(1);
            }
            
            // Add each course to table
            courses.forEach(course => {
                var row = table.insertRow();
                row.insertCell(0).innerHTML = course.cid;
                row.insertCell(1).innerHTML = course.coursename;
                // Use the same button styling as student page
                row.insertCell(2).innerHTML = '<button id="button-1" onclick="deleteCourse(this)">Delete</button>';
                row.insertCell(3).innerHTML = '<button id="button-2" onclick="updateCourse(this)">Edit</button>';
            });
            
            console.log("Table now has", courses.length, "courses");
        })
        .catch(error => {
            console.error("Error loading courses:", error);
            alert("Error loading courses: " + error.message);
        });
}

function deleteCourse(btn) {
    if (confirm('Delete this course?')) {
        var row = btn.parentElement.parentElement;
        var cid = row.cells[0].innerHTML;
        
        fetch('/course/' + cid, { method: 'DELETE' })
            .then(response => {
                if (response.ok) {
                    row.remove();
                    alert("Course deleted successfully!");
                } else {
                    throw new Error("Delete failed");
                }
            })
            .catch(error => {
                console.error("Error:", error);
                alert("Error deleting course: " + error.message);
            });
    }
}
function updateCourse(btn) {
    selectedRow = btn.parentElement.parentElement;
    var cidInput = document.getElementById("cid");
    var cnameInput = document.getElementById("cname");
    
    if (cidInput) cidInput.value = selectedRow.cells[0].innerHTML;
    if (cnameInput) cnameInput.value = selectedRow.cells[1].innerHTML;
    
    var addBtn = document.getElementById("button-add");
    if (addBtn) {
        addBtn.innerHTML = "Update";
    }
}

function saveUpdate() {
    if (!selectedRow) {
        alert("No course selected for update");
        return;
    }
    
    var oldCid = selectedRow.cells[0].innerHTML;
    var cidInput = document.getElementById("cid");
    var cnameInput = document.getElementById("cname");
    
    var updatedData = {
        cid: cidInput ? cidInput.value : "",
        coursename: cnameInput ? cnameInput.value : ""
    };
    
    fetch('/course/' + oldCid, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(updatedData)
    })
    .then(response => {
        if (response.ok) {
            loadAllCourses();
            resetForm();
            var addBtn = document.getElementById("button-add");
            if (addBtn) {
                addBtn.innerHTML = "Add";
            }
            selectedRow = null;
            alert("Course updated successfully!");
        } else {
            throw new Error("Update failed");
        }
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Error updating course: " + error.message);
    });
}

// Make functions globally available
window.addCourse = addCourse;
window.deleteCourse = deleteCourse;
window.updateCourse = updateCourse;
window.saveUpdate = saveUpdate;