package usecases

import (
	"fmt"
	"school21/task/RogueGame/internal/domain/models"
	"school21/task/RogueGame/internal/domain/repositories"
)

type MoveActorUseCase struct {
	repo repositories.IActorRepository
}

func NewMoveActorUseCase(r repositories.IActorRepository) *MoveActorUseCase {
	return &MoveActorUseCase{r}
}

func (uc MoveActorUseCase) Execute(actorId int, newPos models.Position2d) {
	actor, err := uc.repo.FindById(actorId)

	if err != nil {
		fmt.Println(err)
		return
	}
	actor.Position = newPos

	uc.repo.Save(actor)
}
