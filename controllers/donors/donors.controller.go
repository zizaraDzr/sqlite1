package donors

import (
	"mvc/domain/donors"
	services "mvc/services/donors"
	"mvc/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDonor(c *gin.Context) {
	var donor donors.Donor
	if err := c.ShouldBindJSON(&donor); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	// bytes, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// todo: handle error
	// 	fmt.Println("error bytes", err)
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &donor); err != nil {
	// 	// todo: handle error
	// 	fmt.Println("error json Unmarshal", err.Error())
	// 	return
	// }
	result, saveErr := services.CreateDonor(donor)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchDonor(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}

func GetDonor(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}
