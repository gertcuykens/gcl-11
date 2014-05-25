#!/bin/bash
client_id="249348058430770"
client_secret=""

curl -X GET "https://graph.facebook.com/oauth/access_token?grant_type=fb_exchange_token&client_id=$client_id&client_secret=$client_secret&fb_exchange_token=$1"
