package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	."github.com/bella678/MeetingAPI/models"
	."github.com/bella678/MeetingAPI/DataAccess"
	."github.com/bella678/MeetingAPI/Config"
)
 var config=Cob=nfig{}
 var dao=MeetingDAO{}

func scheduleMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meeting meetings
	if err:=json.NewDecoder(r.Body).Decode(&meetings);err!=nil{
		responseWithError(w,http.StatusBadRequest,"Invalid Request PayLoad")
		return
	}
	movie.ID=bson.NewObjectId()
	if err:=da0.Insert(meeting);err!=nil{
		respondWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	respondWithJson(w,http.StatusCreated,meeting)
	
}

func getMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	params:=mux.Vars(r)
	meeting,err:=dao.FindById(params["id"])
	if err!=nil{
		respondWithError(w,http.StatusBadRequest,"Invalid Meeting Id")
		return
	}
	respondWithJson(w http.StatusOK,meeting)
}

func respondWithError(w http.RespondWriter,code int,msg string){
	respondWithJson(w,code,map[string]string{"error":msg})
}
func respondWithJson(w http.ResponseWriter,code int,payload interface{}){
	response,_:=json.Marshal(payload)
	w.Header().set("Content-Type","application/json")
	w.WriterHeader(code)
	w.Write(response)
}

func init(){
	config.Read()
	dao.Server=config.Server
	dao.Database=config.Database
	dao.Connect()
}

func listMeetingInTimeFrameEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not yet implemented!")
}

func listMeetingOfParticipantEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not yet implemented!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/meetings", scheduleMeetingEndPoint).Methods("POST")
	r.HandleFunc("/meetings/{id}", getMeetingEndPoint).Methods("GET")
	r.HandleFunc("/meetings?start={starttime}&end={endtime}", listMeetingInTimeFrameEndPoint).Methods("GET")
	r.HandleFunc("/meetings?participant={email}", listMeetingOfParticipantEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}

