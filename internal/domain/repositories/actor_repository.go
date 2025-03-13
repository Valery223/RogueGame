package repositories

import "school21/task/RogueGame/internal/domain/models/actor"

type IActorRepository interface {
	FindById(id int) (*actor.Actor, error)
	Save(actor *actor.Actor) error
}
