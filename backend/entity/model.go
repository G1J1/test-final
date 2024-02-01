package entity

import(
	"gorm.io/gorm"
)

type Model struct{
	gorm.Model

	Student string `valid:"required~studend is required, matches(^[ASD]\\d{2}$)~wrong format"`
	FirstName string `valid:"required~FirstName is required, maxstringlength(10)~not over 10 char"`
	Lastname string `valid:"required~Lastname is required, minstringlength(5)~ at least 5 char"`
	Email string `valid:"required~Email is required , email~ email worng format"`
	Phone string `valid:"required~phone is required , stringlength(10|10)~ phone is 10 num bro"`
	Photo string `valid:"required~photo is required , url~ photo wrong format"`

	GenderID uint `valid:"required~ Gender is required"`
	Gender Gender `gorm:"foreignKey:GenderID"`


}

type Gender struct{
	gorm.Model
	Name string
}