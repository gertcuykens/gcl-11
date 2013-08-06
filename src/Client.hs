import Control.Exception     ( bracket )
import Data.Acid             ( AcidState, createCheckpoint, closeAcidState )
import Data.Acid.Advanced    ( query', update' )
--import Data.Acid.Local       ( createArchive, openLocalState )
import Data.Acid.Remote      ( openRemoteState, sharedSecretPerform )
import Data.ByteString.Char8 ( pack )
import Data.Set              ( member, fromList )
import Keys                  ( serverKey )
import Network               ( PortID(PortNumber) )
import GroupMap              ( GroupMap(..), InsertKey(..), LookupKey(..) )
import Token                 ( create, verify )

openAcidState :: IO (AcidState GroupMap)
openAcidState = openRemoteState (sharedSecretPerform $ pack serverKey) "localhost" (PortNumber 8080)

runAcidState :: AcidState GroupMap -> IO ()
runAcidState acid = do

    _ <- update' acid (InsertKey 0 (fromList [116469479527388802962,555]))
    _ <- update' acid (InsertKey 1 (fromList [116469479527388802962]))
    _ <- update' acid (InsertKey 2 (fromList [116469479527388802962]))
    _ <- update' acid (InsertKey 3 (fromList [116469479527388802962]))

    c <- check 116469479527388802962 [0,1,2,3] acid
    print c

    t <- create 116469479527388802962 [0,1,2,3]
    print t
    b <- verify t
    print b

    createCheckpoint acid

main :: IO ()
main = bracket openAcidState closeAcidState runAcidState

check :: Integer -> [Int] -> AcidState GroupMap -> IO Bool
check _ [] _ = return True
check uid [x] acid = query' acid (LookupKey x) >>= \(Just set) -> return (uid `member` set)
check uid (x:xs) acid = query' acid (LookupKey x) >>= \(Just set) -> return (uid `member` set) >>= \r -> if (r == True) then (check uid xs acid) else (return False)

