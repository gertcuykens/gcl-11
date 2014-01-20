#!/bin/bash
id=""
secret=""
#scope="https://www.googleapis.com/auth/androidpublisher"
scope="https://www.googleapis.com/auth/devstorage.full_control"

echo "https://accounts.google.com/o/oauth2/auth?scope=$scope&response_type=code&access_type=offline&redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob&client_id=$id&approval_prompt=force"

read code

curl -X POST "https://accounts.google.com/o/oauth2/token"\
    -H "Content-type: application/x-www-form-urlencoded"\
    -d "grant_type=authorization_code&code=$code&redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob&client_id=$id&client_secret=$secret"
