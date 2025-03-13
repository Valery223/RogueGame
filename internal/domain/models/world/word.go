package world

import (
	"school21/task/RogueGame/internal/domain/models"
	"school21/task/RogueGame/internal/domain/models/actor"
	"school21/task/RogueGame/internal/domain/models/item"
)

type GameMap struct {
	Tiles    [][]*Tile
	Actors   map[int]*actor.Actor        // actor by id(ID - > Actor)
	ActorPos map[models.Position2d][]int // all ID in the Position (Pos -> list ID Actors)
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

func (m *GameMap) AddActor(a *actor.Actor, p models.Position2d) {
	m.Actors[a.ID] = a
	m.ActorPos[p] = append(m.ActorPos[p], a.ID)
}

func (m GameMap) GetActor(id int) (*actor.Actor, bool) {
	return m.Actors[id]
}
