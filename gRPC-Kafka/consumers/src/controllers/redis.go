package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Games []Game

type Player struct {
	Name string `json:"name"`
}

type Game struct {
	Name     string   `json:"name"`
	ID       int64    `json:"id"`
	Players  []Player `json:"players"`
	Winner   string   `json:"winner"`
	Broker   string   `json:"broker"`
	Datetime string   `json:"datetime"`
}

func newPool() *redis.Pool {
	//35.184.204.155:6379

	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", os.Getenv("REDIS_URI"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func UnmarshalGames(data []byte) (Games, error) {
	var r Games
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Games) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

var pool = newPool()

func InsertarRedis(R string) (e error) {

	var r Result
	err := json.Unmarshal([]byte(R), &r)
	if err != nil {
		return err
	}

	var plys []Player
	for i := 1; i <= int(r.Players); i++ {
		plys = append(plys, Player{Name: strconv.Itoa(int(i))})
	}

	juego := Game{
		Name:     r.Game_name,
		ID:       int64(r.Game_id),
		Players:  plys,
		Winner:   strconv.Itoa(int(r.Winner)),
		Broker:   r.Queue,
		Datetime: time.Now().Format("2006-01-02 15:04:05"),
	}

	client := pool.Get()
	defer client.Close()

	value, err := redis.String(client.Do("GET", "games"))
	if err != nil {
		juegos := Games{}
		juegos = append(juegos, juego)
		texto, _ := juegos.Marshal()
		client.Do("SET", "games", texto)
	} else {
		juegos, _ := UnmarshalGames([]byte(value))
		juegos = append(juegos, juego)
		texto, _ := juegos.Marshal()
		client.Do("SET", "games", texto)

	}

	data2, err := client.Do("GET", "games")
	if err != nil {
		return err
	}

	fmt.Printf("Se ha Guardado en Redis : %s \n\n ", data2)
	return nil
}
