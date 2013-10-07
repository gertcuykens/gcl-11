curl -X GET "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=$1"
curl -X GET "https://www.googleapis.com/oauth2/v2/userinfo?access_token=$1"
curl -H "Content-type: application/json" -H "Authorization:  Bearer $1" -X POST "https://www.googleapis.com/datastore/v1beta1/datasets/gcl-11/blindWrite" -d \
'{
 "mutation": {
  "upsert": [
   {
    "key": {
     "path": [
      {
       "kind": "person",
       "name": "gert"
      }
     ]
    }
   }
  ]
 }
}'

