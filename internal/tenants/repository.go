package tenants

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Tenant, error) {
	var tenants []Tenant
	if err := r.db.Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *Repository) GetByGUID(guid string) (*Tenant, error) {
	var tenant Tenant
	if err := r.db.Where("guid = ?", guid).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *Repository) Create(t *Tenant) error {
	return r.db.Create(t).Error
}

func (r *Repository) Update(guid string, t *Tenant) error {
	return r.db.Model(&Tenant{}).Where("guid = ?", guid).Updates(t).Error
}

func (r *Repository) Delete(guid string) error {
	return r.db.Where("guid = ?", guid).Delete(&Tenant{}).Error
}
