package usecases

import (
	"errors"
	"school21/task/RogueGame/internal/domain/models"
	"school21/task/RogueGame/internal/domain/models/actor"
	"testing"
)

// MockActorRepository реализует IActorRepository для тестов
type MockActorRepository struct {
	FindByIdCalledWith int
	SaveCalledWith     *actor.Actor
	FindByIdResult     *actor.Actor
	FindByIdError      error
	SaveError          error
}

func (m *MockActorRepository) FindById(id int) (*actor.Actor, error) {
	m.FindByIdCalledWith = id
	return m.FindByIdResult, m.FindByIdError
}

func (m *MockActorRepository) Save(actor *actor.Actor) error {
	m.SaveCalledWith = actor
	return m.SaveError
}

func TestMoveActorUseCase_Execute_Success(t *testing.T) {
	actor := &actor.Actor{
		ID:       1,
		Position: models.Position2d{X: 0, Y: 0},
	}

	repo := &MockActorRepository{
		FindByIdResult: actor,
	}

	useCase := NewMoveActorUseCase(repo)
	newPos := models.Position2d{X: 5, Y: 5}

	// Act
	useCase.Execute(1, newPos)

	// Assert
	// Проверка вызова FindById
	if repo.FindByIdCalledWith != 1 {
		t.Errorf("FindById called with %d, want 1", repo.FindByIdCalledWith)
	}

	// Проверка обновления позиции
	if actor.Position != newPos {
		t.Errorf("Actor position = %v, want %v", actor.Position, newPos)
	}

	// Проверка вызова Save
	if repo.SaveCalledWith != actor {
		t.Errorf("Save called with %v, want %v", repo.SaveCalledWith, actor)
	}
}

func TestMoveActorUseCase_Execute_FindByIdError(t *testing.T) {
	repo := &MockActorRepository{
		FindByIdError: errors.New("actor not found"),
	}

	useCase := NewMoveActorUseCase(repo)

	// Act
	useCase.Execute(1, models.Position2d{X: 5, Y: 5})

	// Assert
	// Проверка вызова FindById
	if repo.FindByIdCalledWith != 1 {
		t.Errorf("FindById called with %d, want 1", repo.FindByIdCalledWith)
	}

	// Проверка что Save не вызывался
	if repo.SaveCalledWith != nil {
		t.Error("Save should not be called when FindById fails")
	}
}

func TestMoveActorUseCase_Execute_SaveError(t *testing.T) {
	actor := &actor.Actor{ID: 1}
	repo := &MockActorRepository{
		FindByIdResult: actor,
		SaveError:      errors.New("save failed"),
	}

	useCase := NewMoveActorUseCase(repo)

	// Act
	useCase.Execute(1, models.Position2d{X: 5, Y: 5})

	// Assert
	// Проверка что Save был вызван несмотря на ошибку
	if repo.SaveCalledWith == nil {
		t.Error("Save should be called even if it returns error")
	}
}
