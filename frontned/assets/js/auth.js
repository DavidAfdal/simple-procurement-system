

$(document).ready(function() {
    $('#loginForm').on('submit', function(e) {
        e.preventDefault();

        const data = {
            username: $('#loginUsername').val(),
            password: $('#loginPassword').val()
        };

        apiRequest('/users/login', 'POST', data, function(res) {
            localStorage.setItem('jwtToken', res.data.access_token);
            Swal.fire('Success', res.message, 'success').then(() => {
                window.location.href = 'dashboard.html';
            });
        });
    });
    
    $('#registerForm').on('submit', function(e) {
        e.preventDefault();

        const data = {
            username: $('#registerUsername').val(),
            password: $('#registerPassword').val()
        };

        apiRequest('/users/register', 'POST', data, function(res) {
            Swal.fire('Success', res.message, 'success').then(() => {
                window.location.href = 'login.html';
            });
        });
    });

    $('#logoutButton').on('click', function() {
        const token = localStorage.getItem('jwtToken');

        if (token) {
            $.ajax({
                url: '/users/logout',
                type: 'POST',
                headers: { 'Authorization': 'Bearer ' + token },
                success: function() {
                    console.log('Logout success');
                },
                error: function(err) {
                    console.error('Logout failed', err);
                }
            });
        }

        localStorage.removeItem('jwtToken');
        window.location.href = 'login.html';
    });
});


