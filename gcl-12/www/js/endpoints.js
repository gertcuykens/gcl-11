window.fbAsyncInit = function() {
    FB.Event.subscribe('auth.authResponseChange', function(response) {
        if (response.status === 'connected') {}
        else if (response.status === 'not_authorized') {FB.login();}
        else {}
    });
};

init = function () {
    var apisToLoad = 2;
    var callback = function() {if (--apisToLoad == 0) {autosignin()}}
    gapi.client.load('rest', 'v0', callback, '//'+window.location.host+'/_ah/api');
    gapi.client.load('oauth2', 'v2', callback);
};

(function(d,s) {
    var js1 = d.createElement('script');
    js1.id = 'facebook-jssdk';
    js1.async = true;
    js1.src = "//connect.facebook.net/en_US/all.js#xfbml=1&appId=1379351942320920";
    d.getElementById('fb-root').appendChild(js1);

    var js2 = d.createElement('script');
    js2.id = 'google-jssdk';
    js2.async = true;
    js2.src = '//apis.google.com/js/client:plusone.js?onload=init';
    s.parentNode.insertBefore(js2, s);

})(document,document.getElementsByTagName('script')[0]);

function signin() {
    var options = { callback : autosignin,
        clientid : '1034966141188-b4cup6jccsjqpdc14c9218fhb488e515.apps.googleusercontent.com',
        requestvisibleactions : 'http://schemas.google.com/AddActivity',
        cookiepolicy : 'single_host_origin',
        scope: 'https://www.googleapis.com/auth/plus.login https://www.googleapis.com/auth/userinfo.email'}
    gapi.auth.signIn(options)
}

function autosignin() {
    var b=document.getElementById('signinButton')
    b.removeEventListener('click', signout);
    b.addEventListener('click',signin)
    document.getElementsByClassName("buttonText")[0].innerHTML='Log In'

    var callback = function (token) { if (!token) return false
        var b=document.getElementById("signinButton")
        b.removeEventListener('click', signin);
        b.addEventListener('click', signout)
        document.getElementsByClassName("buttonText")[0].innerHTML='Log Out'
        console.log('Sign-in state: '+ token['status']['signed_in']) }

    var options = { client_id: "1034966141188-b4cup6jccsjqpdc14c9218fhb488e515.apps.googleusercontent.com",
        scope: "https://www.googleapis.com/auth/plus.login https://www.googleapis.com/auth/userinfo.email",
        immediate: true }

    gapi.auth.authorize(options, callback)
}

function signout() {
    gapi.client.rest.logout({token:gapi.auth.getToken().access_token}).execute(function(response){console.log('Server, Bye, '+response.message)})
    var b=document.getElementById('signinButton')
    b.removeEventListener('click', signout);
    b.addEventListener('click',signin)
    document.getElementsByClassName("buttonText")[0].innerHTML='Log In'
}

function testAPI() {
    console.log('Facebook,  Fetching your information.... ');
    FB.api('/me', function(response) {console.log('Facebook, Good to see you, '+response.name+'.')})
    console.log('Google,  Fetching your information.... ');
    gapi.client.oauth2.userinfo.get().execute(function(response) {console.log('Google, Good to see you, '+response.email+'.')})
    console.log('Server, Fetching your information.... ');
    gapi.client.rest.login({token:FB.getAccessToken()}).execute(function(response){console.log('Server, Good to see you, '+response.message)})
}