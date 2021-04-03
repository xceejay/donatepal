$(document).ready(function () {



    $('input[type=radio][name=donation_type]').change(function () {
        $(".unanymous").toggle()


    })


    $('.registerform').validate({
        rules: {
            password: {
                minlength: 5
            },
            confirm_password: {
                minlength: 5,
                equalTo: "#password"
            }
        }
    });

    $('#registerbtn').click(function (event) {
        // alert("no")
        event.preventDefault()
        $("#passwordnotif").show().delay(2000).fadeOut();
    });







})