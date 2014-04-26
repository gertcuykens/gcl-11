#!/bin/sh
URL='https://gcl-11.appspot.com/_ah/api/discovery/v1/apis/rest1/0/rest'
curl -s $URL > discovery.zip
endpointscfg.py gen_client_lib java discovery.zip
