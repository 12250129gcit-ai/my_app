// Check if user is logged in
function checkAuth() {
    // Send a request to check if cookie exists
    fetch('/students', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
    })
    .then(response => {
        if (response.status === 401) {
            // Not authenticated
            alert("Please login first!");
            window.location.href = "login.html";
        }
    })
    .catch(error => {
        console.error("Auth check error:", error);
    });
}

// Call checkAuth when page loads
checkAuth();