package dal

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type MongoDatastore struct {
	DB      *mongo.Database
	Session *mongo.Client
}

func NewMongoDatastore() *MongoDatastore {
	var mongoDataStore *MongoDatastore
	db, session := connect()
	if db != nil && session != nil {
		mongoDataStore = new(MongoDatastore)
		mongoDataStore.DB = db
		mongoDataStore.Session = session
		return mongoDataStore
	}

	return nil
}

func connect() (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo()
	})

	return db, session
}

func connectToMongo() (a *mongo.Database, b *mongo.Client) {
	uri := viper.GetString("connectionStrings.mongoConnectionString")
	db := viper.GetString("connectionStrings.database")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		//log.Fatalf("Failed to connect to database: %v", db)
		panic(err)
	}

	return client.Database(db), client
}
