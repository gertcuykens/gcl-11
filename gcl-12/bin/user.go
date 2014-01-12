package bin

import "appengine/datastore"

type User struct {
	Key *datastore.Key `datastore:"-"`
	Group []Property `datastore:"group"`
	Token *Token `datastore:"-"`
}

func (u *User) Error() string {
	return u.Token.Status
}

func (u *User) Put() error{
	if u.Key == nil {u.Token.Status="Datastore no Key!"; return u}
	key, err := datastore.Put(u.Token.Context, u.Key, u);
	if err != nil {u.Token.Status="Datastore put error! "+err.Error(); return u}
	u.Token.Status="Stored "+key.String()+"."
	return nil
}

func (u *User) Get() (err error){
	if u.Key == nil {u.Token.Status="Datastore no Key!"; return u}
	if err = datastore.Get(u.Token.Context, u.Key, u); err != nil {u.Token.Status="Datastore get error! "+err.Error(); return u}
	u.Token.Status="Fetched "+u.Key.String()+"."
	return nil
}
