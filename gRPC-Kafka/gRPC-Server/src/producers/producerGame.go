package producers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Result struct {
	Game_id   int32  `json:"game_id"`
	Players   int32  `json:"players"`
	Game_name string `json:"game_name"`
	Winner    int32  `json:"winner"`
	Queue     string `json:"queue"`
	Date_game string `json:"date_game"`
}

func SaveToKafka(gameId, players int32, date string) string {
	result := getResult(gameId, players, date)
	fmt.Println(string(result))

	// Produce messages to topic (asynchronously)
	nvoProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")})
	if err != nil {
		log.Fatalf("Error:no se pudo instanciar el producer :(. Error: %s", err.Error())
		panic(err)
	}

	topic := "games-topic"
	nvoProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(result),
	}, nil)

	return string(result)
}

func getResult(gameId, players int32, date string) string {

	var winner int

	//aca iria el if de los 5 juegos
	rand.Seed(time.Now().Unix())
	winner = rand.Intn(int(players))

	juegoName := "Juego " + strconv.Itoa(rand.Intn(5))

	jsonBytes, err := json.Marshal(Result{
		Game_id:   gameId,
		Players:   players,
		Game_name: juegoName,
		Winner:    int32(winner),
		Queue:     "kafka",
		Date_game: date,
	})

	if err != nil {
		log.Fatalf("Error: No se puede Convertir struct a json. Error: %s", err.Error())
		panic(err)
	}
	return string(jsonBytes)
}
