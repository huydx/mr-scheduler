package main

import (
	mgo "gopkg.in/mgo.v2"
	"log"
)

type DbHandler struct {
	session *mgo.Session
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
	return d
}

func (d *DbHandler) SetEvent(e *Event) {
	//defer session.Close()
	//session.SetMode(mgo.Monotonic, true)

	c := d.session.DB("mrs").C("event")
	err := c.Insert(e)

	if err != nil {
		log.Fatal(err)
	}
}
