package usecase

import "clean-architecture-example/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindAll() ([]*domain.User, error)
}

type AddUserUseCase struct {
	repo UserRepository
}

func NewAddUserUseCase(repo UserRepository) *AddUserUseCase {
	return &AddUserUseCase{repo: repo}
}

func (u *AddUserUseCase) Execute(user *domain.User) error {
	// Business logic (e.g., validation) can be added here
	return u.repo.Save(user)
}
