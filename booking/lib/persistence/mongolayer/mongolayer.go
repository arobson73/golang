package mongolayer

import (
	"andy/booking/lib/persistence"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB         = "myevents"
	USERS      = "users"
	ADMINUSERS = "adminuser"
	EVENTS     = "events"
	LOCATIONS  = "locations"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	log.Println("NewMongoDBLayer")
	s, err := mgo.Dial(connection)
	return &MongoDBLayer{
		session: s,
	}, err
}

func (mgoLayer *MongoDBLayer) AddUser(u persistence.User) ([]byte, error) {
	log.Println("mgoLayer.AddUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !u.ID.Valid() {
		u.ID = bson.NewObjectId()
	}

	//u.ID = bson.NewObjectId()
	return []byte(u.ID), s.DB(DB).C(USERS).Insert(u)
}
func (mgoLayer *MongoDBLayer) AddAdminUser(u persistence.AdminUser) ([]byte, error) {
	log.Println("mgoLayer.AddAdminUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !u.ID.Valid() {
		u.ID = bson.NewObjectId()
	}

	//u.ID = bson.NewObjectId()
	return []byte(u.ID), s.DB(DB).C(ADMINUSERS).Insert(u)
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	log.Println("mgoLayer.AddEvent")
	s := mgoLayer.getFreshSession()
	defer s.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if e.Location.ID == "" {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

func (mgoLayer *MongoDBLayer) AddLocation(l persistence.Location) (persistence.Location, error) {
	log.Println("mgoLayer.AddLocation")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !l.ID.Valid() {
		l.ID = bson.NewObjectId()
	}
	err := s.DB(DB).C(LOCATIONS).Insert(l)
	return l, err
}

func (mgoLayer *MongoDBLayer) AddEventForAdminUser(id []byte, e persistence.Event) error {
	log.Println("mgoLayer.AddEventForAdminUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if e.Location.ID == "" {
		e.Location.ID = bson.NewObjectId()
	}

	//not this push and array so we had array of arrays of structs, instead of array of structs (i.e array of bookings)
	//return s.DB(DB).C(USERS).UpdateId(bson.ObjectId(id), bson.M{"$addToSet": bson.M{"bookings": []persistence.Booking{bk}}})
	//return s.DB(DB).C(ADMINUSERS).UpdateId(bson.ObjectId(id), bson.M{"$push": bson.M{"bookings": bson.M{"date": bk.Date, "eventid": bk.EventID, "seats": bk.Seats, "name": bk.Name}}})
	//	return s.DB(DB).C(ADMINUSERS).UpdateId(bson.ObjectId(id), bson.M{"$push": bson.M{"events": bson.M{"id": e.ID, "name": e.Name, "duration": e.Duration, "startdate": e.StartDate, "enddate": e.EndDate, bson.M{"location": bson.M{"id": e.Location.ID, "name": e.Location.Name,
	//"address": e.Location.Address, "country": e.Location.Country, "opentime": e.Location.OpenTime, "closetime": e.Location.CloseTime, bson.M{"halls": e.Location.Halls}}}}}})
	return s.DB(DB).C(ADMINUSERS).UpdateId(bson.ObjectId(id), bson.M{"$push": bson.M{"events": bson.M{"id": e.ID, "name": e.Name, "duration": e.Duration, "startdate": e.StartDate, "enddate": e.EndDate, "location": e.Location}}})

}

func (mgoLayer *MongoDBLayer) AddBookingForUser(id []byte, bk persistence.Booking) error {
	log.Println("mgoLayer.AddBookingForUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	//not this push and array so we had array of arrays of structs, instead of array of structs (i.e array of bookings)
	//return s.DB(DB).C(USERS).UpdateId(bson.ObjectId(id), bson.M{"$addToSet": bson.M{"bookings": []persistence.Booking{bk}}})
	return s.DB(DB).C(USERS).UpdateId(bson.ObjectId(id), bson.M{"$push": bson.M{"bookings": bson.M{"date": bk.Date, "eventid": bk.EventID, "seats": bk.Seats, "name": bk.Name}}})
}

func (mgoLayer *MongoDBLayer) FindUser(f string, l string) (persistence.User, error) {
	log.Println("mgoLayer.FindUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.User{}
	err := s.DB(DB).C(USERS).Find(bson.M{"first": f, "last": l}).One(&u)
	//fmt.Printf("Found %v \n", u.String())
	return u, err
}
func (mgoLayer *MongoDBLayer) FindUserEmailPass(e string, p string) (persistence.User, error) {
	log.Println("mgoLayer.FindUserEmail")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.User{}
	err := s.DB(DB).C(USERS).Find(bson.M{"email": e, "password": p}).One(&u)
	//fmt.Printf("Found %v \n", u.String())
	return u, err
}
func (mgoLayer *MongoDBLayer) FindAdminUserEmailPass(e string, p string) (persistence.AdminUser, error) {
	log.Println("mgoLayer.FindAdminUserEmail")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.AdminUser{}
	err := s.DB(DB).C(ADMINUSERS).Find(bson.M{"email": e, "password": p}).One(&u)
	//fmt.Printf("Found %v \n", u.String())
	return u, err
}

func (mgoLayer *MongoDBLayer) FindBookingsForUser(id []byte) ([]persistence.Booking, error) {
	log.Println("mgoLayer.FindBookingsForUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.User{}
	err := s.DB(DB).C(USERS).FindId(bson.ObjectId(id)).One(&u)
	//log.Println("email:", u.Email)
	//log.Println("bookings ")
	//log.Println("bookins type=", reflect.TypeOf(u.Bookings))
	//for _, b := range u.Bookings {
	//log.Println("booking:", b.Date)
	//}
	return u.Bookings, err
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	log.Println("mgoLayer.FindEvent")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}

	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAdminUser(id []byte) (persistence.AdminUser, error) {
	log.Println("mgoLayer.FindAdminUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.AdminUser{}

	err := s.DB(DB).C(ADMINUSERS).FindId(bson.ObjectId(id)).One(&u)
	return u, err
}

//not sure we need this.
func (mgoLayer *MongoDBLayer) FindUserFromId(id []byte) (persistence.User, error) {
	log.Println("mgoLayer.FindUser")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	u := persistence.User{}

	err := s.DB(DB).C(USERS).FindId(bson.ObjectId(id)).One(&u)
	return u, err
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	log.Println("mgoLayer.FindEventByName")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	log.Println("mgoLayer.FindAllAvailableEvents")
	s := mgoLayer.getFreshSession()
	defer s.Close()
	events := []persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (l *MongoDBLayer) FindLocation(id string) (persistence.Location, error) {
	log.Println("l.FindLocation")
	s := l.getFreshSession()
	defer s.Close()
	location := persistence.Location{}
	err := s.DB(DB).C(LOCATIONS).Find(bson.M{"_id": bson.ObjectId(id)}).One(&location)
	return location, err
}

func (l *MongoDBLayer) FindAllLocations() ([]persistence.Location, error) {
	log.Println("l.FindAllLocations")
	s := l.getFreshSession()
	defer s.Close()
	locations := []persistence.Location{}
	err := s.DB(DB).C(LOCATIONS).Find(nil).All(&locations)
	return locations, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}
