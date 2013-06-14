jwt1=`echo -n '{"alg":"RS256","typ":"JWT"}' | openssl base64 -e`

jwt2=`echo -n '{\
"iss":"522156758812-u8hj8dhnk5br3vnpqqvuscievhbnl0gg@developer.gserviceaccount.com",\
"scope":"https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/plus.me https://www.googleapis.com/auth/datastore",\
"aud":"https://accounts.google.com/o/oauth2/token",\
"exp":'$(($(date +%s)+3600))',\
"iat":'$(date +%s)'}' | openssl base64 -e`

jwt3=`echo -n "$jwt1.$jwt2" | tr -d '\n' | tr -d '=' | tr '/+' '_-'`

jwt4=`echo -n "$jwt3" | openssl sha -sha256 -sign google.p12 | openssl base64 -e`

jwt5=`echo -n "$jwt4" | tr -d '\n' | tr -d '=' | tr '/+' '_-'`

curl -H "Content-type: application/x-www-form-urlencoded" -X POST "https://accounts.google.com/o/oauth2/token" -d \
"grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Ajwt-bearer&assertion=$jwt3.$jwt5"

#"prn":"gert.cuykens@gmail.com",\
