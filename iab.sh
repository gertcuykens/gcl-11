#!/bin/bash
client_id="1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com"
client_secret=$GCL_13_SECRET
iab_token=$GCL_13_IAB

curl -X GET "https://www.googleapis.com/androidpublisher/v1.1/applications/com.appspot/inapp/gas/purchases/$iab_token"\
    -H "Authorization:  Bearer $1"

#curl -s -X GET "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=$1"

#curl -s -X GET "https://www.googleapis.com/oauth2/v2/userinfo?access_token=$1"

#curl -s -v -X POST "https://gcl-11.appspot.com/_ah/api/rest1/0/greetings/authed" \
#        -H "Authorization:  Bearer $1" \
#        -H "Content-Type: application/json; charset=UTF-8" \
#        -d ""
