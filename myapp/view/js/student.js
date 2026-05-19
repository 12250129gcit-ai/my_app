// Global variable for editing
var selectedRow = null;

// Function to add a new student
function addStudent() {
    // Get form data
    var data = {
        stdid: parseInt(document.getElementById("sid").value),
        fname: document.getElementById("fname").value,
        lname: document.getElementById("lname").value,
        email: document.getElementById("email").value
    };
    
    var sid = data.stdid;
    
    // Validation
    if (isNaN(sid) || sid === "") {
        alert("Please enter a valid Student ID");
        return;
    }
    if (data.fname === "") {
        alert("First name cannot be empty");
        return;
    }
    if (data.email === "") {
        alert("Email cannot be empty");
        return;
    }
    
    console.log("Sending data:", data);
    
    // Send POST request to add student
    fetch('/student', {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        console.log("Response status:", response.status);
        if (response.ok) {
            return response.json();
        } else {
            throw new Error("Server returned " + response.status);
        }
    })
    .then(studentData => {
        console.log("Student created:", studentData);
        // Reload all students to refresh the table
        loadAllStudents();
        // Reset the form
        resetForm();
        // Change button back to "Add" if it was in update mode
        var btn = document.getElementById("button-add");
        if (btn && btn.innerHTML === "Update") {
            btn.innerHTML = "Add";
            btn.setAttribute("onclick", "addStudent()");
            selectedRow = null;
        }
        alert("Student added successfully!");
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Error: " + error.message);
    });
}

// Reset form fields
function resetForm() {
    document.getElementById("sid").value = "";
    document.getElementById("fname").value = "";
    document.getElementById("lname").value = "";
    document.getElementById("email").value = "";
}

// Function to add a row to the table
function addRowToTable(student) {
    var table = document.getElementById("myTable");
    var row = table.insertRow(table.rows.length);
    
    var cell1 = row.insertCell(0);
    var cell2 = row.insertCell(1);
    var cell3 = row.insertCell(2);
    var cell4 = row.insertCell(3);
    var cell5 = row.insertCell(4);
    var cell6 = row.insertCell(5);
    
    cell1.innerHTML = student.stdid;
    cell2.innerHTML = student.fname || student.firstname;
    cell3.innerHTML = student.lname || student.lastname;
    cell4.innerHTML = student.email;
    cell5.innerHTML = '<button class="delete-btn" onclick="deleteStudent(this)">Delete</button>';
    cell6.innerHTML = '<button class="edit-btn" onclick="updateStudent(this)">Edit</button>';
}

// Load all students from API and display in table
function loadAllStudents() {
    console.log("Loading all students...");
    fetch('/students')
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to fetch students");
            }
            return response.json();
        })
        .then(students => {
            console.log("Received students:", students);
            
            // Clear existing table rows (keep the header row)
            var table = document.getElementById("myTable");
            while (table.rows.length > 1) {
                table.deleteRow(1);
            }
            
            // Add each student to table
            students.forEach(student => {
                addRowToTable(student);
            });
            
            console.log("Table now has", table.rows.length - 1, "students");
        })
        .catch(error => {
            console.error("Error loading students:", error);
            alert("Error loading students: " + error.message);
        });
}

// Delete student
function deleteStudent(btn) {
    if (confirm('Are you sure you want to DELETE this student?')) {
        var row = btn.parentElement.parentElement;
        var sid = row.cells[0].innerHTML;
        
        fetch('/student/' + sid, {
            method: "DELETE"
        })
        .then(response => {
            if (response.ok) {
                // Remove the row from the table
                var rowIndex = row.rowIndex;
                document.getElementById("myTable").deleteRow(rowIndex);
                alert("Student deleted successfully");
            } else {
                throw new Error("Delete failed");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("Error deleting student: " + error.message);
        });
    }
}

// Update/Edit student - populate form with selected student data
function updateStudent(btn) {
    selectedRow = btn.parentElement.parentElement;
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML;
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML;
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML;
    document.getElementById("email").value = selectedRow.cells[3].innerHTML;
    
    var btn = document.getElementById("button-add");
    if (btn) {
        btn.innerHTML = "Update";
        btn.onclick = function() { updateStudentData(); };
    }
}

// Perform the actual update
function updateStudentData() {
    var sid = document.getElementById("sid").value;
    var newData = {
        stdid: parseInt(sid),
        fname: document.getElementById("fname").value,
        lname: document.getElementById("lname").value,
        email: document.getElementById("email").value
    };
    
    fetch('/student/' + sid, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            // Reload all students to refresh the table
            loadAllStudents();
            resetForm();
            
            var btn = document.getElementById("button-add");
            if (btn) {
                btn.innerHTML = "Add";
                btn.onclick = function() { addStudent(); };
            }
            selectedRow = null;
            alert("Student updated successfully");
        } else {
            throw new Error("Update failed");
        }
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Error updating student: " + error.message);
    });
}

// Load students when page loads
window.onload = function() {
    console.log("Page loaded, initializing...");
    loadAllStudents();
};