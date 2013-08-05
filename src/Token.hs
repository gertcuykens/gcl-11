module Token where
import Codec.Utils           ( Octet, toTwosComp )
import Data.HMAC             ( hmac_sha1 )
import Data.UnixTime         ( getUnixTime, utSeconds )
import Keys                  ( secretKey )

data Token = Token {seconds::Int, userId::Integer, groupId::[Int], hash::[Octet]} deriving (Show)

groupOctet :: [Int] -> [Octet]
groupOctet [] = []
groupOctet [x] = toTwosComp x
groupOctet (x:xs) = groupOctet xs ++ toTwosComp x

verify :: Token -> IO Bool
verify t = do
    u <- getUnixTime
    let s = fromEnum (utSeconds u) - seconds t
    let h = toTwosComp (seconds t) ++ toTwosComp (userId t) ++ groupOctet (groupId t)
    return $ (s <= 3600) && (hmac_sha1 secretKey h == hash t)

create :: Integer -> [Int] -> IO Token
create u g = do
    t <- getUnixTime
    let s = fromEnum (utSeconds t)
    let h = toTwosComp s ++ toTwosComp u ++ groupOctet g
    return $ Token s u g (hmac_sha1 secretKey h)

