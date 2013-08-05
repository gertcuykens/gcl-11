module Token where
import Codec.Utils           ( Octet, toTwosComp )
import Data.HMAC             ( hmac_sha1 )
import Data.UnixTime         ( getUnixTime, utSeconds )
import Keys                  ( secretKey )

data Token = Token {seconds::Int, groupId::[Int], hash::[Octet]} deriving (Show)

groupOctet :: [Int] -> [Octet]
groupOctet [] = []
groupOctet [x] = toTwosComp x
groupOctet (x:xs) = groupOctet xs ++ toTwosComp x

verify :: Token -> IO Bool
verify t = do
    u <- getUnixTime
    let s = fromEnum (utSeconds u) - seconds t
    let h = toTwosComp (seconds t) ++ groupOctet (groupId t)
    return $ (s <= 3600) && (hmac_sha1 secretKey h == hash t)

create :: [Int] -> IO Token
create g = do
    u <- getUnixTime
    let s = fromEnum (utSeconds u)
    let h = toTwosComp s ++ groupOctet g
    return $ Token s g (hmac_sha1 secretKey h)

