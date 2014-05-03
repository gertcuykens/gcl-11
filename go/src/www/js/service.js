service = {}

service.publish = function() {
  FB.api(
   document.getElementById('feed').value,
   'post',
   {message: document.getElementById('message').value},
   function(response){console.log(response)}
  )
};

service.submit = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  console.log(t);
  gapi.auth.setToken(t)
  console.log(gapi.auth.getToken());
  gapi.client.service.datastore.submit({"list":[{"message":document.getElementById('message').value}]}).execute(function(resp){service.list()})
};

service.list = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  console.log(t);
  gapi.auth.setToken(t)
  console.log(gapi.auth.getToken());
  gapi.client.service.datastore.list().execute(
      function(resp) {
        if (!resp.code) {
          resp.list = resp.list || [];
          document.getElementById('console').innerHTML=""
          for (var i = 0; i < resp.list.length; i++) {print(resp.list[i]);}
        }
      });
};

print = function(s) {
  var element = document.createElement('li');
  //element.classList.add('row');
  element.innerHTML = s.message;
  document.getElementById('console').appendChild(element);
};

testAPI = function () {
  console.log('Browser Facebook, Fetching your information... ');
  FB.api('/me?fields=email', function(response) {console.log('Browser Facebook, '+response.email+'.')})

  console.log('Server Facebook, Fetching your information... ');
  Facebook.access_token=FB.getAccessToken()
  gapi.client.service.facebook.callback(Facebook).execute(function(response){console.log('Server Facebook, '+response.email_token)})
}

/*
function testAPI2() {
  console.log('Browser Google, Fetching your information... ');
  gapi.client.oauth2.userinfo.get().execute(function(response) {console.log('Browser Google, '+response.email+'.')})

  console.log('Server Google, Fetching your information... ');
  Google.access_token=gapi.auth.getToken().access_token
  gapi.client.rest.google.callback().execute(function(response){console.log('Server Google, '+response.result.message)})
}

function testAPI2() {
    console.log('Iab, Fetching your order information... ');
    gapi.client.rest.google.purchases().execute(function(response){console.log('Iab, '+response.message)})
}

function testAPI3() {
    console.log('Storage, setting ACL...');
    gapi.client.rest.google.storage().execute(function(response){console.log('Storage, '+response.message)})
}
*/