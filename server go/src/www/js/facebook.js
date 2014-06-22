FACEBOOK_CLIENT_ID="249348058430770"

$(document).ready(function() {
  $.ajaxSetup({ cache: true });
  $.getScript('//connect.facebook.net/en_UK/all.js', function(){

    FB.init({
        appId: FACEBOOK_CLIENT_ID,
        version: 'v2.0'
    });

    FB.Event.subscribe('auth.authResponseChange', function(response) {
        if (response.status === 'connected') {signout()}
        else if (response.status === 'not_authorized') {alert('Not authorized.')}
        else {signin()}
        console.log(response)
    });

    FB.getLoginStatus(function(response){
        $.getScript('//apis.google.com/js/client:plusone.js?onload=load');
        if (!FB.getAccessToken()){signin()}
        console.log(response)
    });

  });

});

signin = function() {
    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', fsignout);
    b.addEventListener('click',fsignin)
    b.value="Sign in"
    b.style.display="block";
}

signout = function() {
    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', fsignin)
    b.addEventListener('click',fsignout)
    b.value="Sign out"
    b.style.display="block";
}

fsignin = function() {FB.login(function(response){start()}, {scope: 'email,publish_actions,manage_pages'})}

fsignout = function() {FB.logout(); stop();}

load = function () {
    console.log('Loading Google API')
    //var apisToLoad = 2;
    //var callback = function() { if (--apisToLoad == 0) {autosignin()} }
    var http = ( window.location.hostname == "localhost" ? "http://" : "https://" )
    gapi.client.load('service', 'v0', start, http+window.location.host+'/_ah/api')
    //gapi.client.load('oauth2', 'v2', function(){});
};

/*
$(document).ready(function() {
  $.ajaxSetup({ cache: true });
  $.getScript('//connect.facebook.net/en_UK/all.js', function(){
    FB.init({appId: FACEBOOK_CLIENT_ID,});
    $('#loginbutton,#feedbutton').removeAttr('disabled');
    FB.getLoginStatus(updateStatusCallback);
  });
});
*/
