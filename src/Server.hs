import Control.Exception     ( bracket )
import Data.Acid             ( AcidState, closeAcidState )
import Data.Acid.Local       ( openLocalState )
import Data.Acid.Remote      ( acidServer, sharedSecretCheck )
import Data.ByteString.Char8 ( pack )
import Data.IntMap           ( empty )
import Data.Set              ( singleton )
import Keys                  ( serverKey )
import Network               ( PortID(PortNumber) )
import UserMap               ( UserMap(..) )
import GroupMap              ( GroupMap(..) )

openAcidStateU :: IO (AcidState UserMap)
openAcidStateU = openLocalState $ UserMap empty

openAcidStateG :: IO (AcidState GroupMap)
openAcidStateG = openLocalState $ GroupMap empty

runAcidStateU :: AcidState UserMap -> IO ()
runAcidStateU = acidServer (sharedSecretCheck (singleton $ pack serverKey)) (PortNumber 8081)

runAcidStateG :: AcidState GroupMap -> IO ()
runAcidStateG = acidServer (sharedSecretCheck (singleton $ pack serverKey)) (PortNumber 8082)

main :: IO ()
main = do
    bracket openAcidStateU closeAcidState runAcidStateU
    bracket openAcidStateG closeAcidState runAcidStateG

