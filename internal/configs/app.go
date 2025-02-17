package configs

import (
	"PaymentProcessingSystem/internal"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
	"os"
	"strconv"
)

const appConfigPrefix = "APP_"

// NewAppConfig ...
func NewAppConfig() (*AppConfig, error) {
	var defaultUserID int
	var err error
	defaultUserIDStr := os.Getenv(appConfigPrefix + "DEFAULT_USER_ID")
	if !(defaultUserIDStr == "") {
		defaultUserID, err = strconv.Atoi(defaultUserIDStr)
		if err != nil {
			return nil, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "strconv.Atoi")
		}
	} else {
		zap.L().Info("Default user ID not set, using default value", zap.Int("defaultUserID", 1))
		defaultUserID = 1
	}

	return &AppConfig{
		Environment:   EnvironmentType(os.Getenv(appConfigPrefix + "ENVIRONMENT")),
		DefaultUserID: defaultUserID,
	}, nil
}

// AppConfig ...
type AppConfig struct {
	Environment   EnvironmentType
	DefaultUserID int
}

// Validate ...
func (s *AppConfig) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Environment, validation.Required),
		validation.Field(&s.DefaultUserID, validation.Required),
	)
}

// EnvironmentType ...
type EnvironmentType string

const (
	EnvDevelopment EnvironmentType = "development"
	EnvProduction  EnvironmentType = "production"
)

// Validate ...
func (e *EnvironmentType) Validate() error {
	switch *e {
	case EnvDevelopment, EnvProduction:
		return nil
	}

	return validation.NewError(
		"configs.App.Environment",
		fmt.Sprintf("Environment is invalid. Must be one of: %s or %s", EnvDevelopment, EnvProduction),
	)
}

// IsProduction ...
func (e *EnvironmentType) IsProduction() bool {
	return *e == EnvProduction
}

// IsDevelopment ...
func (e *EnvironmentType) IsDevelopment() bool {
	return *e == EnvDevelopment
}
