service.getTrickList = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.trickstore.getTrickList().execute(function(resp){
    $('#trick').html('')
    var s = resp.list || []
    for (var i = 0; i < s.length; i++) {
      var option = document.createElement('option');
      option.innerHTML= s[i].name
      $('#trick').append(option)
    }
    //console.log(resp.list)
  })
};

service.addTrickName = function(trick) {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.trickstore.putTrickName({"list":[{
    "name":trick,
    "difficulty":10 //parseInt($('#heat').val(), 10),
  }]}).execute(function(resp){})
};

service.deleteTrickName = function(trick) {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.trickstore.deleteTrickName({"list":[{
    "name":trick
  }]}).execute(function(resp){})
};

$('#trickAdd').on('click',function (){
  var t = prompt("add trick")
  var option = document.createElement('option');
  option.innerHTML=t
  $('#trick').append(option)
  $('#trick').val(t);
  service.addTrickName(t);
})

$('#trickDel').on('click',function (){
  var t = $('#trick option:selected').text()
  $('#trick option:selected').remove()
  //$("this option[value='X']").remove();
  service.deleteTrickName(t);
})
