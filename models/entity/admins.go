package entity

import (
	"basic_trade/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	UUID      string     `gorm:"not null"`
	Name      string     `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	Email     string     `gorm:"not null;unique" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Products  []Product  `gorm:"foreignKey:AdminID"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(a)

	if errCreate != nil {
		err = errCreate
		return
	}

	a.Password = helpers.HashPass(a.Password)

	err = nil
	return
}
