#!/bin/bash
client_id="249348058430770"

curl -X GET "https://graph.facebook.com/oauth/access_token?grant_type=authorization_code&client_id=$client_id&redirect_uri=localhost"