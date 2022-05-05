package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Result struct {
	Game_id   int32  `json:"game_id"`
	Players   int32  `json:"players"`
	Game_name string `json:"game_name"`
	Winner    int32  `json:"winner"`
	Queue     string `json:"queue"`
	Date_game string `json:"date_game"`
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	//usuario:contraseña@tcp(host:port)

	uri := os.Getenv("TIDB_URI")
	nombreBaseDeDatos := "sopes1"
	db, err := sql.Open("mysql", fmt.Sprintf("%s/%s", uri, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertarTidb(R string) (e error) {

	var r Result
	err := json.Unmarshal([]byte(R), &r)
	if err != nil {
		return err
	}

	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO dataGame (game_id, players, game_name, winner, queue, date_game) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(r.Game_id, r.Players, r.Game_name, r.Winner, r.Queue, r.Date_game)
	if err != nil {
		return err
	}

	fmt.Print("Se ha Guardado en Ti-db :" + R + "\n\n")
	return nil
}
