package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/physics"
)

func main() {
	
	screenWidth := float32(800)
	screenHeight := float32(450)
	
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Falling Sand")

	physics.Init()

	floor := physics.NewBodyRectangle(rl.NewVector2(screenWidth/2, screenHeight), screenWidth + 4, 100, 10)
	wallLeft := physics.NewBodyRectangle(rl.NewVector2(5, screenHeight/2), 10, screenHeight, 10)
	wallRight := physics.NewBodyRectangle(rl.NewVector2(screenWidth-4, screenHeight/2), 10, screenWidth, 10)
	

	floor.Enabled = false
	floor.Restitution = 1
	wallLeft.Enabled = false
	wallRight.Enabled = false

	//Main Loop
	for !rl.WindowShouldClose() {
		physics.Update()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			sand := physics.NewBodyCircle(rl.GetMousePosition(), 10, 10)
			sand.Restitution = 0
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(int32(screenWidth)-90, int32(screenHeight)-30)

		for i, body := range physics.GetBodies() {
			vertexCount := physics.GetShapeVerticesCount(i)
			for j := 0; j < vertexCount; j++ {
				vertexA := body.GetShapeVertex(j)

				jj := 0
				if j+1 < vertexCount {
					jj = j + 1
				}
				vertexB := body.GetShapeVertex(jj)
				rl.DrawLineV(vertexA, vertexB, rl.Green)
			}
		}		
		rl.EndDrawing()
	}
	physics.Close()
	rl.CloseWindow()
}
