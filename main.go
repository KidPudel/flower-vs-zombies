package main

import (
	"fmt"
	"net/http"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/KidPudel/flower-vs-zombiles/common"
)
type Model struct {
	UserName         string   `json:"user_name,omitempty"`
	UserOnlineStatus bool     `json:"user_online_status,omitempty"`
	Dick             string   `json:"dick,omitempty"`
	IsPenis          bool     `json:"is_penis"`
	Test             []string `json:"test"`
	Count            int      `json:"count"`
	Price            float64  `json:"price"`
}

type Game struct{}

func initGame() *Game {
	rl.InitWindow(common.WindowWidth, common.WindowHeight, "flower vs zombies")
	// rl.RenderTexture2DRenderTexture2D

	return &Game{}
}

func (game *Game) update() {
}

func (game Game) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Brown)
	rl.EndDrawing()
}

func (game *Game) run() {
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		game.update()
		game.draw()
	}
}

func main() {
	fmt.Println("hellope from Neovim, I think I could actually stick around this time")
	game := initGame()
	game.run()
	fmt.Errorf("aaa %s", "aaa")

    mux := http.NewServeMux()
    mux.HandleFunc("GET clients", func(w http.ResponseWriter, r *http.Request) {})
}
