Google=new Token()
Google.type_token="google"

init = function () {
    console.log('Loading Google')
    var apisToLoad = 2;
    var callback = function() { if (--apisToLoad == 0) {autosignin()}}
    var http = ( window.location.hostname == "localhost" ? "http://" : "https://" )
    gapi.client.load('rest', 'v0', callback, http+window.location.host+'/_ah/api')
    gapi.client.load('oauth2', 'v2', callback);
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
        clientid : '781166019436-vs3a5b1nva8kefmsk1mscccur6rkpos4.apps.googleusercontent.com',
        requestvisibleactions : 'http://schemas.google.com/AddActivity',
        cookiepolicy : 'single_host_origin',
        scope: 'https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/plus.login'
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
        client_id: "781166019436-vs3a5b1nva8kefmsk1mscccur6rkpos4.apps.googleusercontent.com",
        scope: "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/plus.login",
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

