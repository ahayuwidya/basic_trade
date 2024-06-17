package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" valid:"int"`
	UUID        string     `gorm:"not null" valid:"uuid"`
	VariantName string     `gorm:"not null" json:"variant_name" form:"variant_name" valid:"required"`
	Quantity    uint       `gorm:"not null" json:"quantity" form:"quantity" valid:"int"`
	ProductID   uint       `gorm:"not null" json:"product_id" form:"product_id" valid:"int"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(v)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
