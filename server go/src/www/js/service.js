service = {}

service.truncate = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.truncate().execute(function(resp){service.list()})
};

service.delete = function() {
 var i=0;
 var list = []
 $(this).closest('table').children('tbody').children('tr').each(function(){
  if (!$(this).hasClass('active')) {return true}
  list[i++]={'id':$(this).attr('id'), 'event':$('#event').val()}
  //list[i++]={'id':parseInt($(this).attr('id'))}
 })

  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.delete({'list':list}).execute(function(resp){service.list()})
};

service.editor = function() {
  console.log('Server Editor, Fetching your information... ');
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.editor().execute(function(response){
    console.log(response)
    if (!response.error){ view(2) } // else { form(1) }
  })
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
    //"id":0,
    //"date":d,
    //"judge":"",
    "event":$('#event').val(),
    "division":$('#division').val(),
    "heat":parseInt($('#heat').val(), 10),
    "rider":$('#rider').val(),
    "trick":$('#trick').val(),
    "score":parseInt($("#score").val(), 10),
    "attempt":parseInt($("#attempt").val(), 10)
  }]}).execute(function(resp){service.list()})
};

service.list = function() {
  document.getElementById('console2').innerHTML="Loading..."
  document.getElementById('console3').innerHTML="Loading..."
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.getHeat({"list":[{
                                            "event":$('#event').val(),
                                            "division":$('#division').val(),
                                            "heat":parseInt($('#heat').val(), 10),
                                          }]}
  ).execute(
      function(resp) {
        if (!resp.code) {
            resp.list = resp.list || []
            document.getElementById('console2').innerHTML=""
            document.getElementById('console3').innerHTML=""
            mapreduce(resp.list);
        }
      }
  );
};

service.getFirst = function() {
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.getFirst({"list":[{"event":$('#event').val(),}]}).execute(
      function(resp) {
        if (!resp.code) {
            resp.list = resp.list || []
            if (resp.list[0]){heatf(resp.list[0])}
        }
      }
  );
};

act = function () {
 if (this.className=="active") {this.className="";return;}
 for (r in $(this).parent('tbody').children('tr')){$(this).parent('tbody').children('tr')[r].className=""}
 this.className="active"
}

heatf = function (s) {
 $('#heat2').html(s['event']+' '+s['division']+' heat '+s['heat'])
 $('#heat3').html(s['event']+' '+s['division']+' heat '+s['heat'])
 $('#event').attr('value',s['event'])
 $('#division').val(s['division'])
 $('#heat').attr('value',s['heat'])
 service.list()
}

heatb = function (v) {
 var e=$('#event').attr('value')
 var d=$('#division').val()
 var h=parseInt($('#heat').attr('value'))
 heatf({'event':e,'division':d,'heat':h+v})
}

division = function () {
 var e=$('#event').attr('value')
 var h=parseInt($('#heat').attr('value'))
 var d=$('#division').val()
 switch (d) {
  case "men": $('#division').val('woman'); break;
  default: $('#division').val('men'); break;
 }
 d=$('#division').val()
 heatf({'event':e,'division':d,'heat':h})
}

mapreduce = function(s) {
  var object = {}
  for (var i = 0; i < s.length; i++) {
    var id=s[i]['id']
    var rider=s[i]['rider']
    var trick=s[i]['trick']
    var attempt=s[i]['attempt']
    var score=s[i]['score']
    var judge=s[i]['judge']
    if (!object[rider]) {object[rider]=[];print4(rider);}
    if (!object[rider][attempt]) {object[rider][attempt]={}}
    if (!object[rider][attempt][judge]) {object[rider][attempt][judge]=[score,trick,id]}
  }
  console.log(object)

   var trick={}
   for (rider in object) {
    if(object.hasOwnProperty(rider)){

         var list={}

         if (!trick[rider]) {trick[rider]=[]}
         for (attempt in object[rider]) {
           //console.log(attempt)

           var name=""
           var max=0
           var score=0
           var score2=[]
           for (judge in object[rider][attempt]) {
            if(object[rider][attempt].hasOwnProperty(judge)){
             max+=10
             scorej=object[rider][attempt][judge][0]
             score+=object[rider][attempt][judge][0]
             score2.push(object[rider][attempt][judge][0])
             name=object[rider][attempt][judge][1]
             id=object[rider][attempt][judge][2]
             print3(id,rider,name,scorej,attempt,judge)
             if (!list[name]){list[name]=[score,max,name,score2]}
             else if (list[name][0]<score){list[name]=[score,max,name,score2]}
            }
           }

           //console.log(name+':'+score+'/'+max)

         }

         var keys = Object.keys(list).sort(function(a,b){return list[b][0]-list[a][0]})
         for (k in keys) {
         trick[rider].push(list[keys[k]])


         }

     }
    }
    console.log(trick)

    for (rider in trick){
     if(trick.hasOwnProperty(rider)){


               for (score in trick[rider]){

  print1(rider,trick[rider][score][2],trick[rider][score][0],trick[rider][score][3])

               }

     }
    }

     var score={}
     for (rider in trick) {
       if (trick.hasOwnProperty(rider)) {
        var count=0
        var total=0
        for (t in trick[rider]) {total+=trick[rider][t][0]}
        count++
        if (count>5){break}
        score[rider]=total
       }
     }
     console.log(score)

     var keys=Object.keys(score).sort(function(a,b){return score[b]-score[a]})

     var rider
     if (keys[0]) rider=keys[0]
     for (k in keys) {

         var rider2= keys[k]

          document.getElementById("-"+rider2  ).innerHTML=" score "+score[rider2]
          $('#--'+rider2).data('score',score[rider2])
          $('#'+rider2).data('score',score[rider2])
          document.getElementById("---"+rider2).innerHTML=" score "+score[rider2]

          if (k=='0' && keys[parseInt(k)+1]) {
                 rider2= keys[parseInt(k)+1]
                 console.log(rider)
                 console.log(rider2)
                 document.getElementById("-"+rider  ).innerHTML+=" difference "+(score[rider]-score[rider2])
                 document.getElementById("---"+rider).innerHTML+=" difference "+(score[rider]-score[rider2])
          } else {
                 document.getElementById("-"+rider2  ).innerHTML+=" difference "+(score[rider2]-score[rider])
                 document.getElementById("---"+rider2).innerHTML+=" difference "+(score[rider2]-score[rider])
          }

     }

     $('#console2 table').sort(function(a,b){
        return $(b).children('tbody').data('score') - $(a).children('tbody').data('score')
     }).each(function(i,v){
        $('#console2').append(v)
     })

      $('#console3 table').sort(function(a,b){
         return $(b).children('tbody').data('score') - $(a).children('tbody').data('score')
      }).each(function(i,v){
         $('#console3').append(v)
      })

};

print1 = function(rider,trick,score,score2) {

    var trickf = document.createElement('td');
    trickf.innerHTML=trick

    var scoref = document.createElement('td');
    scoref.innerHTML=score //+' <= ('+score2.toString()+')'

    var row = document.createElement('tr');
    row.appendChild(trickf);
    row.appendChild(scoref);

    document.getElementById(rider).appendChild(row);

};

print3 = function (id,rider,trick,score,attempt,judge){

   print5(rider)

   var riderf = document.createElement('td');
   riderf.innerHTML=rider

    var trickf = document.createElement('td');
    trickf.innerHTML=trick

    var scoref = document.createElement('td');
    scoref.innerHTML=score

    var attemptf = document.createElement('td');
    attemptf.innerHTML=attempt

    var judgef = document.createElement('td');
    judgef.innerHTML=judge

    var row = document.createElement('tr');
    row.id=id
    row.onclick=act
    row.appendChild(attemptf);
    row.appendChild(trickf);
    row.appendChild(scoref);
    row.appendChild(judgef);
    $('#--'+rider).prepend(row);

}

print4 = function(rider) {

    //if (document.getElementById('-'+rider)) return

    var caption = document.createElement('caption');
    caption.innerHTML='<h1>'+rider+'<span class="result" id="-'+rider+'"></span></h1>'

    var thead = document.createElement('thead');
    thead.innerHTML="<tr><th>Trick</th><th>Score</th></tr>"

    var tbody = document.createElement('tbody');
    tbody.id=rider
    //$(tbody).data('score',score)

    var table = document.createElement('table');
    table.className="table table-hover"
    table.appendChild(caption)
    table.appendChild(thead)
    table.appendChild(tbody)

    document.getElementById('console3').appendChild(table);

}

print5 = function(rider) {

    if (document.getElementById('--'+rider)) return

    var del = document.createElement('a');
    del.innerHTML='delete'
    del.onclick=service.delete

    var caption = document.createElement('caption');
    caption.innerHTML='<h1>'+rider+'<span class="result" id="---'+rider+'"></span></h1>'
    caption.appendChild(del);

    var thead = document.createElement('thead');
    thead.innerHTML="<tr><th>Attempt</th><th>Trick</th><th>Score</th><th>Judge</th></tr>"

    var tbody = document.createElement('tbody');
    tbody.id='--'+rider

    var table = document.createElement('table');
    table.className="table table-hover"
    table.appendChild(caption)
    table.appendChild(thead)
    table.appendChild(tbody)

    document.getElementById('console2').appendChild(table);
}

start = service.getFirst

stop = service.getFirst

view = function(x) {
    switch(x){
        case 1: $( "#form1" ).show(); $( "#form2" ).hide(); $( "#form3" ).hide();   break;
        case 2: $( "#form1" ).hide(); $( "#form2" ).show(); $( "#form3" ).hide();   break;
        default: $( "#form1" ).hide(); $( "#form2" ).hide(); $( "#form3" ).show();  break;
    }
}

testAPI = function () {
  console.log('Browser Facebook, Fetching your information... ');
  FB.api('/me?fields=email', function(response) {console.log('Browser Facebook, '+response.email+'.')})
}

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