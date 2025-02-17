package configs

import (
	"PaymentProcessingSystem/internal"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/url"
	"os"
	"strconv"
)

const dbConfigPrefix = "DB_"

// NewDatabaseConfig ...
func NewDatabaseConfig() (*DatabaseConfig, error) {
	port, err := strconv.Atoi(os.Getenv(dbConfigPrefix + "PORT"))
	if err != nil {
		return &DatabaseConfig{}, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, dbConfigPrefix+"PORT is not a valid number")
	}
	return &DatabaseConfig{
		Host: os.Getenv(dbConfigPrefix + "HOST"),
		Port: port,
		User: os.Getenv(dbConfigPrefix + "USER"),
		Pass: os.Getenv(dbConfigPrefix + "PASS"),
		Name: os.Getenv(dbConfigPrefix + "NAME"),
	}, nil
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

// BuildDSN ...
func (c *DatabaseConfig) BuildDSN() *url.URL {
	return &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.User, c.Pass),
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Name,
	}
}

// Validate ...
func (c *DatabaseConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.Port, validation.Required, validation.Min(1000)),
		validation.Field(&c.User, validation.Required),
		validation.Field(&c.Pass, validation.Required),
		validation.Field(&c.Name, validation.Required),
	)
}
