package database

import "gopkg.in/mgo.v2"

/*
Database session
*/
var Session *mgo.Session

/*
Workype's model connection
*/
var Users *mgo.Collection

var Company *mgo.Collection

var Address *mgo.Collection

/*
Init database
*/
func Init(uri, dbname string) error {
	session, err := mgo.Dial(uri)
	if err != nil {
		return err
	}

	// See https://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	// Expose session and models

	Session = session
	Users = session.DB(dbname).C("user")
	Company = session.DB(dbname).C("company")
	Address = session.DB(dbname).C("address")
	return nil
}