package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbUrl  = viper.Get("DB_URL")
	dbName = viper.Get("DB_NAME")
	dbPort = viper.Get("DB_PORT")
)

func ConnectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stringDeconexao := fmt.Sprintf("%s://%s:%s", dbName, dbUrl, dbPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(stringDeconexao))
	if err != nil {
		log.Panic("database connection error")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return client, nil
}
