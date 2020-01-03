package models

import (
	"time"

	"github.com/MohabMohamed/try-gin/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

type Todo struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Description string
	CreatedAt   time.Time
	Done        bool
}

func init() {

	var err error
	DB, err = gorm.Open("sqlite3", config.DbURL)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	DB.AutoMigrate(&Todo{})
}
