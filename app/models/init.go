package models

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

const (
	DBHOST     = "localhost"
	DBNAME     = "unival"
	COLLECTION = "users"
	SALT       = "uniconsole"
)

var DBSession *mgo.Session
var err error

func init() {
	if DBSession, err = mgo.Dial(DBHOST); err != nil {
		revel.ERROR.Fatalf("Connecting to Database error: %s", err)
	}
}

func GetConnection() *mgo.Collection {
	return DBSession.DB("uniconsole").C("user")
}
