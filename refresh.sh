#!/bin/bash
client_id="1093123836049-ilqfjb2s2tqal2fobuduj8b790hnnfju.apps.googleusercontent.com"
client_secret=$GCL_13_SECRET
refresh_token=$GCL_13_REFRESH

curl -X POST "https://accounts.google.com/o/oauth2/token"\
    -H "Content-type: application/x-www-form-urlencoded"\
    -d "grant_type=refresh_token&client_id=$client_id&client_secret=$client_secret&refresh_token=$refresh_token"
