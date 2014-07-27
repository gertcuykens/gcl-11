unit = {}

unit.user = function () {
  console.log('Browser Facebook, Fetching your information... ');
  FB.api('/me?fields=email', function(response) {console.log('Browser Facebook, '+response.email+'.')})
}

unit.truncate = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.truncate({"list":[{"event":"Tarifa 2014"}]}).execute(function(resp){service.list()})
};

unit.get = function(id) {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.get({"list":[{"event":$('#event').val(),"id":id}]}).execute(
      function(resp) {
        if (!resp.code) {
            resp.list = resp.list || []
            console.log(resp.list)
            //if (resp.list[0]){heatf(resp.list[0])}
        }
      }
  );
};

unit.data = function() {
  //var d = new Date()
  //d.parseRfc3339("0001-01-01T00:00:00Z");
  //"date":d,
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.put({"list":[

    {
      "judge":"test-judge1",
      "event":"Tarifa 2014",
      "division":"woman",
      "heat":1,
      "rider":"Gert",
      "trick":"POP",
      "score":0,
      "attempt":1
    },

    {
      "judge":"test-judge2",
      "event":"Tarifa 2014",
      "division":"woman",
      "heat":1,
      "rider":"Gert",
      "trick":"POP",
      "score":0,
      "attempt":1
    },

    {
        "judge":"test-judge1",
        "event":"Tarifa 2014",
        "division":"woman",
        "heat":1,
        "rider":"Annelous",
        "trick":"BJ",
        "score":5,
        "attempt":1
    },

    {
        "judge":"test-judge2",
        "event":"Tarifa 2014",
        "division":"woman",
        "heat":1,
        "rider":"Annelous",
        "trick":"BJ",
        "score":6,
        "attempt":1
    },

    {
        "judge":"test-judge1",
        "event":"Tarifa 2014",
        "division":"woman",
        "heat":1,
        "rider":"Annelous",
        "trick":"BJ",
        "score":4,
        "attempt":2
    },

    {
        "judge":"test-judge2",
        "event":"Tarifa 2014",
        "division":"woman",
        "heat":1,
        "rider":"Annelous",
        "trick":"BJ",
        "score":5,
        "attempt":2
    },

    {
        "judge":"test-judge1",
        "event":"Tarifa 2014",
        "division":"woman",
        "heat":1,
        "rider":"Annelous",
        "trick":"BM",
        "score":6,
        "attempt":3
    },

    {
        "judge":"test-judge2",
        "event":"Tarifa 2014",
        "division":"woman",
        "heat":1,
        "rider":"Annelous",
        "trick":"BM",
        "score":7,
        "attempt":3
    }

  ]}).execute(function(resp){service.list()})

  service.addTrickName('BJ')
  service.addTrickName('KGB')
  service.addTrickName('HB')
  service.addTrickName('POP')

};

/*
http://jsfiddle.net/spetnik/gFzCk/1953/

truncateAPI = function () {
  console.log('Browser Facebook, Fetching your information... ');
  FB.api('/me?fields=email', function(response) {console.log('Browser Facebook, '+response.email+'.')})

  //console.log('Server Facebook, Fetching your information... ');
  //Facebook.access_token=FB.getAccessToken()
  //gapi.client.service.facebook.callback(Facebook).execute(function(response){console.log('Server Facebook, '+response.email_token)})
}

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

                 /*
                 var x = document.URL.match(/[^/]*$/)
                 switch(x[0]) {
                     case "":
                         for (var i = 0; i < resp.list.length; i++) {print1(resp.list[i]);}
                         break;
                     case "result":
                         print2(resp.list);
                         break;
                 }
                 */

/*
service.submit2 = function() {
  var d = new Date()
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.put2({"list":[{
    "event":$('#event').val(),
    "heat":$('#heat').val(),
  }]}).execute(function(resp){service.list2()})
};
*/

/*
service.list2 = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.get2({}).execute(
      function(resp) {
        if (!resp.code) {
            resp.list = resp.list || []
            console.log(resp.list)
        }
      }
  );
};
*/
