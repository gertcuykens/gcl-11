service.getTrickList = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.trickstore.getTrickList().execute(function(resp){})
};

service.addTrickName = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.trickstore.putTrickName({"list":[{
    "name":$('#trick').val(),
    "difficulty":10 //parseInt($('#heat').val(), 10),
  }]}).execute(function(resp){})
};

service.deleteTrickName = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.trickstore.deleteTrickName({"list":[{
    "name":$('#trick').val()
  }]}).execute(function(resp){})
};

service.jsTrickList = function(trick) {
  var option = document.createElement('option');
  option.innerHTML=trick
  $('#trick').append(option)
};

$('#trickAdd').on('click',function (){
  var t = prompt("add trick")
  service.jsTrickList(t)
  //service.addTrickName()
})
