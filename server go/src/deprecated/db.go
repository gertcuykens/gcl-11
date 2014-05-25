package deprecated

import (
	"appengine/datastore"
)

type Data struct {
	Key *datastore.Key `datastore:"-"`
	Data string `datastore:"data"`
    Token *Token `datastore:"-"`
}

func (d *Data) Put() error{
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

/*func (s *Service) DataService(r *http.Request, req *Token, res *Token) (err error) {
	return nil
}*/

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
	if err = datastore.Get(u.Token.Context, u.Key, u); err != nil {u.Token.Status="Datastore get error! "+err.Error(); return err}
	u.Token.Status="Fetched "+u.Key.String()+"."
	return nil
}

/*
  if req.Limit <= 0 {req.Limit = 10}
  c := endpoints.NewContext(r)
  q := datastore.NewQuery("Greeting").Order("-Date").Limit(req.Limit)
  greets := make([]*Greeting, 0, req.Limit)
  keys, err := q.GetAll(c, &greets)
  if err != nil {return err}
  for i, k := range keys {greets[i].Key = k}
  resp.Items = greets
*/
