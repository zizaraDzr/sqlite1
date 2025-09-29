package tenants

import (
	"mvc/domain/tenants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTenants(c *gin.Context) {
	result, err := tenants.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetTenant(c *gin.Context) {
	guid := c.Param("guid")
	result, err := tenants.GetByGuid(guid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateTenant(c *gin.Context) {
	var tenant tenants.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tenants.Create(&tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tenant)
}

func UpdateTenant(c *gin.Context) {
	guid := c.Param("guid")
	var tenant tenants.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tenants.Update(guid, &tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tenant)
}

func DeleteTenant(c *gin.Context) {
	guid := c.Param("guid")
	println(guid)
	if err := tenants.Delete(guid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
