GOOGLE_CLIENT_ID="522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com"
GOOGLE_SCOPE="https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/plus.login"

Google=new Token()
Google.type_token="google"

init = function () {
    console.log('Loading Google API')
    //var apisToLoad = 2;
    //var callback = function() { if (--apisToLoad == 0) {autosignin()} }
    var http = ( window.location.hostname == "localhost" ? "http://" : "https://" )
    gapi.client.load('service', 'v0', service.list, http+window.location.host+'/_ah/api')
    //gapi.client.load('oauth2', 'v2', function(){});
};

(function(d,s) {
    var js = d.createElement('script');
    js.id = 'google-jssdk';
    js.async = true;
    js.src = '//apis.google.com/js/client:plusone.js?onload=init';
    s.parentNode.insertBefore(js, s);
})(document,document.getElementsByTagName('script')[0]);

signin = function() {
    var options = {
        callback : autosignin,
        clientid : GOOGLE_CLIENT_ID,
        requestvisibleactions : 'http://schemas.google.com/AddActivity',
        cookiepolicy : 'single_host_origin',
        scope: GOOGLE_SCOPE
    }
    gapi.auth.signIn(options)
}

autosignin = function() {
    var b=document.getElementById('gsigninButton')
    b.removeEventListener('click', signout);
    b.addEventListener('click',signin)
    //document.getElementsByClassName("buttonText")[1].innerHTML='Log In'

    var callback = function (t) {
        if (!t) return false
        var b=document.getElementById("gsigninButton")
        b.removeEventListener('click', signin);
        b.addEventListener('click', signout)
        //document.getElementsByClassName("buttonText")[1].innerHTML='Log Out'
        console.log('Sign-in state: '+ t['status']['signed_in'])
        Google.access_token=gapi.auth.getToken().access_token
        border()
    }

    var options = {
        client_id: GOOGLE_CLIENT_ID,
        scope: GOOGLE_SCOPE,
        immediate: true
    }

    gapi.auth.authorize(options, callback)
}

signout = function () {
    Google.access_token=gapi.auth.getToken().access_token
    gapi.client.rest.google.revoke(Google).execute(function(response){console.log('Server, Bye, '+response.message);Google.access_token=null;border()})
    var b=document.getElementById('gsigninButton')
    b.removeEventListener('click', signout)
    b.addEventListener('click',signin)
    //document.getElementsByClassName("buttonText")[1].innerHTML='Log In'
}

