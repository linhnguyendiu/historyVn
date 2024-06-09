package app

import (
	"fmt"
	"go-pzn-restful-api/model/domain"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		&domain.Lesson{},
		&domain.LessonContent{},
		&domain.Course{},
		&domain.Chapter{},
		&domain.Category{},
		&domain.CategoryCourse{},
		&domain.Transaction{},
		&domain.Post{},
		&domain.Comment{},
		&domain.Question{},
		&domain.Option{},
		&domain.ExamResult{},
	)

	if err != nil {
		return err
	}

	log.Println("Migration is successfully")
	return nil
}
