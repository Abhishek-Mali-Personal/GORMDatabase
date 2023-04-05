package Connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (DBString ConnectDBString) ConnectToPostgresDB() (DB *gorm.DB, ConnectionError error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DBString.DBHost, DBString.DBUser, DBString.DBPassword, DBString.DBName, DBString.DBPort)
	DB, ConnectionError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}
