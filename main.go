package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/sqlite"
	"runtime/debug"
)

type User struct {
	ID uint64 `gorm:'primaryKey'`
	Name string
	Tweets []Tweet
}

type Tweet struct {
	ID uint64 `gorm:'primaryKey'`
	UserID uint64
	Content string
}

func main() {
	debug.SetGCPercent(1)

	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	sqlite3, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db, err := sqlite3.DB()
	defer db.Close()
	if err != nil {
		db.Close()
	}

	r.GET("/", func(c *gin.Context) {
		var users []User
		sqlite3.Preload("Tweets").Find(&users)
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}
