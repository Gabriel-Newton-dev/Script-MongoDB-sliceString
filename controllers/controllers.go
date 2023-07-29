package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Gabriel-Newton-dev/Script-MongoDB-sliceString/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbName = viper.Get("DB_NAME")
)

func UpdateUserData(client *mongo.Client) error {
	dbNames := fmt.Sprintf("%s", dbName)
	collection := client.Database(dbNames).Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// definido filtro para recuperar todos os dados na colletcion que possua string
	filter := bson.M{"data": bson.M{"$type": "string"}}

	// Recuperar os documentos que correspondente ao filtro
	cursor, appErr := collection.Find(context.Background(), filter)
	if appErr != nil {
		fmt.Errorf("error")
	}

	defer cursor.Close(ctx)

	for cursor.Next(context.Background()) {
		var user models.Users
		if err := cursor.Decode(&user); err != nil {
			return err
		}
	}
	// Converter a string para Slice usando ","
	sliceData := strings.Split(user.Name, ",") // verificar erro.

	// Atualiza o documento com os novos dados
	update := bson.M{"$set": bson.M{"data": sliceData}}

	filter2 := bson.M{"id": user.ID}
	_, err := collection.UpdateOne(context.Background(), filter2, update)
	if err != nil {
		return err
	}
	return nil
}
