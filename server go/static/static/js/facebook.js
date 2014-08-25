FACEBOOK_CLIENT_ID="249348058430770"

$(document).ready(function() {
  $.ajaxSetup({ cache: true });
  $.getScript('//connect.facebook.net/en_UK/all.js', function(){

    FB.init({
        appId: FACEBOOK_CLIENT_ID,
        version: 'v2.0'
    });

    FB.Event.subscribe('auth.authResponseChange', function(response) {
        console.log(response)
        if (response.status=='connected') {$('#loginButton').hide()}
    });

    FB.getLoginStatus(function(response){

        $.getScript('//apis.google.com/js/client.js?onload=load',function(){
            console.log('Loading Google API')
        });

        console.log(response)

    });

  });

});

signin2 = function(){FB.login(function(response){}, {scope: 'manage_pages'})} //{scope: 'email,manage_pages,publish_actions'}
$('#loginButton').val('Login').off('click').on('click', signin2)

load = function() {
    var http = ( window.location.hostname == "localhost" ? "http://localhost:8081" : "https://"+window.location.host )
    gapi.client.load('service', 'v0', start, http+'/_ah/api')
}

            //var apisToLoad = 2;
            //var callback = function() { if (--apisToLoad == 0) {autosignin()} }
            //var http = ( window.location.hostname == "localhost" ? "http://" : "https://" )
            //gapi.client.load('service', 'v0', start, http+window.location.host+'/_ah/api')
            //gapi.client.load('oauth2', 'v2', function(){});

/*
signin = function() {
    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', signout2);
    b.addEventListener('click',signin2)
    b.value="Sign in"
    b.style.display="block";
}

signout = function() {
    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', signin2)
    b.addEventListener('click',signout2)
    b.value="Sign out"
    b.style.display="block";
}
*/

        //if (response.status === 'connected') {}
        //else if (response.status === 'not_authorized') {alert('Not authorized.')}
        //else {}

                //if (!FB.getAccessToken()){signin()}

                //signout2 = function(){FB.logout(); stop();}
