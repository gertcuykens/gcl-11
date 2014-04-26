function testAPI() {
  console.log('Browser Facebook, Fetching your information... ');
  FB.api('/me?fields=email', function(response) {console.log('Browser Facebook, '+response.email+'.')})

  console.log('Browser Google, Fetching your information... ');
  gapi.client.oauth2.userinfo.get().execute(function(response) {console.log('Browser Google, '+response.email+'.')})

  console.log('Server Facebook, Fetching your information... ');
  Facebook.access_token=FB.getAccessToken()
  gapi.client.rest.facebook.callback(Facebook).execute(function(response){console.log('Server Facebook, '+response.email_token)})

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

print = function(greeting) {
  var element = document.createElement('span');
  element.classList.add('row');
  element.innerHTML = greeting.message;
  document.getElementById('outputLog').appendChild(element);
};

getGreeting = function(id) {
  gapi.client.rest.greetings.getGreeting({'id': +id}).execute(
      function(resp) {
        if (!resp.code) {print(resp);}
      });
};

listGreeting = function() {
  gapi.client.rest.greetings.listGreeting().execute(
      function(resp) {
        if (!resp.code) {
          resp.items = resp.items || [];
          for (var i = 0; i < resp.items.length; i++) {print(resp.items[i]);}
        }
      });
};

multiplyGreeting = function(greeting, times) {
  gapi.client.rest.greetings.multiply({
      'message': greeting,
      'times': +times
    }).execute(function(resp) {
      if (!resp.code) {print(resp);}
    });
};

authedGreeting = function() {
  gapi.client.rest.greetings.authed().execute(function(resp) {print(resp);});
};

soap = function() {
  gapi.client.rest.greetings.soap().execute(function(resp) {print(resp);});
};

datastore = function() {
  gapi.client.rest.greetings.datastore().execute(function(resp) {print(resp);});
};

document.getElementById('getGreeting').onclick = function() {getGreeting(document.getElementById('id').value);}
document.getElementById('listGreeting').onclick = function() {listGreeting();}
document.getElementById('multiplyGreetings').onclick = function() {multiplyGreeting(document.getElementById('greeting').value, document.getElementById('count').value);}
document.getElementById('authedGreeting').onclick = function() {authedGreeting();}
document.getElementById('soap').onclick = function() {soap();}
document.getElementById('datastore').onclick = function() {datastore();}

