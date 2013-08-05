{-# LANGUAGE DeriveDataTypeable, TypeFamilies, TemplateHaskell#-}
module GroupMap where
import Control.Lens ((?=), at, from, makeIso, view)
import Data.Acid (Update, Query, makeAcidic)
import Data.SafeCopy (deriveSafeCopy, base)
import Data.Text (Text)
import Data.Typeable (Typeable)
import Data.IntMap (IntMap)
import Data.Set (Set)

newtype GroupMap = GroupMap (IntMap (Set Integer)) deriving (Show, Typeable)

$(deriveSafeCopy 0 'base ''GroupMap)

$(makeIso ''GroupMap)

insertKey :: Int -> Set Text -> Update GroupMap ()
insertKey k v = (from groupMap.at k) ?= v

lookupKey :: Int -> Query GroupMap (Maybe (Set Text))
lookupKey k = view (from groupMap.at k)

$(makeAcidic ''GroupMap ['insertKey, 'lookupKey])

