package dao

import (
	"log"

	. "github.com/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MeetingDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "meetings"
)

//Establish Connection using Connect...
func (m *MeetingDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//Find the Scheduled Meetings Using FindAll
func (m *MeetingDAO) FindAll() ([]Meetings, error) {
	var meetings []Meetings
	err := db.C(COLLECTION).Find(bson.M{}).All(&meetings)
	return meetings, err
}

//Find Meetings By Id
func (m *MeetingDAO) FindByID(id string) (Meetings, error) {
	var meeting Meetings
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&meeting)
	return meeting, err
}
//Adding meeting
func (m *MeetingDAO) Insert(meeting Meetings) error {
	err := db.C(COLLECTION).Insert(&meeting)
	return err
}
