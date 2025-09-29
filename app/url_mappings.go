package app

import (
	"mvc/controllers/donors"
	"mvc/controllers/tenants"
	"mvc/controllers/users"
)

func mapUrls() {
	router.GET("/donors/:donorId", donors.GetDonor)
	router.GET("/donors", donors.SearchDonor)
	router.POST("/donors", donors.CreateDonor)
	router.GET("iam/v1/users", users.GetUsers)
	router.GET("iam/v1/users/:userId", users.GetUser)
	router.PUT("iam/v1/users/:userId", users.PutUser)
	router.GET("iam/v1/tenants", tenants.GetTenants)
	router.GET("iam/v1/tenants/:guid", tenants.GetTenant)
	router.POST("iam/v1/tenants", tenants.CreateTenant)
	router.PUT("iam/v1/tenants/:guid", tenants.UpdateTenant)
	router.DELETE("iam/v1/tenants/:guid", tenants.DeleteTenant)
}
