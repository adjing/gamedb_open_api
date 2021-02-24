package game_log

import (
	"context"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgo_client *mongo.Client
var m_MongoDBUrl = "mongodb://127.0.0.1:27017"

func StartInit() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m_MongoDBUrl))
	if err != nil {
		//log
	}
	mgo_client = client

}

func GetClient() *mongo.Client {
	if mgo_client == nil {
		StartInit()
	}
	return mgo_client

}

func Get_Collection(db_name string, datatable_name string) *mongo.Collection {

	c := GetClient()
	var dt = c.Database(db_name).Collection(datatable_name)
	return dt
}
