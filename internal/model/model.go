package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type Source struct {
	gorm.Model
	Name    string
	Comment string
	Ips     []Ip `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Ip struct {
	gorm.Model
	Ip       string
	Port     string
	SourceID uint
}

type Command struct {
	gorm.Model
	Name           string
	Path           string
	Value          string
	MessageLogs    string
	DefaultMessage string
}

type Log struct {
	gorm.Model
	Value string
}

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Source{}, &Ip{}, &Command{}, &Log{})

	return db
}
