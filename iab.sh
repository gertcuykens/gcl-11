#!/bin/bash
iab=""

curl -X GET "https://www.googleapis.com/androidpublisher/v1.1/applications/com.appspot/inapp/gas/purchases/$iab"\
    -H "Authorization:  Bearer $1"

#curl -s -X GET "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=$1"
#curl -s -X GET "https://www.googleapis.com/oauth2/v2/userinfo?access_token=$1"
#curl -s -v -X POST "https://gcl-11.appspot.com/_ah/api/rest1/0/greetings/authed" \
#        -H "Authorization:  Bearer $1" \
#        -H "Content-Type: application/json; charset=UTF-8" \
#        -d ""
