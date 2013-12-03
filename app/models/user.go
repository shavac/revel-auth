package models

import (
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Roles    []string
}

func Authenticate(username, password string) *User {
	user := User{}
	c := GetConnection()
	if err := c.Find(bson.M{"username": username, "password": crypt(password)}).One(&user); err != nil {
		return nil
	}
	return &user
}
