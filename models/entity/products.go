package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" valid:"int"`
	UUID      string     `gorm:"not null" valid:"uuid"`
	Name      string     `gorm:"not null" json:"name" valid:"required"`
	ImageURL  string     `gorm:"not null" json:"image_url" valid:"url"`
	AdminID   uint       `gorm:"not null" json:"admin_id" valid:"int"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Variants  []Variant  `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

func (p *Product) BeforeDelete(tx *gorm.DB) (err error) {
	// get product info
	_ = tx.Debug().Where("uuid = ?", p.UUID).First(&p)

	err = tx.Where("product_id = ?", p.ID).Delete(&Variant{}).Error
	if err != nil {
		return err
	}
	return nil
}
