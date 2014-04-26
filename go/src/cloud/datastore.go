package cloud

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"appengine/datastore"
)

type Entity struct {}

type DataStore struct {
	Root *datastore.Key
	Entity *Entity
	Context endpoints.Context
}
