package service

import (
	"PaymentProcessingSystem/internal/domain"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const otelName = "github.com/MarioCarrion/todo-api/internal/service"

// BankRepository ...
type BankRepository interface {
	Create(ctx context.Context, params domain.BankCreateParams) (domain.BankID, error)
	UpdateName(ctx context.Context, id domain.BankID, name string) error
	Delete(ctx context.Context, id domain.BankID) error
	GetOneByID(ctx context.Context, id domain.BankID) (domain.Bank, error)
	GetAllByPlanetID(ctx context.Context, planetID int32) ([]domain.Bank, error)
}

// BankService ...
type BankService struct {
	repo BankRepository
}

// NewBankService ...
func NewBankService(repo BankRepository) *BankService {
	return &BankService{
		repo: repo,
	}
}

// Create ...
func (s *BankService) Create(ctx context.Context, params domain.BankCreateParams) (domain.BankID, error) {
	defer newOTELSpan(ctx, "Task.Update").End()

	return s.repo.Create(ctx, params)
}

// UpdateName ...
func (s *BankService) UpdateName(ctx context.Context, id domain.BankID, name string) error {
	defer newOTELSpan(ctx, "Task.Update").End()

	return s.repo.UpdateName(ctx, id, name)
}

// Delete ...
func (s *BankService) Delete(ctx context.Context, id domain.BankID) error {
	defer newOTELSpan(ctx, "Task.Delete").End()

	return s.repo.Delete(ctx, id)
}

// GetOneByID ...
func (s *BankService) GetOneByID(ctx context.Context, id domain.BankID) (domain.Bank, error) {
	defer newOTELSpan(ctx, "Task.GetOneByID").End()

	return s.repo.GetOneByID(ctx, id)
}

// GetAllByPlanetID ...
func (s *BankService) GetAllByPlanetID(ctx context.Context, planetID int32) ([]domain.Bank, error) {
	defer newOTELSpan(ctx, "Task.GetAllByPlanetID").End()

	return s.repo.GetAllByPlanetID(ctx, planetID)
}

func newOTELSpan(ctx context.Context, name string) trace.Span {
	_, span := otel.Tracer(otelName).Start(ctx, name)

	return span
}
