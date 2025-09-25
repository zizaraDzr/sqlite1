package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Настройка CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},                                       // Замените на список разрешенных источников
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},            // Укажите методы, которые вы хотите разрешить
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}, // Укажите заголовки, которые вы хотите разрешить
	}))
	mapUrls()
	router.Run(":8000")
}
