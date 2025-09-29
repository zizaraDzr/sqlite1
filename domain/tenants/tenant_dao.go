package tenants

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=appdb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// миграция
	db.AutoMigrate(&Tenant{})
}

func GetAll() ([]Tenant, error) {
	var users []Tenant
	result := db.Find(&users)
	return users, result.Error
}

func GetByGuid(guid string) (Tenant, error) {
	var user Tenant
	result := db.Where("guid = ?", guid).First(&user)
	return user, result.Error
}

func Create(tenant *Tenant) error {
	result := db.Create(tenant)
	return result.Error
}

func Update(guid string, user *Tenant) error {
	var existing Tenant
	if err := db.Where("guid = ?", guid).First(&existing).Error; err != nil {
		return err
	}
	existing.Name = user.Name
	existing.Website = user.Website
	existing.Properties = user.Properties
	existing.IsBlocked = user.IsBlocked
	existing.IsDeleted = user.IsDeleted
	existing.IsPlatform = user.IsPlatform
	fmt.Println(existing)
	return db.Save(&existing).Error
}

func Delete(guid string) error {
	return db.Where("guid = ?", guid).Delete(&Tenant{}).Error
}
