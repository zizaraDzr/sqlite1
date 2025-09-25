package items

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Organization struct {
	GUID       string `json:"guid"`
	ID         int64  `json:"id"`
	IsBlocked  bool   `json:"is_blocked"`
	IsDeleted  bool   `json:"is_deleted"`
	IsPlatform bool   `json:"is_platform"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Website    string `json:"website"`
	Code       string `json:"code"`
	SsoId      string `json:"sso_id"`
}

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func generateItems() []Organization {

	organizations := []Organization{}
	for i := 0; i < 100; i++ {
		organizations = append(organizations, generateTenant())
		organizations[i].ID = int64(i + 1)
	}
	return organizations
}
func generateTenant() Organization {
	organizations := Organization{
		GUID:       "550e8400-e29b-41d4-a716-446655440000",
		ID:         1,
		IsBlocked:  true,
		IsDeleted:  true,
		IsPlatform: true,
		Name:       RandString(15),
		Properties: RandString(15),
		Website:    "https://example.com",
	}
	return organizations
}
func putGenerateTenant() Organization {
	organizations := Organization{
		GUID:       "550e8400-e29b-41d4-a716-446655440000",
		ID:         1,
		IsBlocked:  true,
		IsDeleted:  true,
		IsPlatform: true,
		Name:       "123g",
		Properties: "333es",
		Website:    "https://example.com",
	}
	return organizations
}
func TestItems(c *gin.Context) {
	result := generateItems()

	c.JSON(http.StatusOK, result)
}
func GetTenant(c *gin.Context) {
	result := generateTenant()

	c.JSON(http.StatusOK, result)
}
func PutTenant(c *gin.Context) {
	result := putGenerateTenant()
	c.JSON(http.StatusOK, result)
}
