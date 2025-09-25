package tenants

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	GUID       string `json:"guid" gorm:"type:uuid"`
	ID         int64  `json:"id" gorm:"primaryKey"`
	IsBlocked  bool   `json:"is_blocked"`
	IsDeleted  bool   `json:"is_deleted"`
	IsPlatform bool   `json:"is_platform"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Website    string `json:"website"`
	Code       string `json:"code"`
	SsoId      string `json:"sso_id"`
}

func (t *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	t.GUID = uuid.New().String()
	return
}
