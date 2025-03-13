package world

import (
	"school21/task/RogueGame/internal/domain/models"
	"school21/task/RogueGame/internal/domain/models/actor"
	"school21/task/RogueGame/internal/domain/models/item"
)

type GameMap struct {
	Tiles []*Tile
}

type Tile struct {
	Position models.Position2d
	Entities []*actor.Actor
	Items    []*item.Item
}
