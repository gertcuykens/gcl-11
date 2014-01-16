#!/bin/bash
client_id="1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com"
redirect_uri="https%3A%2F%2Fgcl-13.appspot.com%2Foauth2callback"

url="https://accounts.google.com/o/oauth2/auth?scope=https://www.googleapis.com/auth/androidpublisher&response_type=code&access_type=offline&redirect_uri=$redirect_uri&client_id=$client_id"
echo $url

read code

curl -X POST "https://accounts.google.com/o/oauth2/token"\
    -H "Content-type: application/x-www-form-urlencoded"\
    -d "grant_type=authorization_code&code=$code&client_id=$client_id&client_secret=$GCL_13_SECRET&redirect_uri=$redirect_uri"

curl -X POST "https://accounts.google.com/o/oauth2/token"\
    -H "Content-type: application/x-www-form-urlencoded"\
    -d "grant_type=refresh_token&client_id=$client_id&client_secret=$GCL_13_SECRET&refresh_token=$GCL_13_REFRESH"

#curl -s -X GET "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=$1"
#curl -s -X GET "https://www.googleapis.com/oauth2/v2/userinfo?access_token=$1"
#curl -s -v -X POST "https://gcl-11.appspot.com/_ah/api/rest1/0/greetings/authed" \
#        -H "Authorization:  Bearer $1" \
#        -H "Content-Type: application/json; charset=UTF-8" \
#        -d ""

