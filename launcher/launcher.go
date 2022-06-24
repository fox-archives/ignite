package launcher

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var isActive bool = false

func Launch() {
	fmt.Println("Launching")
	if !isActive {
		isActive = true
		createWindow()
		isActive = false
	}
}

var items int32 = 0

func AddItem(screenWidth int32, text string) {
	var margin int32 = 5

	var rectangleHeight int32 = 40
	rectangleStart := margin + ((rectangleHeight + margin) * items)

	rl.DrawRectangle(margin, rectangleStart, screenWidth-(margin*2), rectangleHeight, rl.Gray)
	rl.DrawText(text, margin*2, rectangleStart+2, 32, rl.Black)
	items += 1
}

func createWindow() {
	var screenWidth int32 = 800
	var screenHeight int32 = 450

	rl.InitWindow(screenWidth, screenHeight, "Ignite 🔥")
	rl.SetTargetFPS(60)

	pos := rl.GetWindowPosition()

	fmt.Printf("AAAAAAAAAAAAAAAAA %f, %f\n", pos.X, pos.Y)
	for !rl.WindowShouldClose() {
		items = 0
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		AddItem(screenWidth, "Open Browser")
		AddItem(screenWidth, "Open VSCode")
		// AddItem(screenHeight)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
