package tenants

import (
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

func GetByID(id string) (Tenant, error) {
	var user Tenant
	result := db.First(&user, id)
	return user, result.Error
}

func Create(tenant *Tenant) error {
	result := db.Create(tenant)
	return result.Error
}

func Update(id string, user *Tenant) error {
	var existing Tenant
	if err := db.First(&existing, id).Error; err != nil {
		return err
	}
	existing.Name = user.Name
	return db.Save(&existing).Error
}

func Delete(id string) error {
	return db.Delete(&Tenant{}, id).Error
}
