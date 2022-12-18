package mysql

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDbStore struct {
	DB *gorm.DB
}

func DBConn(dbsconn string) (*gorm.DB, error) {

	fmt.Println("dbsconn :", dbsconn)

	var (
		databasename, hostPort, userPass, userName string
	)

	if dbsconn == "db_one_readwrite" || true {
		userName = viper.GetString("database.db_one_readwrite.dbuser")
		userPass = viper.GetString("database.db_one_readwrite.dbpassword")
		hostPort = viper.GetString("database.db_one_readwrite.hostname")
		databasename = viper.GetString("database.db_one_readwrite.dbname")
	} else if dbsconn == "db_one_read" {
		userName = viper.GetString("database.db_one_read.dbuser")
		userPass = viper.GetString("database.db_one_read.dbpassword")
		hostPort = viper.GetString("database.db_one_read.hostname")
		databasename = viper.GetString("database.db_one_read.dbname")
	} else if dbsconn == "db_one_write" {
		userName = viper.GetString("database.db_one_write.dbuser")
		userPass = viper.GetString("database.db_one_write.dbpassword")
		hostPort = viper.GetString("database.db_one_write.hostname")
		databasename = viper.GetString("database.db_one_write.dbname")
	}

	var DB_URI = userName + ":" + userPass + "@tcp(" + hostPort + ")/" + databasename + "?parseTime=true&loc=Local"
	//dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("DB_URI:", DB_URI)
	// dataConfig := &gorm.Config{}
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       strings.Trim(DB_URI, "'"),
	// 	DefaultStringSize:         256,
	// 	DisableDatetimePrecision:  true,
	// 	DontSupportRenameIndex:    true,
	// 	DontSupportRenameColumn:   true,
	// 	SkipInitializeWithVersion: false,
	// }), dataConfig)
	db, err := gorm.Open(mysql.Open(DB_URI), &gorm.Config{})
	if err != nil {
		return db, err
	}

	rawSqlForSetDababase := "USE " + databasename + ";"
	db.Exec(rawSqlForSetDababase)
	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func GetDbConnect(c *gin.Context) {
	c.JSON(200, gin.H{"message": "In GetDbConnect test"})
}
