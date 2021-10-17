package db

import (
	"context"
	"fmt"
	"gofiber/internal/model"
	"time"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
	Student  *mongo.Collection
}

var _ Connection = (*Mongo)(nil)
var _ Database = (*Mongo)(nil)
var _ Collection = (*Mongo)(nil)

//Function Connect will create a new session for mongo db
func (m *Mongo) Connect() string {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, e := mongo.Connect(context.TODO(), clientOptions)
	CheckError(e)

	// Check the connection
	e = client.Ping(context.TODO(), nil)
	CheckError(e)

	m.Client = client

	return "Mongo Connected!"
}

// List databases
func (m *Mongo) List() string {
	databases, err := m.Client.ListDatabaseNames(context.TODO(), bson.M{})
	CheckError(err)
	s := ""
	for _, a := range databases {
		s += "Database: " + a
	}
	fmt.Println(databases)

	database := m.Client.Database("jai")
	m.Database = database

	collections, err := m.Database.ListCollectionNames(context.TODO(), bson.M{})
	CheckError(err)
	for _, a := range collections {
		s += "Collections: " + a
	}
	fmt.Println(collections)
	return s
}

func (m *Mongo) Init() {
	student := m.Database.Collection("student")
	m.Student = student
}

func (m *Mongo) Insert() {
	res, err := m.Student.InsertOne(context.TODO(),
		model.Student{
			Phone:     "65345343324",
			Email:     "fesrewrew@mongodb.com",
			Status:    "Incomplete",
			CreatedAt: time.Now(),
		})
	CheckError(err)
	fmt.Println(res.InsertedID)
}

func (m *Mongo) Update() {
	update := bson.D{
		{"$set",
			bson.D{
				{"phone", "999999999"},
			}},
	}

	upd := m.Student.FindOneAndUpdate(
		context.TODO(),
		model.Student{Status: "Incomplete"},
		update,
	)
	b := model.Student{}
	upd.Decode(&b)
	fmt.Println(b)
}
func (m *Mongo) Find() {
	s := model.Student{}
	res, err := m.Student.Find(context.TODO(), model.Student{Status: "Incomplete"})
	CheckError(err)
	for res.Next(context.TODO()) {
		err := res.Decode(&s)
		CheckError(err)
		fmt.Println("Student Record:", s.Email, s.Phone, s.Status)
	}
}

func (m *Mongo) Delete() {
}

func CheckError(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}

func RecoverError() {
	if r := recover(); r != nil {
		fmt.Println("Panic:", r)
	}
}
