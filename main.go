package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"time"

	"github.com/fthyilmz/go-challange/src"

	"github.com/joho/godotenv"
)

var (
	db          *src.DB
	userService src.Service
)

func loadDependencies() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	apiTimeout, _ := strconv.Atoi(os.Getenv("BACKEND_API_TIMEOUT"))
	timeout := time.Duration(int64(apiTimeout))
	db = src.NewDB()
	userService = src.NewService(db, src.NewClient(os.Getenv("BACKEND_API_URL"), timeout))
}

func simulate(n int) {
	fmt.Println("--- Simulation started ---")
	userService.CreateRandomUsers(n)
	fmt.Println(n, "%c users created.")
	done := make(chan bool)

	for i := 0; i < n/2; i++ {
		randPlayers := userService.GetRandomUnPlayedUsers(2)
		for _, user := range randPlayers {
			user.Score = src.GenerateRandomInt(10) + 1
			user.IsPlayed = true
			db.Update(user.ID, *user)
		}
		go func(randPlayers []*src.User) {
			err := userService.SavePlayersScore(randPlayers)
			if err != nil {
				log.Error(err)
			}
			done <- true
		}(randPlayers)
	}
	// wait for all goroutines to complete before exiting
	for i := 0; i < n/2; i++ {
		<-done
	}
	fmt.Println("--- Simulation ended ---")
}

func main() {
	loadDependencies()
	simulate(100)
}
