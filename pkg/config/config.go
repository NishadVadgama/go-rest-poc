package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Environment string
	DB          DBConfig
}

type DBConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBname          string
	DBSchema        string
	CertificatePath string
}

var config Config

func init() {
	config = Config{
		Port:        getEnvVariable("PORT", "8080"),
		Environment: getEnvVariable("ENVIRONMENT", "local"),
		DB: DBConfig{
			Host:            getEnvVariable("RDM_RDS_URL", "localhost"),
			Port:            getEnvVariable("RDM_RDS_PORT", "5432"),
			User:            getEnvVariable("RDM_RDS_USERNAME", "postgres"),
			Password:        getEnvVariable("RDM_RDS_PASSWORD", ""),
			DBname:          getEnvVariable("RDM_RDS_DATABASE", "postgres"),
			DBSchema:        getEnvVariable("RDM_RDS_SCHEMA", ""),
			CertificatePath: getEnvVariable("RDM_RDS_CERTIFICATE_PATH", ""),
		},
	}
}

func GetPort() string {
	return config.Port
}

func GetEnvironment() string {
	return config.Environment
}

func GetPGConnectionString() string {
	if config.DB.Host != "localhost" {
		return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslrootcert=%s sslmode=verify-full", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DBname, config.DB.CertificatePath)
	}
	return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DBname)
}

func getEnvVariable(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
