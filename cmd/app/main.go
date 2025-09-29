package main

import (
	"mvc/internal/tenants"
	"mvc/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	db.Init()

	// миграции
	db.DB.AutoMigrate(&tenants.Tenant{})

	// init modules
	tenantRepo := tenants.NewRepository(db.DB)
	tenantService := tenants.NewService(tenantRepo)
	tenantHandler := tenants.NewHandler(tenantService)

	// setup gin
	r := gin.Default()
	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},                                       // Замените на список разрешенных источников
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},            // Укажите методы, которые вы хотите разрешить
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}, // Укажите заголовки, которые вы хотите разрешить
	}))
	// register routes
	tenantHandler.RegisterRoutes(r)

	// run server
	r.Run(":8000")
}
