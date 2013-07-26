{-# LANGUAGE OverloadedStrings #-}
module Keys where
import Codec.Utils (Octet)
import Network.OAuth.OAuth2 (OAuth2(..))

serverKey :: String
serverKey = "12345"

secretKey :: [Octet]
secretKey = [48,57]

googleKey :: OAuth2
googleKey = OAuth2 { oauthClientId = "522156758812-09f5qv0e4gqjdjqfocerqcud5m5jutau.apps.googleusercontent.com"
                   , oauthClientSecret = ""
                   , oauthCallback = Just "http://localhost:9160"
                   , oauthOAuthorizeEndpoint = "https://accounts.google.com/o/oauth2/auth"
                   , oauthAccessTokenEndpoint = "https://accounts.google.com/o/oauth2/token"}

facebookKey :: OAuth2
facebookKey = OAuth2 { oauthClientId = ""
                     , oauthClientSecret = ""
                     , oauthCallback = Just "http://localhost:9160"
                     , oauthOAuthorizeEndpoint = "https://www.facebook.com/dialog/oauth"
                     , oauthAccessTokenEndpoint = "https://graph.facebook.com/oauth/access_token"}

-- https://developers.google.com/oauthplayground

