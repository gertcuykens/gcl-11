service = {}

service.rider = function(rider) {
  var option = document.createElement('option');
  option.innerHTML=rider
  option.id = 'option-'+rider
  $(option).data('attempt','1')
  $('#rider').append(option)
};

$('#rider').on('change',function (){var i = this.selectedIndex; $('#attempt').val($(this.options[i]).data('attempt'))})

$('#riderAdd').on('click',function (){service.rider(prompt("add rider"))})

$('#scoreAdd').on('click',function (){service.submit()})

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
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.editor().execute(function(response){
    if (!response.error){
      $('#loginButton').hide()
    } else {
      $('#loginButton').show()
    }
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
  var t= new Tokeng()
  t.access_token = FB.getAccessToken()
  gapi.auth.setToken(t)
  gapi.client.service.datastore.put({"list":[{
    "judge":"admin",
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
            $('#rider').html('')
            mapreduce(resp.list);
            $('#rider').trigger('change')
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
            if (resp.list[0]){heatf(resp.list[0]); service.list()}
            else {view(2)}
        }
      }
  );
};

act = function () {
 if (this.className=="active") {this.className="";return;}
 for (var r in $(this).parent('tbody').children('tr')){$(this).parent('tbody').children('tr')[r].className=""}
 this.className="active"
}

heatf = function (s) {
 $('#heat2').html(s.event+' '+s.division+' heat '+s.heat)
 $('#heat3').html(s.event+' '+s.division+' heat '+s.heat)
 $('#event').val(s.event)
 $('#division').val(s.division)
 $('#heat').val(s.heat)
}

heatb = function (v) {
 var e=$('#event').val()
 var d=$('#division').val()
 var h=parseInt($('#heat').val())
 heatf({'event':e,'division':d,'heat':h+v})
 service.list()
}

division = function () {
 var e=$('#event').val()
 var h=parseInt($('#heat').val())
 var d=$('#division').val()
 switch (d) {
  case "men": $('#division').val('woman'); break;
  default: $('#division').val('men'); break;
 }
 d=$('#division').val()
 heatf({'event':e,'division':d,'heat':h})
 service.list()
}

mapreduce = function(s) {
  var object = {}
  for (var i = 0; i < s.length; i++) {
    var id=s[i].id
    var rider=s[i].rider
    if (!rider) rider="Unknown"
    var trick=s[i].trick
    var attempt=s[i].attempt
    var score=s[i].score
    var judge=s[i].judge
    if (!object[rider]) {object[rider]=[];print4(rider);}
    if (!object[rider][attempt]) {object[rider][attempt]={}}
    if (!object[rider][attempt][judge]) {object[rider][attempt][judge]=[score,trick,id]}
  }
  //console.log(object)

   var trick={}
   for (rider in object) {
    if(object.hasOwnProperty(rider)){

         var list={}

         if (!trick[rider]) {trick[rider]=[]}
         for (attempt in object[rider]) {
           //console.log(attempt)

           var name=""
           var score=0
           var score2=[]
           for (judge in object[rider][attempt]) {
            if(object[rider][attempt].hasOwnProperty(judge)){
             scorej=object[rider][attempt][judge][0]
             score+=object[rider][attempt][judge][0]
             score2.push(object[rider][attempt][judge][0])
             name=object[rider][attempt][judge][1]
             id=object[rider][attempt][judge][2]
             print3(id,rider,name,scorej,attempt,judge)
             if (!list[name]){list[name]=[score,score2,name]}
             else if (list[name][0]<score){list[name]=[score,score2,name]}
            }
           }

           //console.log(name+':'+score+'/'+max)

         }

         var keys = Object.keys(list).sort(function(a,b){return list[b][0]-list[a][0]})
         for (var k in keys) {
         trick[rider].push(list[keys[k]])


         }

     }
    }
    //console.log(trick)

    for (rider in trick){
     if(trick.hasOwnProperty(rider)){


               for (var score in trick[rider]){

  print1(rider,trick[rider][score][0],trick[rider][score][1],trick[rider][score][2])

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
     //console.log(score)

     var keys=Object.keys(score).sort(function(a,b){return score[b]-score[a]})

     var rider
     if (keys[0]) rider=keys[0]
     for (var k in keys) {

         var rider2= keys[k]

          document.getElementById("-"+rider2  ).innerHTML=" score "+score[rider2]
          $('#--'+rider2).data('score',score[rider2])
          $('#'+rider2).data('score',score[rider2])
          document.getElementById("---"+rider2).innerHTML=" score "+score[rider2]

          if (k=='0' && keys[parseInt(k)+1]) {
                 rider2= keys[parseInt(k)+1]
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

print1 = function(rider,score,score2,trick) {

    var trickf = document.createElement('td');
    trickf.innerHTML=trick

    var scoref = document.createElement('td');
    scoref.innerHTML=score +'/'+ score2.length*10

    var row = document.createElement('tr');
    row.appendChild(trickf);
    row.appendChild(scoref);

    document.getElementById(rider).appendChild(row);

};

print3 = function (id,rider,trick,score,attempt,judge){

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

    $('#option-'+rider).data('attempt',parseInt(attempt)+1)

}

print4 = function(rider) {

    if (document.getElementById(rider)) return

    var add = document.createElement('a');
    add.innerHTML='add'
    add.onclick=function(){view(2)}

    var caption = document.createElement('caption');
    caption.innerHTML='<h1>'+rider+'<span class="result" id="-'+rider+'"></span></h1>'
    caption.appendChild(add);

    var thead = document.createElement('thead');
    thead.innerHTML="<tr><th>Trick</th><th>Score</th></tr>"

    var tbody = document.createElement('tbody');
    tbody.id=rider

    var table = document.createElement('table');
    table.className="table table-hover"
    table.appendChild(caption)
    table.appendChild(thead)
    table.appendChild(tbody)

    document.getElementById('console3').appendChild(table);

    service.rider(rider)

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

start = function(){service.editor(); service.getTrickList(); service.getFirst();}

stop = service.getFirst

view = function(x) {
    switch(x){
        case 1: $( "#form1" ).show(); $( "#form2" ).hide(); $( "#form3" ).hide(); break;
        case 2: $( "#form1" ).hide(); $( "#form2" ).show(); $( "#form3" ).hide(); service.editor(); break;
        default: $( "#form1" ).hide(); $( "#form2" ).hide(); $( "#form3" ).show(); break;
    }
}
