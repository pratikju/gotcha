$(document).ready(function(){
    $('#login_button').on('click', function(){
        window.location.href = "/home?name=" + $('#username').val();
    });
});
