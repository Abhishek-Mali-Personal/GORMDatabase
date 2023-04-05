package Connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once

var instance *gorm.DB

func (DBString ConnectDBString) SingletonConnectToPostgresDB() (*gorm.DB, error) {
	var initializeDBError error
	once.Do(func() {
		instance, initializeDBError = DBString.initializeDB()
	})
	return instance, initializeDBError
}

func (DBString ConnectDBString) CloseSingletonDatabase() error {
	database, connectDBError := DBString.SingletonConnectToPostgresDB()
	if connectDBError != nil {
		return connectDBError
	}
	db, getDBError := database.DB()
	if getDBError != nil {
		return getDBError
	}
	return db.Close()
}

func (DBString ConnectDBString) initializeDB() (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", DBString.DBHost, DBString.DBUser, DBString.DBPassword, DBString.DBName, DBString.DBPort)
	db, OpenPostgresDBError := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}))
	return db, OpenPostgresDBError
}
