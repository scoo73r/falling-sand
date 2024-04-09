package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth = float32(1000)
	screenHeight = float32(1000)
)

var (
		sandMap [int(screenWidth)][int(screenHeight)]int
		gameTick int

)
func drawMap() {
		for x := 0; x < int(screenWidth); x++ {
			for y := 0; y < int(screenHeight); y++ {
			   if sandMap[x][y] == 1 {
					rl.DrawCircle(int32(x) , int32(y), 30, rl.White)
			}
		}
	}
}

func input() {
if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			sandMap[rl.GetMouseX()][rl.GetMouseY()] = 1
		}
}

func gameLoop() {
	for x := int(screenWidth)-1; x > 0; x--{
		for y := int(screenHeight)-1; y > 0; y-- {
		   if sandMap[x][y] == 0 && y < (int(screenHeight) - 1 ){
				if sandMap[x][y-1] != 0{
				sandMap[x][y] = sandMap[x][y-1] 
                sandMap[x][y-1] = 0
			}
		}
	}
}
}

func main() {


	rl.SetTargetFPS(60)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Falling Sand")
	gameTick = 0
	for !rl.WindowShouldClose() {

		gameTick++	
		rl.BeginDrawing()
		input()
		//if gameTick == 10 {
      	gameLoop()
		//gameTick = 0}
		drawMap()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(int32(screenWidth)-90, int32(screenHeight)-30)
		rl.DrawText( fmt.Sprintf("%d", gameTick), int32(screenWidth)-90, int32(screenHeight)-60, 14, rl.White)
		rl.EndDrawing()
	}
	rl.CloseWindow()

}
