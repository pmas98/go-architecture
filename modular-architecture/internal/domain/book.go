package domain

import (
	"errors"
	"fmt"
	"time"
)

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Author      string    `gorm:"type:varchar(255);not null"`
	PublishedAt time.Time `gorm:"type:date"`
	IsAvailable bool      `gorm:"default:true"`

	RentedByID *uint `gorm:"index"`
	RentedBy   *User `gorm:"foreignKey:RentedByID"`
}

func (b *Book) Rent(userID uint) error {
	fmt.Println(b.IsAvailable)
	if !b.IsAvailable {
		return errors.New("book is not available")
	}
	b.IsAvailable = false
	b.RentedByID = &userID
	return nil
}

func (b *Book) Return() {
	b.IsAvailable = true
	b.RentedByID = nil
	b.RentedBy = nil
}
