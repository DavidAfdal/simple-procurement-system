

$(document).ready(function() {
    $('#loginForm').on('submit', function(e) {
        e.preventDefault();

        const username = $('#loginUsername').val();

        if (!username) {
            Swal.fire({
                icon: 'error',
                title: 'Username required!',
                text: 'Please input username  first',
                confirmButtonColor: '#3b82f6'
            });
            return;
        }

        const password = $('#loginPassword').val();

        if (!password) {
            Swal.fire({
                icon: 'error',
                title: 'Password required!',
                text: 'Please input password first',
                confirmButtonColor: '#3b82f6'
            });
            return;
        }

        const data = {
            username: username,
            password: password
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


        const username = $('#registerUsername').val();

        if (!username) {
            Swal.fire({
                icon: 'error',
                title: 'Username required!',
                text: 'Please input username  first',
                confirmButtonColor: '#3b82f6'
            });
            return;
        }

        const password = $('#registerPassword').val();

        if (!password) {
            Swal.fire({
                icon: 'error',
                title: 'Password required!',
                text: 'Please input password first',
                confirmButtonColor: '#3b82f6'
            });
            return;
        }


        const data = {
            username: username,
            password: password
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


