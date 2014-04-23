#!/bin/bash

FNAME=$1

if [ "${FNAME##.*}" == "des" ]; then
    openssl des3 -d -salt -in "$FNAME" -out "${FNAME%.*}"
else
    openssl des3 -salt -in "$FNAME" -out "$FNAME.des"
fi
