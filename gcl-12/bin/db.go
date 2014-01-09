package bin

import (
	"appengine/datastore"
)

type Data struct {
	Key *datastore.Key `datastore:"-"`
	Data string `datastore:"data"`
    Token *Token `datastore:"-"`
}

func (d *Data) Store() error{
	if d.Key == nil {d.Token.Status="Datastore no Key!"; return d.Token}
	key, err := datastore.Put(d.Token.Context, d.Key, d);
	if err != nil {d.Token.Status="Datastore put error! "+err.Error(); return d.Token}
	d.Token.Status="Stored "+key.StringID()+"."
	return nil
}

func (d *Data) Get() (err error){
	if d.Key == nil {d.Token.Status="Datastore no Key!"; return d.Token}
	if err = datastore.Get(d.Token.Context, d.Key, d.Token); err != nil {d.Token.Status="Datastore get error! "+err.Error(); return d.Token}
	d.Token.Status="Fetched "+d.Key.StringID()+"."
	return nil
}

func (d *Data) Query() (err error){
	return nil
}
