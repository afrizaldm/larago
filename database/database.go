package database

import (
	"fmt"
	"simple-api/config"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	d := &Database{}

	d.Connect()

	return d
}

func (d *Database) Connect() *Database {
	DBConfig := config.NewDatabaseConfig().Load()

	var db *gorm.DB
	var err error

	switch DBConfig.DB_CONNECTION {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect to SQLite: %v", err))
		}

	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DBConfig.DB_USERNAME,
			DBConfig.DB_PASSWORD,
			DBConfig.DB_HOST,
			DBConfig.DB_PORT,
			DBConfig.DB_DATABASE,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect to MySQL: %v", err))
		}

	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			DBConfig.DB_HOST,
			DBConfig.DB_USERNAME,
			DBConfig.DB_PASSWORD,
			DBConfig.DB_DATABASE,
			DBConfig.DB_PORT,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect to PostgreSQL: %v", err))
		}

	default:
		panic("unknown database driver specified in configuration")
	}

	d.DB = db

	return d

}
