package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/davidlux123/service/src/controllers"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT"),
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"games-topic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n \n", msg.TopicPartition, string(msg.Value))

			/*/Tidbi
			if errmysql := controllers.InsertarTidb(string(msg.Value)); errmysql != nil {
				fmt.Printf("Mysql error: %v \n\n", errmysql)
			}
			//redis
			if errredis := controllers.InsertarRedis(string(msg.Value)); errredis != nil {
				fmt.Printf("Mysql error: %v \n\n", errredis)
			}*/
			//mongo
			if errmongo := controllers.InsertMongo(string(msg.Value)); errmongo != nil {
				fmt.Printf("Mysql error: %v \n\n", errmongo)
			}

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
	c.Close()
}
