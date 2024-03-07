package main

import (
	life "github.com/Asp1kkk/Go-Life/game"

	"fmt"
)

func main() {
	world := life.NewWorld(10, 10)
	world.RandInit(30)
	fmt.Println(world)
}
