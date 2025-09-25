package app

import (
	"mvc/controllers/donors"
	"mvc/controllers/items"
	"mvc/controllers/tenants"
	"mvc/controllers/users"
)

func mapUrls() {
	router.GET("/donors/:donorId", donors.GetDonor)
	router.GET("/donors", donors.SearchDonor)
	router.POST("/donors", donors.CreateDonor)
	router.GET("iam/v1/tenants", items.TestItems)
	router.GET("iam/v1/tenants/:tenantId", items.GetTenant)
	router.PUT("iam/v1/tenants/:tenantId", items.PutTenant)
	router.GET("iam/v1/users", users.GetUsers)
	router.GET("iam/v1/users/:userId", users.GetUser)
	router.PUT("iam/v1/users/:userId", users.PutUser)
	router.GET("/tenants", tenants.GetTenants)
	router.GET("/tenants/:id", tenants.GetTenant)
	router.POST("/tenants", tenants.CreateTenant)
	router.PUT("/tenants/:id", tenants.UpdateUser)
	router.DELETE("/tenants/:id", tenants.DeleteUser)
}
