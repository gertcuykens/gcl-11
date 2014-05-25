#!/bin/sh
REST='https://gcl-11.appspot.com/_ah/api/discovery/v1/apis/service/v0/rest'
curl -s $REST > rest.zip
endpointscfg.py gen_client_lib java rest.zip
