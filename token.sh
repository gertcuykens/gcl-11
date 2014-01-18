#!/bin/bash
client_id="1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com"
client_secret=$GCL_13_SECRET
#redirect_uri="http://localhost:8080/oauth2callback"
redirect_uri="http%3A%2F%2Flocalhost%3A8080%2Foauth2callback"
#force="&approval_prompt=force"
force=""

echo "https://accounts.google.com/o/oauth2/auth?scope=https://www.googleapis.com/auth/androidpublisher&response_type=code&access_type=offline&redirect_uri=$redirect_uri&client_id=$client_id$force"

read code

curl -X POST "https://accounts.google.com/o/oauth2/token"\
    -H "Content-type: application/x-www-form-urlencoded"\
    -d "grant_type=authorization_code&code=$code&client_id=$client_id&client_secret=$client_secret&redirect_uri=$redirect_uri"
