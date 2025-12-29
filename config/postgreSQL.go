package config

import (
	"os"

	"github.com/WilliamOut/gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgreSQL")

	dsn := os.Getenv("DB_SOURCE")

	if dsn == "" {
		dsn = "postgresql://user:password@db:5432/gopportunities?sslmode=disable"
	}

	// Check if the database exists
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar no banco PostgreSQL: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Erro ao migrar o banco PostgreSQL: %v", err)
		return nil, err
	}

	return db, nil
}
