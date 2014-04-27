#!/bin/bash
client_id=""
client_secret=""
refresh_token=""

curl -X POST "https://accounts.google.com/o/oauth2/token"\
    -H "Content-type: application/x-www-form-urlencoded"\
    -d "grant_type=refresh_token&client_id=$client_id&client_secret=$client_secret&refresh_token=$refresh_token"
