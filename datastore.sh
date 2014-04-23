#!/bin/bash
jwt1=`echo -n '{"alg":"RS256","typ":"JWT"}' | openssl base64 -e`

jwt2=`echo -n '{\
"iss":"'$DATASTORE_SERVICE_ACCOUNT'",\
"scope":"https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/plus.me https://www.googleapis.com/auth/datastore",\
"aud":"https://accounts.google.com/o/oauth2/token",\
"exp":'$(($(date +%s)+3600))',\
"iat":'$(date +%s)'}' | openssl base64 -e`

#"iss":"522156758812-u8hj8dhnk5br3vnpqqvuscievhbnl0gg@developer.gserviceaccount.com",\
#"prn":"gert.cuykens@gmail.com",\

jwt3=`echo -n "$jwt1.$jwt2" | tr -d '\n' | tr -d '=' | tr '/+' '_-'`

jwt4=`echo -n "$jwt3" | openssl sha -sha256 -sign <(openssl pkcs12 -passin pass:notasecret -in $DATASTORE_PRIVATE_KEY_FILE -nocerts -nodes) | openssl base64 -e`

jwt5=`echo -n "$jwt4" | tr -d '\n' | tr -d '=' | tr '/+' '_-'`

json=`curl -s -H "Content-type: application/x-www-form-urlencoded" -X POST "https://accounts.google.com/o/oauth2/token" -d "grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Ajwt-bearer&assertion=$jwt3.$jwt5"`

function jsonval {
    temp=`echo $json | sed 's/\\\\\//\//g' | sed 's/[{}]//g' | awk -v k="text" '{n=split($0,a,","); for (i=1; i<=n; i++) print a[i]}' | sed 's/\"\:\"/\|/g' | sed 's/[\,]/ /g' | sed 's/\"//g' | grep -w $prop | cut -d":" -f2 | sed -e 's/^ *//g' -e 's/ *$//g'`
    echo ${temp##*|}
}

prop='access_token'
export TOKEN=`jsonval`
echo $TOKEN

curl -s -X GET "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=$TOKEN"
curl -s -X GET "https://www.googleapis.com/oauth2/v2/userinfo?access_token=$TOKEN"
curl -s -H "Content-type: application/json" -H "Authorization:  Bearer $TOKEN" -X POST "https://www.googleapis.com/datastore/v1beta1/datasets/gcl-11/blindWrite" -d \
'{
 "mutation": {
  "upsert": [
   {
    "key": {
     "path": [
      {
       "kind": "person",
       "name": "gert"
      }
     ]
    }
   }
  ]
 }
}'
