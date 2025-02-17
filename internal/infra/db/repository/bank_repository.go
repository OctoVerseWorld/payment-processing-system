package repository

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/domain"
	"PaymentProcessingSystem/internal/infra/db/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.9.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

const otelName = "internal/infra/db/repository/bank_repository.go"

// BankRepository represents the repository used for interacting with Bank records.
type BankRepository struct {
	q *sqlc.Queries
}

// NewBankRepository instantiates the BankRepository.
func NewBankRepository(d sqlc.DBTX) *BankRepository {
	return &BankRepository{
		q: sqlc.New(d),
	}
}

// Create inserts a new bank record.
func (t *BankRepository) Create(ctx context.Context, params domain.BankCreateParams) (domain.BankID, error) {
	defer newOTELSpan(ctx, "BankRepository.Create").End()

	_, err := t.q.CreateBank(ctx, sqlc.CreateBankParams{
		ID:             int16(params.ID),
		PlanetID:       params.PlanetID,
		OrganizationID: params.OrganizationID,
		Name:           params.Name,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return params.ID, internal.NewErrorf(internal.ErrorCodeInvalidArgument,
					"bank with id %d already exists", params.ID)
			}
		}
		return params.ID, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert bank")
	}

	return params.ID, nil
}

// UpdateName updates the existing record matching the id.
func (t *BankRepository) UpdateName(ctx context.Context, id domain.BankID, name string) error {
	defer newOTELSpan(ctx, "BankRepository.UpdateName").End()

	err := t.q.UpdateBankName(ctx, sqlc.UpdateBankNameParams{
		Name: name,
		ID:   int16(id),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "bank %d not found", id)
		}
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "update bank name")
	}

	return nil
}

// Delete deletes the existing record matching the id.
func (t *BankRepository) Delete(ctx context.Context, id domain.BankID) error {
	defer newOTELSpan(ctx, "BankRepository.Delete").End()

	err := t.q.DeleteBank(ctx, int16(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "bank not found")
		}
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "delete bank")
	}

	return nil
}

// GetOneByID retrieves the existing record matching the id.
func (t *BankRepository) GetOneByID(ctx context.Context, id domain.BankID) (domain.Bank, error) {
	defer newOTELSpan(ctx, "BankRepository.GetByID").End()

	bank, err := t.q.SelectBankByID(ctx, int16(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Bank{}, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "bank not found")
		}
		return domain.Bank{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "get bank")
	}

	return domain.Bank{
		ID:             domain.BankID(bank.ID),
		PlanetID:       bank.PlanetID,
		OrganizationID: bank.OrganizationID,
		Name:           bank.Name,
	}, nil
}

// GetAllByPlanetID retrieves all bank records by PlanetID.
func (t *BankRepository) GetAllByPlanetID(ctx context.Context, planetId int32) ([]domain.Bank, error) {
	defer newOTELSpan(ctx, "BankRepository.GetAllByPlanetID").End()

	banks, err := t.q.SelectBanksByPlanetID(ctx, planetId)
	if err != nil {
		zap.L().Error("get all banks", zap.Error(err))
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "no banks found for planet: %d", planetId)
		}
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "get all banks")
	}

	var result []domain.Bank
	for _, bank := range banks {
		result = append(result, domain.Bank{
			ID:             domain.BankID(bank.ID),
			PlanetID:       bank.PlanetID,
			OrganizationID: bank.OrganizationID,
			Name:           bank.Name,
		})
	}

	return result, nil
}

func newOTELSpan(ctx context.Context, name string) trace.Span {
	_, span := otel.Tracer(otelName).Start(ctx, name)

	span.SetAttributes(semconv.DBSystemPostgreSQL)

	return span
}
