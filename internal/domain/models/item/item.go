package item

import (
	"school21/task/RogueGame/internal/domain/models"
)

type Item struct {
	ID       int
	Name     string
	Position models.Position2d // FIX MEEEEEEE (delete and make metod for word.GameMap)
	Value    int               // Price
}
