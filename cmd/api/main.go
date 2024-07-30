package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Red-Chot/simple-rpg-api-Aula03/internal/handler"
	"github.com/Red-Chot/simple-rpg-api-Aula03/internal/repository"
	"github.com/Red-Chot/simple-rpg-api-Aula03/internal/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgresql://postgres:postgres@localhost:5432/go-simple-rpg-api?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	 playerRepository := repository.NewPlayerRepository(db)
	playerService := service.NewPlayerService(*playerRepository)
	playerHandler := handler.NewPlayerHandler(playerService)

	enemyRepository := repository.NewEnemyRepository(db)
	enemyService := service.NewEnemyService(*enemyRepository)
	 enemyHandler := handler.NewEnemyHandler(enemyService)

	battleRepository := repository.NewBattleRepository(db)
	 battleService := service.NewBattleService(*playerRepository, *enemyRepository, *battleRepository)
	 battleHandler := handler.NewBattleHandler(battleService)
 
	router := mux.NewRouter()

	 router.HandleFunc("/player", playerHandler.AddPlayer).Methods("POST")
	router.HandleFunc("/player", playerHandler.LoadPlayers).Methods("GET")
	router.HandleFunc("/player/{id}", playerHandler.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/player/{id}", playerHandler.LoadPlayer).Methods("GET")
	router.HandleFunc("/player/{id}", playerHandler.SavePlayer).Methods("PUT")

	router.HandleFunc("/enemy", enemyHandler.AddEnemy).Methods("POST")
	router.HandleFunc("/enemy", enemyHandler.LoadEnemies).Methods("GET")
	 router.HandleFunc("/enemy/{id}", enemyHandler.DeleteEnemy).Methods("DELETE")
	router.HandleFunc("/enemy/{id}", enemyHandler.LoadEnemy).Methods("GET")
	router.HandleFunc("/enemy/{id}", enemyHandler.SaveEnemy).Methods("PUT")

	router.HandleFunc("/battle", battleHandler.CreateBattle).Methods("POST")
	router.HandleFunc("/battles", battleHandler.LoadBattles).Methods("GET")

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}





























