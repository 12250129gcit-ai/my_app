function signup() {
    var firstname = document.getElementById("firstname").value;
    var lastname = document.getElementById("lastname").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
    var confirmPassword = document.getElementById("confirm_password").value;
    
    if (firstname === "") {
        alert("Please enter first name");
        return;
    }
    if (email === "") {
        alert("Please enter email");
        return;
    }
    if (password === "") {
        alert("Please enter password");
        return;
    }
    if (password !== confirmPassword) {
        alert("Passwords do not match!");
        return;
    }
    
    var data = {
        firstname: firstname,
        lastname: lastname,
        email: email,
        password: password
    };
    
    fetch('/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (response.status === 201) {
            return response.json();
        } else {
            return response.json().then(err => {
                throw new Error(err.error || "Signup failed");
            });
        }
    })
    .then(data => {
        console.log("Signup success:", data);
        alert("Account created successfully! Please login.");
        window.location.href = "login.html";
    })
    .catch(error => {
        console.error("Error:", error);
        alert("Signup failed: " + error.message);
    });
}