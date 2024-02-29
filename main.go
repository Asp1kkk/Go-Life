package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) String() string {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"
	var result string
	for i := range w.Cells {
		for j := range w.Cells[i] {
			if w.Cells[i][j] {
				result += greenSquare
			} else {
				result += brownSquare
			}
		}
		if i+1 != len(w.Cells) {
			result += "\n"
		}
	}
	return result
}

func NewWorld(height, width int) *World {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) Neighbors(x, y int) int {
	var result int

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			newX := (w.Width + x + dx + w.Width) % w.Width
			newY := (w.Height + y + dy + w.Height) % w.Height

			if w.Cells[newY][newX] {
				result++
			}
		}
	}

	return result
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	}
	if n == 3 && !alive {
		return true
	}

	return false
}

func (w *World) Seed() {
	for _, row := range w.Cells {
		for i := range row {
			if rand.Intn(5) == 1 {
				row[i] = true
			}
		}
	}
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(i, j)
		}
	}
}

func (w *World) SaveState(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for i, row := range w.Cells {
		for _, val := range row {
			if val {
				fmt.Fprint(f, 1)
			} else {
				fmt.Fprint(f, 0)
			}
		}
		if i+1 != len(w.Cells) {
			fmt.Fprint(f, "\n")
		}
	}

	return nil
}

func (w *World) LoadState(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	slice := strings.Split(string(data), "\n")

	for i := 1; i < len(slice); i++ {
		if len(slice[i]) != len(slice[0]) {
			return errors.New("несоответствие размерности в строках файла")
		}
	}

	w.Height = len(slice)
	w.Width = len(slice[0])
	w.Cells = make([][]bool, w.Height)

	for i, str := range slice {
		w.Cells[i] = make([]bool, w.Width)
		for j, v := range str {
			if v == '1' {
				w.Cells[i][j] = true
			} else {
				w.Cells[i][j] = false
			}
		}
	}
	return nil
}

func main() {
	height := 10
	width := 10

	currentWorld := NewWorld(height, width)

	nextWorld := NewWorld(height, width)

	currentWorld.Seed()

	for {
		fmt.Println(currentWorld)

		NextState(currentWorld, nextWorld)

		currentWorld, nextWorld = nextWorld, currentWorld

		time.Sleep(2 * time.Second)

		fmt.Print("\033[H\033[2J")
	}
}
