{-# LANGUAGE OverloadedStrings, TemplateHaskell #-}
module Google where
import Data.Aeson (FromJSON)
import Data.Aeson.TH (deriveJSON)
import qualified Data.ByteString.Char8 as BS
import Data.Text (Text, unpack)
import Keys (googleKey)
import Network.OAuth.OAuth2
import Prelude hiding (id)
import qualified Prelude as P (id)
import Token (create, verify)

data Email = Email { id             :: Text
                   , email          :: Text
                   , verified_email :: Bool
                   } deriving (Show)

$(deriveJSON P.id ''Email)

googleScopeEmail :: QueryParams
googleScopeEmail = [("scope", "https://www.googleapis.com/auth/userinfo.email")]

googleAccessOffline :: QueryParams
googleAccessOffline = [("access_type", "offline")
                      ,("approval_prompt", "force")]

userinfo :: FromJSON a => AccessToken -> IO (OAuth2Result a)
userinfo token = authGetJSON token "https://www.googleapis.com/oauth2/v2/userinfo"

test :: IO ()
test = do
    BS.putStrLn $ authorizationUrl googleKey `appendQueryParam` (googleScopeEmail ++ googleAccessOffline)
    code <- fmap BS.pack getLine
    (Right token) <- fetchAccessToken googleKey code
    f token
    case refreshToken token of
        Nothing -> putStrLn "Failed to fetch refresh token"
        Just rt -> do
            (Right nt) <- fetchRefreshToken googleKey rt
            f nt
    where f token = (userinfo token :: IO (OAuth2Result Email)) >>= \(Right x) -> print x >> create (read $ unpack $ id x) [] >>= \y -> print y >> verify y >>= print

-- googleScopeProfile :: QueryParams
-- googleScopeProfile = [("scope", "https://www.googleapis.com/auth/userinfo.profile")]

-- validateToken :: FromJSON a => AccessToken -> IO (OAuth2Result a)
-- validateToken token = authGetJSON token "https://www.googleapis.com/oauth2/v1/tokeninfo"

-- data Token = Token { issued_to      :: Text
--                    , audience       :: Text
--                    , user_id        :: Text
--                    , scope          :: Text
--                    , expires_in     :: Integer
--                    , email          :: Text
--                    , verified_email :: Bool
--                    , access_type    :: Text
--                    } deriving (Show)
--
-- $(deriveJSON P.id ''Token)

-- data UserInfo = Profile { id          :: Text
--                         , name        :: Text
--                         , given_name  :: Text
--                         , family_name :: Text
--                         , link        :: Text
--                         , picture     :: Text
--                         , gender      :: Text
--                         , birthday    :: Text
--                         , locale      :: Text
--                         } deriving (Show)
--
-- $(deriveJSON P.id ''Profile)

