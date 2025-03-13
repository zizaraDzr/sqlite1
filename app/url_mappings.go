package app

import (
	controllers "mvc/controllers/users"
)

func mapUrls() {
	router.GET("/donors/:donorId", controllers.GetDonor)
	router.GET("/donors", controllers.SearchDonor)
	router.POST("/donors", controllers.CreateDonor)
}
