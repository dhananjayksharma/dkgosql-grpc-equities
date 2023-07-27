package pgsql

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// "github.com/spf13/viper"
	"grpcequitiesapi/internals/adapter/pgsql/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MySQLDbStore struct {
	DB *gorm.DB
}

type DBConnector interface {
	DBConn(dbsconn string) (*gorm.DB, error)
}

func NewDbConnector(dbsconn string) (*MySQLDbStore, error) {
	var connector DBConnector = &dbConnectorImpl{}
	db, err := connector.DBConn(dbsconn)
	if err != nil {
		log.Fatalln(err)
	}
	return &MySQLDbStore{DB: db}, nil
}

type dbConnectorImpl struct{}

func (dc *dbConnectorImpl) DBConn(dbsconn string) (*gorm.DB, error) {
	fmt.Println("dbsconn :", dbsconn)
	dbURL := dbsconn

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.Merchant{})
	db.AutoMigrate(&entities.OrdersProcessed{})
	db.AutoMigrate(&entities.Users{})
	db.AutoMigrate(&entities.Orders{})
	db.AutoMigrate(&entities.Company{})

	return db, nil
}

func GetDbConnect(c *gin.Context) {
	c.JSON(200, gin.H{"message": "In GetDbConnect test"})
}
