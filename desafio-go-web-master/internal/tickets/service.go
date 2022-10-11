package tickets

import (
	"context"

	"github.com/fzunda/desafio-goweb-AgustinZunda/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return s.repository.GetTicketByDestination(ctx, destination)
}
