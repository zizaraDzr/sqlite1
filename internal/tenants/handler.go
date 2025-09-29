package tenants

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("iam/v1/tenants", h.GetTenants)
	r.GET("iam/v1/tenants/:guid", h.GetTenant)
	r.POST("iam/v1/tenants", h.CreateTenant)
	r.PUT("iam/v1/tenants/:guid", h.UpdateTenant)
	r.DELETE("iam/v1/tenants/:guid", h.DeleteTenant)
}

func (h *Handler) GetTenants(c *gin.Context) {
	tenants, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tenants)
}

func (h *Handler) GetTenant(c *gin.Context) {
	guid := c.Param("guid")
	tenant, err := h.service.GetByGUID(guid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}
	c.JSON(http.StatusOK, tenant)
}

func (h *Handler) CreateTenant(c *gin.Context) {
	var tenant Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tenant)
}

func (h *Handler) UpdateTenant(c *gin.Context) {
	guid := c.Param("guid")
	var tenant Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Update(guid, &tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tenant)
}

func (h *Handler) DeleteTenant(c *gin.Context) {
	guid := c.Param("guid")
	if err := h.service.Delete(guid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
