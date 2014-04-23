#!/bin/sh
URL='https://gcl-13.appspot.com/_ah/api/discovery/v1/apis/rest/v0/rest'
curl -s $URL > discovery.zip
endpointscfg.py gen_client_lib java discovery.zip
