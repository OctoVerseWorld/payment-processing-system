package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// BankID ...
type BankID int32

//var _ validation.Validatable = (*BankID)(nil)

// Validate ...
func (bid BankID) Validate() error {
	if bid < 1000 || bid >= 10000 {
		return validation.NewError("BankID", "must be greater than 999 and less than 10000")
	}

	return nil
}

//-

// Bank ...
type Bank struct {
	ID             BankID
	PlanetID       int32
	OrganizationID int32
	Name           string
}

//-

// BankCreateParams ...
type BankCreateParams struct {
	ID             BankID
	PlanetID       int32
	OrganizationID int32
	Name           string
}

// Validate ...
func (b BankCreateParams) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.ID, validation.Required),
		validation.Field(&b.PlanetID, validation.Required),
		validation.Field(&b.OrganizationID, validation.Required),
		validation.Field(&b.Name, validation.Required),
	)
}
