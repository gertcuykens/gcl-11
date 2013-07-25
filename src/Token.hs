module Token where
import Codec.Utils           ( Octet, toTwosComp )
import Data.HMAC             ( hmac_sha1 )
import Data.UnixTime         ( getUnixTime, utSeconds )

data Token = Token {seconds::Int, userId::Int, groupId::Int, hash::[Octet]} deriving (Show)

secret :: [Octet]
secret = [48,57]

verify :: Token -> IO Bool
verify t = do
    u <- getUnixTime
    let s = fromEnum (utSeconds u) - seconds t
    let h = toTwosComp (seconds t) ++ toTwosComp (userId t) ++ toTwosComp (groupId t)
    return $ (s <= 3600) && (hmac_sha1 secret h == hash t)

create :: Int -> Int -> IO Token
create i1 i2 = do
    u <- getUnixTime
    let s = fromEnum (utSeconds u)
    let h = toTwosComp s ++ toTwosComp i1 ++ toTwosComp i2
    return $ Token s i1 i2 (hmac_sha1 secret h)

