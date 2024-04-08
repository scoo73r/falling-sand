package falling_sand

import rl "github.com/gen2brain/raylib-go/raylib"

var sandMap [10][10]int

func drawMap() {

}

func main() {
	screenWidth := float32(1000)
	screenHeight := float32(500)

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Falling Sand")

	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rl.DrawCircle(rl.GetMouseX(), 10, 10, rl.Green)
			//	sandColumn := rl.GetMouseX()
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(int32(screenWidth)-90, int32(screenHeight)-30)

		rl.EndDrawing()
	}
	rl.CloseWindow()

}
