package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Red-Chot/simple-rpg-api-Aula03/internal/handler"
	repository "github.com/Red-Chot/simple-rpg-api-Aula03/internal/repositoy"
	"github.com/Red-Chot/simple-rpg-api-Aula03/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	// "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
	dsn := "postgresql://postgres:postgres@localhost/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	playerRepository := repository.NewPlayerRepository(db)
	playerService := service.NewPlayerService(*playerRepository)
	playerHandler := handler.NewPlayerHandler(playerService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /player", playerHandler.AddPlayer)
	mux.HandleFunc("GET /player", playerHandler.LoadPlayers)
	mux.HandleFunc("DELETE /player/{id}", playerHandler.DeletePlayer)
	mux.HandleFunc("GET /player/{id}", playerHandler.LoadPlayer)
	mux.HandleFunc("PUT /player/{id}", playerHandler.SavePlayer)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}
