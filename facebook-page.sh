#!/bin/bash
access_token=""

curl -X POST https://graph.facebook.com/gcl11/feed \
     -H "Authorization:  Bearer $access_token" \
     -d "message=This+is+a+test+message+from+app+gcl-11" \
     -w "\n"
