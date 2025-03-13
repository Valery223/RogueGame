package item

import (
	"school21/task/RogueGame/internal/domain/models"
)

type Item struct {
	ID       int
	Name     string
	Position models.Position2d
	Value    int // Price
}
