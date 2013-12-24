#!/bin/sh
URL='https://gcl-11.appspot.com/_ah/api/discovery/v1/apis/rest1/0/rest'
curl -s $URL > greetings.rest.discovery
endpointscfg.py gen_client_lib java greetings.rest.discovery
