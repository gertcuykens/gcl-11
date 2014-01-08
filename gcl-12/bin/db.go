package bin

import (
	"appengine/datastore"
	"appengine"
)

type Data struct {
	Key *datastore.Key `datastore:"-"`
	Data string `datastore:"data"`
	Context appengine.Context `datastore:"-"`
	Status string `datastore:"-"`
}

func (d *Data) Error() string {
	return d.Status
}

func (d *Data) Store() error{
	if d.Key == nil {d.Status="Datastore no Key!"; return d}
	key, err := datastore.Put(d.Context, d.Key, d);
	if err != nil {d.Status="Datastore put error! "+err.Error(); return d}
	d.Status="Stored "+key.StringID()+"."
	return nil
}

func (d *Data) Get() (err error){
	if d.Key == nil {d.Status="Datastore no Key!"; return d}
	if err = datastore.Get(d.Context, d.Key, d); err != nil {d.Status="Datastore get error! "+err.Error(); return d}
	d.Status="Fetched "+d.Key.StringID()+"."
	return nil
}

func (d *Data) Query() (err error){
	return nil
}

