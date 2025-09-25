package users

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	AvatarGuid string `json:"avatarGuid"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	Guid       string `json:"guid"`
	Id         int64  `json:"id"`
	IsAdmin    bool   `json:"isAdmin"`
	IsBlocked  bool   `json:"isBlocked"`
	IsDeleted  bool   `json:"isDeleted"`
	LastName   string `json:"lastName"`
	Login      string `json:"login"`
	MiddleName string `json:"middleName"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Properties string `json:"properties"`
	TenantId   int64  `json:"tenantId"`
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
func generateUsers() []User {
	user := []User{}
	for i := 0; i < 100; i++ {
		user = append(user, generateUser())
		user[i].Id = int64(i + 1)
	}

	return user
}
func generateUser() User {
	user := User{
		AvatarGuid: "550e8400-e29b-41d4-a716-446655440000",
		Email:      "123g",
		FirstName:  RandString(15),
		Guid:       RandString(15),
		Id:         1,
		IsAdmin:    true,
		IsBlocked:  true,
		IsDeleted:  true,
		LastName:   RandString(15),
		Login:      RandString(15),
		MiddleName: RandString(15),
		Password:   RandString(15),
		Phone:      RandString(15),
		Properties: RandString(15),
		TenantId:   1,
	}
	return user
}
func putGeneratUser() User {
	user := User{
		AvatarGuid: "550e8400-e29b-41d4-a716-446655440000",
		Email:      "123g",
		FirstName:  RandString(15),
		Guid:       "550e8400-e29b-41d4-a716-446655440000",
		Id:         1,
		IsAdmin:    true,
		IsBlocked:  true,
		IsDeleted:  true,
		LastName:   RandString(15),
		Login:      RandString(15),
		MiddleName: RandString(15),
		Password:   RandString(15),
		Phone:      RandString(15),
		Properties: RandString(15),
		TenantId:   1,
	}
	return user
}
func GetUsers(c *gin.Context) {
	result := generateUsers()

	c.JSON(http.StatusOK, result)
}
func GetUser(c *gin.Context) {
	result := generateUser()

	c.JSON(http.StatusOK, result)
}
func PutUser(c *gin.Context) {
	result := putGeneratUser()
	c.JSON(http.StatusOK, result)
}
