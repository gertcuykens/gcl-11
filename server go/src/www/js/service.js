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
                                            "heat":parseInt($('#heat').val(), 10),
                                          }]}
  ).execute(
      function(resp) {
        if (!resp.code) {
            resp.list = resp.list || []
            document.getElementById('console2').innerHTML=""
            document.getElementById('console3').innerHTML=""
            if (resp.list[0]){heatf(resp.list[0])}
            mapreduce(resp.list);
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
 $('#heat2').html(s['event']+' Heat '+s['heat'])
 $('#heat3').html(s['event']+' Heat '+s['heat'])
 $('#event').attr('value',s['event'])
 $('#heat').attr('value',s['heat'])
}

heatb = function (v) {
 var h=parseInt($('#heat').attr('value'))
 $('#heat').attr('value',h+v)
 service.list()
}

mapreduce = function(s) {
  var object = {}
  for (var i = 0; i < s.length; i++) {
    var id=s[i]['id']
    var rider=s[i]['rider']
    var event=s[i]['event']
    var heat=s[i]['heat']
    var trick=s[i]['trick']
    var attempt=s[i]['attempt']
    var score=s[i]['score']
    var judge=s[i]['judge']
    if (!object[rider]) {object[rider]={};}
    if (!object[rider][event]) {object[rider][event]={}}
    if (!object[rider][event][heat]) {object[rider][event][heat]=[];}
    if (!object[rider][event][heat][attempt]) {object[rider][event][heat][attempt]=[]}
    if (!object[rider][event][heat][attempt][judge]) {object[rider][event][heat][attempt][judge]=[score,trick,id]}
  }
  //console.log(object)

   var trick={}
   for (rider in object) {
    if(object.hasOwnProperty(rider)){

     //console.log(rider)
     if (!trick[rider]) {trick[rider]={};}
     for (event in object[rider]) {
      if(object[rider].hasOwnProperty(event)){

       //console.log(event)
       if (!trick[rider][event]) {trick[rider][event]={}}
       for (heat in object[rider][event]) {
         //console.log(heat)

         var list={}

         if (!trick[rider][event][heat]) {trick[rider][event][heat]=[]}
         for (attempt in object[rider][event][heat]) {
           //console.log(attempt)

           var name=""
           var max=0
           var score=0
           var score2=[]
           for (judge in object[rider][event][heat][attempt]) {
            if(object[rider][event][heat][attempt].hasOwnProperty(judge)){
             max+=10
             scorej=object[rider][event][heat][attempt][judge][0]
             score+=object[rider][event][heat][attempt][judge][0]
             score2.push(object[rider][event][heat][attempt][judge][0])
             name=object[rider][event][heat][attempt][judge][1]
             id=object[rider][event][heat][attempt][judge][2]
             print3(id,rider,name,scorej,attempt,judge)
             if (!list[name]){list[name]=[score,max,name,score2]}
             else if (list[name][0]<score){list[name]=[score,max,name,score2]}
            }
           }

           //console.log(name+':'+score+'/'+max)

         }

         var keys = Object.keys(list).sort(function(a,b){return list[b][0]-list[a][0]})
         for (k in keys) {trick[rider][event][heat].push(list[keys[k]])}

       }
      }
     }
    }
   }
   //console.log(trick)

     var score={}
     for (var i = 0; i < s.length; i++) {
       var rider=s[i]['rider']
       var event=s[i]['event']
       var heat=s[i]['heat']
       if (!score[event]) {score[event]={}}
       if (!score[event][heat]) {score[event][heat]=[]}
       if (!score[event][heat][rider]) {
        var count=0
        var total=0
        for (t in trick[rider][event][heat]) {total+=trick[rider][event][heat][t][0]}
        count++
        if (count>5){break}
        score[event][heat][rider]=total
       }
     }
     //console.log(score)

     var rank={}
     for (event in score) {
       if(score.hasOwnProperty(event)){
         if (!rank[event]) {rank[event]={}}
         for (heat in score[event]) {
            if(score[event].hasOwnProperty(heat)){
                  if (!rank[event][heat]) {rank[event][heat]=[]}
                  var list = score[event][heat]
                  rank[event][heat]=Object.keys(list).sort(function(a,b){return list[b]-list[a]})
            }
         }
       }
     }
     //console.log(rank)

     for (event in rank) {
       if(rank.hasOwnProperty(event)){
         //console.log(event)
         for (heat in rank[event]) {
            if(rank[event].hasOwnProperty(heat)){
              //console.log(heat)
              for (place in rank[event][heat]) {
               if(rank[event][heat].hasOwnProperty(place)){
                 //console.log(rank[event][heat][place])
                 var rider= rank[event][heat][place]
                 var rider2= rider
                 if (place=='0' && rank[event][heat][parseInt(place)+1]) {rider2= rank[event][heat][parseInt(place)+1]}
                 else {rider2= rank[event][heat][0]}

                 print4(rider, score[event][heat][rider])
                 //TODO:clean
                 document.getElementById("-"+rider  ).innerHTML=" score "+score[event][heat][rider]
                 $('#--'+rider).data('score',score[event][heat][rider])
                 document.getElementById("---"+rider).innerHTML=" score "+score[event][heat][rider]
                 document.getElementById("-"+rider  ).innerHTML+=" difference "+(score[event][heat][rider]-score[event][heat][rider2])
                 document.getElementById("---"+rider).innerHTML+=" difference "+(score[event][heat][rider]-score[event][heat][rider2])

               }
              }
            }
         }
       }
     }

     $('#console2 table').sort(function(a,b){
        return $(b).children('tbody').data('score') - $(a).children('tbody').data('score')
     }).each(function(i,v){
      $('#console2').append(v)
     })

    for (rider in trick){
     if(trick.hasOwnProperty(rider)){
         for (event in trick[rider]){
          if(trick[rider].hasOwnProperty(event)){
            for (heat in trick[rider][event]){

               for (score in trick[rider][event][heat]){
                   print1(rider,trick[rider][event][heat][score][2],trick[rider][event][heat][score][0],trick[rider][event][heat][score][3])
               }

            }
          }
        }
      }
    }

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

print4 = function(rider,score) {

    //if (document.getElementById('-'+rider)) return

    var caption = document.createElement('caption');
    caption.innerHTML='<h1>'+rider+'<span class="result" id="-'+rider+'"></span></h1>'

    var thead = document.createElement('thead');
    thead.innerHTML="<tr><th>Trick</th><th>Score</th></tr>"

    var tbody = document.createElement('tbody');
    tbody.id=rider
    $(tbody).data('score',score)

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

start = service.list

stop = service.list

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