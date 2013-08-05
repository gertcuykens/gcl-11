import Control.Exception     ( bracket )
import Data.Acid             ( AcidState, createCheckpoint, closeAcidState )
import Data.Acid.Advanced    ( query', update' )
--import Data.Acid.Local       ( createArchive, openLocalState )
import Data.Acid.Remote      ( openRemoteState, sharedSecretPerform )
import Data.ByteString.Char8 ( pack )
import Data.Set              ( member, fromList )
import Data.Text             ( Text )
import Keys                  ( serverKey )
import Network               ( PortID(PortNumber) )
import GroupMap              ( GroupMap(..), InsertKey(..), LookupKey(..) )
import Token                 ( create, verify )

openAcidState :: IO (AcidState GroupMap)
openAcidState = openRemoteState (sharedSecretPerform $ pack serverKey) "localhost" (PortNumber 8080)

runAcidState :: AcidState GroupMap -> IO ()
runAcidState acid = do

    _ <- update' acid (InsertKey 0 (fromList ["116469479527388802962","555"]))
    _ <- update' acid (InsertKey 1 (fromList ["116469479527388802962"]))
    _ <- update' acid (InsertKey 2 (fromList ["116469479527388802962"]))
    _ <- update' acid (InsertKey 3 (fromList ["116469479527388802962"]))

    Just s <- query' acid (LookupKey 0)
    print ("116469479527388802962" `member` s)
    print ("555" `member` s)
    print ("666" `member` s)

    t <- create "116469479527388802962" [0,1,2,3]
    print t
    b <- verify t
    print b

    createCheckpoint acid

main :: IO ()
main = bracket openAcidState closeAcidState runAcidState

