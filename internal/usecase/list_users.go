package usecase

import "clean-architecture-example/internal/domain"

type ListUsersUseCase struct {
    repo UserRepository
}

func NewListUsersUseCase(repo UserRepository) *ListUsersUseCase {
    return &ListUsersUseCase{repo: repo}
}

func (u *ListUsersUseCase) Execute() ([]*domain.User, error) {
    return u.repo.FindAll()
}