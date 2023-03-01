package database

import (
	"context"
	"fmt"
	"time"

	mongoDb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongoDb.Database

func InitializeMongoConnection(username string, password string, host string, port int, dbname string) *mongoDb.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	dbUrl := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
	client, err := mongoDb.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		fmt.Println("connectionerror:", err)
	}
	DB = client.Database(dbname)
	return DB
}
