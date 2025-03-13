package world

import (
	"school21/task/RogueGame/internal/domain/models"
	"school21/task/RogueGame/internal/domain/models/actor"
	"school21/task/RogueGame/internal/domain/models/item"
)

type GameMap struct {
	Tiles [][]*Tile
}

type TileType = int

const (
	EmptyType TileType = iota
	FloorType
	WallType
)

type Tile struct {
	Type     TileType
	Position models.Position2d
	Entities []*actor.Actor
	Items    []*item.Item
}
