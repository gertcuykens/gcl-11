FACEBOOK_CLIENT_ID="249348058430770"

Facebook=new Token()
Facebook.type_token="facebook"

window.fbAsyncInit = function() {

    FB.Event.subscribe('auth.authResponseChange', function(response) {
        if (response.status === 'connected') {signout()}
        else if (response.status === 'not_authorized') {alert('Not authorized.')}
        else {signin()}
        console.log(response)
    })

    FB.getLoginStatus(function(response){
        start()
        if (!FB.getAccessToken()) signin()
        console.log(response)
    })

}

(function(d,s) {
    console.log('Loading Facebook')
    var js = d.createElement('script');
    js.id = 'facebook-jssdk';
    js.async = true;
    js.src = "//connect.facebook.net/en_US/all.js#xfbml=0&appId="+FACEBOOK_CLIENT_ID;
    //d.getElementById('fb-root').appendChild(js);
    s.parentNode.insertBefore(js, s);
})(document,document.getElementsByTagName('script')[0]);

signin = function() {
    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', fsignout);
    b.addEventListener('click',fsignin)
    b.value="Sign in"
    b.style.display="block";
    Facebook.access_token=null
}

fsignin = function() {FB.login(function(response){service.list()}, {scope: 'email,publish_actions,manage_pages'})}

signout = function() {
    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', fsignin)
    b.addEventListener('click',fsignout)
    b.value="Sign out"
    b.style.display="none";
    Facebook.access_token=FB.getAccessToken()
}

fsignout = function() {FB.logout(); Facebook.access_token=null; stop()}

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

