package repositories

import (
	"errors"
	"school21/task/RogueGame/internal/domain/models/actor"
)

type InMemoryActorRepo struct {
	actors []actor.Actor
}

func (repo *InMemoryActorRepo) FindById(id int) (*actor.Actor, error) {
	if id >= len(repo.actors) || id < 0 {
		return nil, errors.New("actor not found")
	}

	return &repo.actors[id], nil

}

func (repo *InMemoryActorRepo) Save(actor *actor.Actor) error {
	// Pass
	return nil
}
