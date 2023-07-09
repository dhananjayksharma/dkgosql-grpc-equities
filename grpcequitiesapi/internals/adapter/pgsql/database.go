package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grpcequitiesapi/internals/adapter/pgsql/entities"
	"log"
)

type MySQLDbStore struct {
	DB *gorm.DB
}

func DBConn(dbsconn string) (*gorm.DB, error) {

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

	return db, nil
}

func GetDbConnect(c *gin.Context) {
	c.JSON(200, gin.H{"message": "In GetDbConnect test"})
}
