

const baseUrl = 'http://localhost:8080/api';
function getToken() {
    return localStorage.getItem('jwtToken');
}

function apiRequest(endpoint, method = 'GET', data = {}, successCallback, errorCallback) {
    $.ajax({
        url: baseUrl + endpoint,
        method: method,
        data: method === 'GET' ? data : JSON.stringify(data),
        contentType: 'application/json',
        headers: {
            Authorization: getToken() ? 'Bearer ' + getToken() : ''
        },
        success: successCallback,
        error: function(xhr) {
            console.log(xhr)
            const msg = xhr.responseJSON?.message || 'Something went wrong';
            Swal.fire('Error', msg, 'error');
            if (errorCallback) errorCallback(xhr);
        }
    });
}