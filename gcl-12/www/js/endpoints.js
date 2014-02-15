Facebook=new Token()
Facebook.type_token="facebook"

Google=new Token()
Google.type_token="google"

Server=new Token()
Server.type_token="server"

window.fbAsyncInit = function() {
    FB.Event.subscribe('auth.authResponseChange', function(response) {
        if (response.status === 'connected') {
            var b=document.getElementById('fsigninButton')
            b.removeEventListener('click', fsignin);
            b.addEventListener('click',fsignout)
            document.getElementsByClassName("buttonText")[0].innerHTML='Log Out'
            Facebook.access_token=FB.getAccessToken()
        }
        else if (response.status === 'not_authorized') {}
        else {
            var b=document.getElementById('fsigninButton')
            b.removeEventListener('click', fsignout);
            b.addEventListener('click',fsignin)
            document.getElementsByClassName("buttonText")[0].innerHTML='Log In'
            Facebook.access_token=null
        }
        border()
    });

    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', fsignout);
    b.addEventListener('click',fsignin)
    document.getElementsByClassName("buttonText")[0].innerHTML='Log In'
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

function fsignin() {FB.login(function(response){}, {scope: 'email,user_likes'})}

function fsignout() {FB.logout()}

function signin() {
    var options = {
        callback : autosignin,
        clientid : '1034966141188-b4cup6jccsjqpdc14c9218fhb488e515.apps.googleusercontent.com',
        requestvisibleactions : 'http://schemas.google.com/AddActivity',
        cookiepolicy : 'single_host_origin',
        scope: 'https://www.googleapis.com/auth/plus.login https://www.googleapis.com/auth/userinfo.email'
    }
    gapi.auth.signIn(options)
}

function autosignin() {
    var b=document.getElementById('gsigninButton')
    b.removeEventListener('click', signout);
    b.addEventListener('click',signin)
    document.getElementsByClassName("buttonText")[1].innerHTML='Log In'

    var callback = function (t) {
        if (!t) return false
        var b=document.getElementById("gsigninButton")
        b.removeEventListener('click', signin);
        b.addEventListener('click', signout)
        document.getElementsByClassName("buttonText")[1].innerHTML='Log Out'
        console.log('Sign-in state: '+ t['status']['signed_in'])
        Google.access_token=gapi.auth.getToken().access_token
        border()
    }

    var options = {
        client_id: "1034966141188-b4cup6jccsjqpdc14c9218fhb488e515.apps.googleusercontent.com",
        scope: "https://www.googleapis.com/auth/plus.login https://www.googleapis.com/auth/userinfo.email",
        immediate: true
    }

    gapi.auth.authorize(options, callback)
}

function signout() {
    gapi.client.rest.google.revoke(Google).execute(function(response){console.log('Server, Bye, '+response.message);Google.access_token=null;border()})
    var b=document.getElementById('gsigninButton')
    b.removeEventListener('click', signout);
    b.addEventListener('click',signin)
    document.getElementsByClassName("buttonText")[1].innerHTML='Log In'
}

$('#logindropdown').on('hidden.bs.dropdown', function () {border()})
$('#logindropdown').on('shown.bs.dropdown', function () {border()})

function border() {
    var b=document.getElementById('login')
    if (document.getElementById('login-menu').classList.contains('open')){b.style.borderBottom='1px solid transparent'; return}
    if (Google.access_token) {b.style.borderBottom='1px solid #dd4b39'; return}
    if (Facebook.access_token) {b.style.borderBottom='1px solid #5f78ab'; return}
    b.style.borderBottom='1px solid transparent'
}

function testAPI1() {
    console.log('Browser, Fetching your Facebook information... ');
    FB.api('/me?fields=email', function(response) {console.log('Facebook, '+response.email+'.')})

    console.log('Browser, Fetching your Google information... ');
    gapi.client.oauth2.userinfo.get().execute(function(response) {console.log('Google, '+response.email+'.')})
}

function testAPI2() {
    console.log('Facebook, Fetching your Facebook information... ');
    gapi.client.rest.facebook.callback(Facebook).execute(function(response){console.log('Facebook, '+response.email_token)})

    console.log('Google, Fetching your Google information... ');
    gapi.client.rest.google.callback().execute(function(response){console.log('Google, '+response.email_token)})
}

function testAPI3() {
    console.log('Server, Register Facebook user... ');
    gapi.client.rest.register(Facebook).execute(function(response){console.log('Server, '+JSON.stringify(response.result));Server=response.result})

    console.log('Server, Register Google user... ');
    gapi.client.rest.register(Google).execute(function(response){console.log('Server, '+JSON.stringify(response.result));Server=response.result})
}

function testAPI4() {
    console.log('Server, CheckSum Facebook token... ');
    gapi.client.rest.checksum(Server).execute(function(response){console.log('Server, '+JSON.stringify(response.result));Server=response.result})
}
