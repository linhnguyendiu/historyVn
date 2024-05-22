package app

import (
	"fmt"
	"go-pzn-restful-api/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DBConnection() *gorm.DB {

	dbGorm, err := gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	} else {
		fmt.Println("Connect to database sucessfull")
	}

	return dbGorm
}

func DBMigrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(
		&domain.User{},
		&domain.Author{},
		&domain.LessonTitle{},
		&domain.LessonContent{},
		&domain.Course{},
		&domain.Category{},
		&domain.CategoryCourse{},
		&domain.Transaction{},
	)

	if err != nil {
		return err
	}

	log.Println("Migration is successfully")
	return nil
}
