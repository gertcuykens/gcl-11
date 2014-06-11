service = {}

service.truncate = function() {
  //t.access_token = FB.getAccessToken()
  //gapi.auth.setToken(t)
  gapi.client.service.datastore.truncate().execute(function(resp){service.list()})
};

service.publish = function() {
  FB.api(
   'gcl11/feed', //document.getElementById('feed').value,
   'post',
   {message: document.getElementById('message').value},
   function(response){console.log(response)}
  )
};

service.submit = function() {
  //parseRfc3339("0001-01-01T00:00:00Z");
  var d = new Date()
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.put({"list":[{
    "id":0,
    "date":d,
    "user":"",
    "heat":0,
    "event":"",
    "rider":$('#rider').val(),
    "trick":$('#trick').val(),
    "score":parseInt($("#score").val(), 10)
  }]}).execute(function(resp){})
};

service.list = function() {
  document.getElementById('console').innerHTML="Loading..."
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.getall().execute(
      function(resp) {
        if (!resp.code) {
          resp.list = resp.list || [];
          document.getElementById('console').innerHTML=""
          for (var i = 0; i < resp.list.length; i++) {print(resp.list[i]);}
        }
      }
  );
};

print = function(s) {
  var row = document.createElement('tr');
  //row.onclick=form

  var rider = document.createElement('td');
  rider.innerHTML=s['rider']
  row.appendChild(rider);

  var trick = document.createElement('td');
  trick.innerHTML=s['trick']
  row.appendChild(trick);

  var score = document.createElement('td');
  score.innerHTML=s['score']
  row.appendChild(score);

  var judge = document.createElement('td');
  judge.innerHTML=s['judge']
  row.appendChild(judge);

  document.getElementById('console').appendChild(row);
};

init = function () {
    console.log('Loading Google API')
    //var apisToLoad = 2;
    //var callback = function() { if (--apisToLoad == 0) {autosignin()} }
    var http = ( window.location.hostname == "localhost" ? "http://" : "https://" )
    gapi.client.load('service', 'v0', service.list, http+window.location.host+'/_ah/api')
    //gapi.client.load('oauth2', 'v2', function(){});
};

start = function() {
    var s = document.getElementsByTagName('script')[0]
    var j = document.createElement('script');
    j.id = 'google-jssdk';
    j.async = true;
    j.src = '//apis.google.com/js/client:plusone.js?onload=init';
    s.parentNode.insertBefore(j, s);
};

stop = function() {document.getElementById('console').innerHTML=""};

form = function(x) {
    switch(x){
        case 1: $( "#form1" ).show(),$( "#form2" ).hide(); break;
        case 2: $( "#form1" ).hide(),$( "#form2" ).show(); break;
    }
}

testAPI = function () {
  console.log('Browser Facebook, Fetching your information... ');
  FB.api('/me?fields=email', function(response) {console.log('Browser Facebook, '+response.email+'.')})

  //console.log('Server Facebook, Fetching your information... ');
  //Facebook.access_token=FB.getAccessToken()
  //gapi.client.service.facebook.callback(Facebook).execute(function(response){console.log('Server Facebook, '+response.email_token)})
}

truncate = function () {
  console.log('Browser Facebook, Fetching your information... ');
  FB.api('/me?fields=email', function(response) {console.log('Browser Facebook, '+response.email+'.')})

  //console.log('Server Facebook, Fetching your information... ');
  //Facebook.access_token=FB.getAccessToken()
  //gapi.client.service.facebook.callback(Facebook).execute(function(response){console.log('Server Facebook, '+response.email_token)})
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

 /*
  for (var x in s) {
      var td = document.createElement('td');
      if (x=="date") {td.innerHTML = new Date(s.date).toLocaleString()}
      else {td.innerHTML = s[x]}
      tr.appendChild(td);
  }*/


      //element.classList.add('row');
