package db

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Resource struct {
	DB *mongo.Database
}

// Close use this method to close database connection
func (r *Resource) Close() {
	logrus.Warning("Closing all db connections")
}

func InitResource() (*Resource, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print(err)
	}

	host :=os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(host))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return &Resource{DB: mongoClient.Database(dbName)}, nil
}
