package app

import (
	"mvc/controllers/donors"
)

func mapUrls() {
	router.GET("/donors/:donorId", donors.GetDonor)
	router.GET("/donors", donors.SearchDonor)
	router.POST("/donors", donors.CreateDonor)
}
