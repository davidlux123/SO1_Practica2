package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameJson struct {
	Gameid   int32  `json:"gameid"`
	Players  int32  `json:"players"`
	Gamename string `json:"gamename"`
	Winner   int32  `json:"winner"`
	Queue    string `json:"queue"`
	Datetime string `json:"datetime"`
}

func GetCollection(collection string) *mongo.Collection {
	//"root:password@host:port

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	return client.Database("ProyectoF2").Collection(collection)
}

func InsertMongo(R string) (e error) {

	var r Result
	err := json.Unmarshal([]byte(R), &r)
	if err != nil {
		return err
	}

	data := GameJson{
		Gameid:   r.Game_id,
		Players:  r.Players,
		Gamename: r.Game_name,
		Winner:   r.Winner,
		Queue:    r.Queue,
		Datetime: r.Date_game,
	}

	var coleccion = GetCollection("Logs")
	coleccion.InsertOne(context.Background(), data)

	fmt.Print("Se ha Guardado en Mongo :" + R + "\n\n")
	return nil
}
