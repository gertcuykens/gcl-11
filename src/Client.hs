import Control.Exception     ( bracket )
import Data.Acid             ( AcidState, createCheckpoint, closeAcidState )
import Data.Acid.Advanced    ( query', update' )
--import Data.Acid.Local       ( createArchive, openLocalState )
import Data.Acid.Remote      ( openRemoteState, sharedSecretPerform )
import Data.ByteString.Char8 ( pack, putStrLn)
import Data.Set              ( member, fromList )
import Data.Text             ( unpack )
import Keys                  ( serverKey )
import Network               ( PortID(PortNumber) )
import Google                ( Uid(..), uid, tid, url, f2)
import GroupMap              ( GroupMap(..), InsertKey(..), LookupKey(..) )
import Token                 ( create, verify )
import Prelude hiding        ( putStrLn )


openAcidState :: IO (AcidState GroupMap)
openAcidState = openRemoteState (sharedSecretPerform $ pack serverKey) "localhost" (PortNumber 8080)

runAcidState :: AcidState GroupMap -> IO ()
runAcidState acid = do
    putStrLn $ url [(pack "state", pack "[0,1,2,3]")]
    (code,state) <- fmap f2 getLine
    (Right t) <- tid code
    (Right (Uid u _ _)) <- uid t

    let i = read . unpack $ u
    c <- check i state acid
    create i (if c then state else []) >>= (\x -> print x >> verify x >>= print)

    runAcidState acid

main :: IO ()
main = bracket openAcidState closeAcidState runAcidState

check :: Integer -> [Int] -> AcidState GroupMap -> IO Bool
check _ [] _ = return True
check i [x] acid = query' acid (LookupKey x) >>= \(Just set) -> return (i `member` set)
check i (x:xs) acid = query' acid (LookupKey x) >>= \(Just set) -> if i `member` set then check i xs acid else return False

setup :: IO ()
setup = do
    acid <- openRemoteState (sharedSecretPerform $ pack serverKey) "localhost" (PortNumber 8080)
    _ <- update' acid (InsertKey 0 (fromList [116469479527388802962,555]))
    _ <- update' acid (InsertKey 1 (fromList [116469479527388802962]))
    _ <- update' acid (InsertKey 2 (fromList [116469479527388802962]))
    _ <- update' acid (InsertKey 3 (fromList [116469479527388802962]))
    createCheckpoint acid
    closeAcidState acid

