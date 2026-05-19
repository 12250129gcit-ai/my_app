function login() {
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
    
    if (email === "") {
        alert("Please enter email");
        return;
    }
    if (password === "") {
        alert("Please enter password");
        return;
    }
    
    var data = {
        email: email,
        password: password
    };
    
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error("Invalid email or password");
        }
    })
    .then(data => {
        console.log("Login success:", data);
        alert("Login successful!");
        window.location.href = "student.html";
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Login failed: " + error.message);
    });
}

function logout() {
    console.log("Logout button clicked");
    if (confirm("Are you sure you want to logout?")) {
        fetch('/logout', { 
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (response.ok) {
                window.location.href = "login.html";
            } else {
                alert("Logout failed");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("Network error during logout");
        });
    }
}