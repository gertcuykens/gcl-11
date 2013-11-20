user = {
    'json':function(){
        $.ajax({
            type: "GET",
            url: "openid",
        })
        .done(function(j) {
            console.log(j)
            if (j) { $("#user").html("Hello "+j.email); $( "#logout" ).show(); }
            else { $( "#login" ).show(); }
        });
    }
}