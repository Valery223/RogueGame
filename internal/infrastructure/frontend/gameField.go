package frontend

import (
	"school21/task/RogueGame/internal/domain/models/world"

	"github.com/rthornton128/goncurses"
)

type GameField struct {
	win            *goncurses.Window
	width, height  int
	startX, startY int
}

func NewGameField(height, width, y, x int) (*GameField, error) {
	win, err := goncurses.NewWindow(height, width, y, x)
	if err != nil {
		return nil, err
	}
	return &GameField{
		win:    win,
		height: height,
		width:  width,
		startY: y,
		startX: x,
	}, nil
}

func (gf *GameField) Draw(gameMap *world.GameMap) {
	gf.win.Box(0, 0)

	var char goncurses.Char
	for _, tiles := range gameMap.Tiles {
		for _, tile := range tiles {
			if len(tile.Entities) > 0 {
				char = '@'
			} else {
				switch tile.Type {
				case world.EmptyType:
					char = ' '
				case world.FloorType:
					char = '.'
				case world.WallType:
					char = '#'
				default:
					char = '!'
				}
			}
			// Отрисовываем символ в позиции тайла (с учётом смещения рамки окна)
			gf.win.MoveAddChar(tile.Position.Y+1, tile.Position.X+1, char)
		}
	}

	gf.win.Refresh()
}

func (gf *GameField) GetChar() goncurses.Key {
	return gf.win.GetChar()
}
