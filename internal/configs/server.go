package configs

import (
	"PaymentProcessingSystem/internal"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"os"
	"strconv"
)

const serverConfigPrefix = "SERVER_"

// NewServerConfig ...
func NewServerConfig() (*ServerConfig, error) {
	intPort, err := strconv.Atoi(os.Getenv(serverConfigPrefix + "PORT"))
	if err != nil {
		return &ServerConfig{}, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, serverConfigPrefix+"PORT is not a valid number")
	}
	return &ServerConfig{
		Host: os.Getenv(serverConfigPrefix + "HOST"),
		Port: intPort,
	}, nil
}

// ServerConfig ...
type ServerConfig struct {
	Host string
	Port int
}

// GetAddress ...
func (s *ServerConfig) GetAddress() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

// Validate ...
func (s *ServerConfig) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Host, validation.Required),
		validation.Field(&s.Port, validation.Required, validation.Min(1)),
	)
}
