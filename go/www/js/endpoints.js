Google=new Token()
Google.type_token="google"

init = function () {
    var apisToLoad = 2;
    var callback = function() {if (--apisToLoad == 0) {autosignin()}}
    gapi.client.load('rest', 'v0', callback, '//'+window.location.host+'/_ah/api');
    gapi.client.load('oauth2', 'v2', callback);
};

(function(d,s) {

    var js2 = d.createElement('script');
    js2.id = 'google-jssdk';
    js2.async = true;
    js2.src = '//apis.google.com/js/client:plusone.js?onload=init';
    s.parentNode.insertBefore(js2, s);

})(document,document.getElementsByTagName('script')[0]);

function signin() {
    var options = {
        callback : autosignin,
        clientid : '1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com',
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
    document.getElementsByClassName("buttonText")[0].innerHTML='Log In'

    var callback = function (t) {
        if (!t) return false
        var b=document.getElementById("gsigninButton")
        b.removeEventListener('click', signin);
        b.addEventListener('click', signout)
        document.getElementsByClassName("buttonText")[0].innerHTML='Log Out'
        console.log('Sign-in state: '+ t['status']['signed_in'])
        Google.access_token=gapi.auth.getToken().access_token
        border()
    }

    var options = {
        client_id: "1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com",
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
    document.getElementsByClassName("buttonText")[0].innerHTML='Log In'
}

$('#logindropdown').on('hidden.bs.dropdown', function () {border()})
$('#logindropdown').on('shown.bs.dropdown', function () {border()})

function border() {
    var b=document.getElementById('login')
    if (document.getElementById('login-menu').classList.contains('open')){b.style.borderBottom='1px solid transparent'; return}
    if (Google.access_token) {b.style.borderBottom='1px solid #dd4b39'; return}
    b.style.borderBottom='1px solid transparent'
}

function testAPI1() {
    console.log('Browser, Fetching your Google information... ');
    gapi.client.oauth2.userinfo.get().execute(function(response) {console.log('Google, '+response.email+'.')})
}

function testAPI2() {
    console.log('Server, Fetching your Google information... ');
    gapi.client.rest.google.user().execute(function(response){console.log('Server, '+response.message)})
}

function testAPI3() {
    console.log('Iab, Fetching your order information... ');
    gapi.client.rest.google.purchases().execute(function(response){console.log('Iab, '+response.message)})
}

function testAPI4() {
    console.log('Storage, setting ACL...');
    gapi.client.rest.google.storage().execute(function(response){console.log('Storage, '+response.message)})
}
