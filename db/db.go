package db

import (
	"fmt"
	"github.com/agusalex/golang-employee-crud/config"
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB *gorm.DB
)

func Connect() error {
	var err error
	c := config.Get()
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Name)
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatal("An error occured while trying to connect to your db: ", err)
		return err
	}
	Migrate(db)
	DB = db
	return err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Member{}, &models.Tag{})
}
