package database

import (
	"context"
	"github.com/skyepic/privateapi/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var CLIENT *mongo.Client
var DB *mongo.Database
var CTX context.Context

func Connect() {
	CLIENT, err := mongo.NewClient(options.Client().ApplyURI(config.MongodbUri))
	if err != nil {
		log.Println(err)
	}
	CTX, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = CLIENT.Connect(CTX); err != nil {
		log.Println("ConnectDbError", err)
		return
	}
	log.Println("Connection openned.")
	DB = CLIENT.Database("skyepic-main")
}

func Disconnect() {
	err := CLIENT.Disconnect(CTX)
	if err != nil {
		log.Println("DisconnectDbError", err)
		return
	}
	log.Println("Connection finalized.")
}
