import Control.Exception     ( bracket )
import Data.Acid             ( AcidState, createCheckpoint, closeAcidState )
import Data.Acid.Advanced    ( query', update' )
--import Data.Acid.Local       ( createArchive, openLocalState )
import Data.Acid.Remote      ( openRemoteState, sharedSecretPerform )
import Data.ByteString.Char8 ( pack )
import Keys                  ( serverKey )
import Network               ( PortID(PortNumber) )
import Table                 ( GroupMap(..), Group(..), InsertKey(..), LookupKey(..), group)
import Token                 ( create, verify )

openAcidState :: IO (AcidState GroupMap)
openAcidState = openRemoteState (sharedSecretPerform $ pack serverKey) "localhost" (PortNumber 8080)

runAcidState :: AcidState GroupMap -> IO ()
runAcidState acid = do
    _ <- update' acid (InsertKey 0 (Group [116469479527388802962,555]))
    _ <- update' acid (InsertKey 1 (Group [116469479527388802962]))
    _ <- update' acid (InsertKey 2 (Group [116469479527388802962]))
    _ <- update' acid (InsertKey 3 (Group [116469479527388802962]))

    Just p <- query' acid (LookupKey 0)
    print (116469479527388802962 `elem` group p)
    print (555 `elem` group p)
    print (666 `elem` group p)

    t <- create 116469479527388802962 [0,1,2,3]
    print t
    b <- verify t
    print b

    createCheckpoint acid



main :: IO ()
main = bracket openAcidState closeAcidState runAcidState

