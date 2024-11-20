package main

import (
	"math/rand"
	"strings"
	"time"

	"gofr.dev/pkg/gofr"
)

type StatusResponse struct {
	Status      string `json:"status"`
	Emoji       string `json:"emoji"`
	Activity    string `json:"activity"`
	TimeElapsed string `json:"time_elapsed"`
	GameStatus  string `json:"game_status,omitempty"`
}

var activities = []string{
	"Watching Netflix",
	"Grinding Valorant",
	"Touching grass (rare)",
	"In a Discord call",
	"Vibing to spotify",
	"Doing absolutely nothing",
}

var emojis = []string{"🎮", "🎧", "💀", "✨", "🌟", "🎯", "🔥", "💯"}

var gameStatuses = map[string][]string{
	"Valorant":  {"Hardstuck Iron", "Radiant BTW", "Whiffing every shot"},
	"Minecraft": {"Finding diamonds", "Getting lost again", "Building a dirt house"},
	"Among Us":  {"Looking kinda sus", "Emergency Meeting!", "In electrical..."},
}

func generateTimeElapsed() string {
	hours := rand.Intn(12) + 1
	return strings.Join([]string{string(rune(hours)), "hours"}, " ")
}

func main() {
	app := gofr.New()

	//Generate random status
	app.GET("/discord/status", func(ctx *gofr.Context) (interface{}, error) {

		rand.Seed(time.Now().UnixNano())

		game := ctx.Param("game")
		activity := activities[rand.Intn(len(activities))]
		emoji := emojis[rand.Intn(len(emojis))]

		response := StatusResponse{
			Status:      "online",
			Emoji:       emoji,
			Activity:    activity,
			TimeElapsed: generateTimeElapsed(),
		}

		if game != "" && gameStatuses[game] != nil {
			response.GameStatus = gameStatuses[game][rand.Intn(len(gameStatuses[game]))]
		}

		return response, nil
	})

	app.Run()

}
