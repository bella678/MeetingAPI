package models
import(
    "gopkg.in/mgo.v2/bson"
)
type Meetings struct {
		ID                bson.ObjectId `bson:"_id" json:"id"`
		Title             string        `bson:"title" json:"title"`
		Participats       []string      `bson:"participants" json:"participants"`
		StartTime         string        `bson:"starttime" json:"starttime"`
		EndTime           string        `bson:"endtime" json:"endtime"`
		CreationTimestamp string        `bson:"creationimestamp" json:"creationtimestamp"`
	}

	type Participants struct {
		Name  string `bson:"name" json:"name"`
		Email string `bson:"email" json:"email"`
		RSPV  string `bson:"rspv" json:"rspv"`
	}
