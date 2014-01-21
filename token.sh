#!/bin/bash
client_id="...apps.googleusercontent.com"
client_secret="..."

scope="https://www.googleapis.com/auth/androidpublisher+https://www.googleapis.com/auth/devstorage.full_control+https://www.googleapis.com/auth/userinfo.email"
redirect="urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob"

echo "https://accounts.google.com/o/oauth2/auth?scope=$scope&response_type=code&access_type=offline&redirect_uri=$redirect&client_id=$client_id&approval_prompt=force"

read code

curl -X POST "https://accounts.google.com/o/oauth2/token"\
    -H "Content-type: application/x-www-form-urlencoded"\
    -d "grant_type=authorization_code&code=$code&redirect_uri=$redirect&client_id=$client_id&client_secret=$client_secret"
