package configs

import (
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var (
	configInstance *Configuration
	once           sync.Once
)

type Configuration struct {
	AppConfig      *AppConfig
	DatabaseConfig *DatabaseConfig
	ServerConfig   *ServerConfig
}

// InitConfig ...
func InitConfig(envFile string) error {
	// Load configuration from dotenv file
	loadConfig(envFile)

	// Initialize configuration
	var err error
	once.Do(func() {
		appConfig, errApp := NewAppConfig()
		if errApp != nil {
			err = errApp
			return
		}

		databaseConfig, errDB := NewDatabaseConfig()
		if errDB != nil {
			err = errDB
			return
		}

		serverConfig, errSrv := NewServerConfig()
		if errSrv != nil {
			err = errSrv
			return
		}

		configInstance = &Configuration{
			AppConfig:      appConfig,
			DatabaseConfig: databaseConfig,
			ServerConfig:   serverConfig,
		}
	})
	return err
}

// GetConfig возвращает уже созданную конфигурацию
func GetConfig() *Configuration {
	if configInstance == nil {
		log.Fatal("Config is not initialized. Call InitConfig() first.")
	}
	return configInstance
}

// LoadConfig ...
func loadConfig(envFile string) {
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("%s file not found: %v", envFile, err)
	}
}
