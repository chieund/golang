package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	NickName  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;unique" json:"password"`
	CreatedAt time.Time `gorm:"default:null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:null" json:"updated_at"`
}

func CreateUser(db *gorm.DB, user *User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUsers(db *gorm.DB, user *[]User) (err error) {
	err = db.Find(user).Error

	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, user *User, id uint64) (err error) {
	err = db.Where("id=?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, user *User) (err error) {
	db.Save(user)
	return nil
}

func DeleteUser(db *gorm.DB, user *User, id uint64) (err error) {
	db.Where("id=?", id).Delete(user)
	return nil
}
