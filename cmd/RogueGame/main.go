package main

import (
	"log"
	"school21/task/RogueGame/internal/domain/models"
	"school21/task/RogueGame/internal/domain/models/actor"
	"school21/task/RogueGame/internal/domain/models/world"
	"school21/task/RogueGame/internal/infrastructure/frontend"

	"github.com/rthornton128/goncurses"
	gc "github.com/rthornton128/goncurses"
)

func main() {
	_, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	goncurses.Echo(false)
	goncurses.Cursor(0)

	// Размеры экрана
	// rows, cols := stdscr.MaxYX()

	gameField, err := frontend.NewGameField(12, 22, 0, 0)
	if err != nil {
		return
	}

	tiles := [][]*world.Tile{}
	for y := 0; y < 10; y++ {
		tiles = append(tiles, []*world.Tile{})
		for x := 0; x < 20; x++ {
			tileType := world.FloorType
			// Границы карты сделаем стенами
			if y == 0 || y == 9 || x == 0 || x == 19 {
				tileType = world.WallType
			}
			tiles[y] = append(tiles[y], &world.Tile{
				Type:     tileType,
				Position: models.Position2d{X: x, Y: y},
			})
		}
	}

	tiles[5][10].Entities = append(tiles[5][10].Entities, &actor.Actor{})

	gameField.Draw(&world.GameMap{Tiles: tiles})

	// stdscr.MovePrint(0, 0, "Push ESC.")
	// stdscr.Refresh()

	// Обработка клавиш в основном окне
	for {
		key := gameField.GetChar()
		if key == gc.KEY_ESC {
			break
		}
	}
}
