FACEBOOK_CLIENT_ID="249348058430770"

Facebook=new Token()
Facebook.type_token="facebook"

window.fbAsyncInit = function() {

    FB.Event.subscribe('auth.authResponseChange', function(response) {
        if (response.status === 'connected') {
            var b=document.getElementById('fsigninButton')
            b.removeEventListener('click', fsignin);
            b.addEventListener('click',fsignout)
            //document.getElementsByClassName("buttonText")[0].innerHTML='Log Out'
            Facebook.access_token=FB.getAccessToken()
        }
        else if (response.status === 'not_authorized') {}
        else {
            var b=document.getElementById('fsigninButton')
            b.removeEventListener('click', fsignout);
            b.addEventListener('click',fsignin)
            //document.getElementsByClassName("buttonText")[0].innerHTML='Log In'
            Facebook.access_token=null
        }
        border()
    });

    var b=document.getElementById('fsigninButton')
    b.removeEventListener('click', fsignout);
    b.addEventListener('click',fsignin)
    //document.getElementsByClassName("buttonText")[0].innerHTML='Log In'
};

(function(d,s) {
    console.log('Loading Facebook')
    var js = d.createElement('script');
    js.id = 'facebook-jssdk';
    js.async = true;
    js.src = "//connect.facebook.net/en_US/all.js#xfbml=0&appId="+FACEBOOK_CLIENT_ID;
    //d.getElementById('fb-root').appendChild(js);
    s.parentNode.insertBefore(js, s);
})(document,document.getElementsByTagName('script')[0]);

fsignin = function() {FB.login(function(response){}, {scope: 'email,publish_actions,manage_pages'})}

fsignout = function() {FB.logout(); Facebook.access_token=null}

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