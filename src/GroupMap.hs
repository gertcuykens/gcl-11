{-# LANGUAGE DeriveDataTypeable, TypeFamilies, TemplateHaskell#-}
module GroupMap where
import Control.Lens ((?=), at, from, makeIso, view)
import Data.Acid (Update, Query, makeAcidic)
import Data.SafeCopy (deriveSafeCopy, base)
import Data.Typeable (Typeable)
import Data.IntMap (IntMap)
import Data.IntSet (IntSet)

newtype GroupMap = GroupMap (IntMap IntSet) deriving (Show, Typeable)

$(deriveSafeCopy 0 'base ''GroupMap)

$(makeIso ''GroupMap)

insertKey :: Int -> IntSet -> Update GroupMap ()
insertKey k v = (from groupMap.at k) ?= v

lookupKey :: Int -> Query GroupMap (Maybe IntSet)
lookupKey k = view (from groupMap.at k)

$(makeAcidic ''GroupMap ['insertKey, 'lookupKey])

