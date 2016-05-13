package main

import (
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"log"
)

type DbHandler struct {
	session *mgo.Session
	dbname  string
}

type Event struct {
	Hash string
	Name string
}

func NewDbHandler(url string) *DbHandler {
	s, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	d := new(DbHandler)
	d.session = s
	d.dbname = "mrs"
	return d
}

func (d *DbHandler) SetEvent(e *Event) {
	//defer session.Close()
	//session.SetMode(mgo.Monotonic, true)

	c := d.session.DB(d.dbname).C("event")
	err := c.Insert(e)

	if err != nil {
		log.Fatal(err)
	}
}

func (d *DbHandler) GetEventByHash(hash string) *Event {
	result := Event{}
	c := d.session.DB(d.dbname).C("event")
	err := c.Find(bson.M{"hash": hash}).One(&result)
	return result
}
