#!/bin/bash

curl -X GET https://graph.facebook.com/gert.cuykens?fields=id,name \
     -H "Authorization:  Bearer $1" \
     -w "\n"

curl -X GET https://graph.facebook.com/gert.cuykens/permissions \
     -H "Authorization:  Bearer $1" \
     -w "\n"

curl -X GET https://graph.facebook.com/gert.cuykens/accounts \
     -H "Authorization:  Bearer $1" \
     -w "\n"
