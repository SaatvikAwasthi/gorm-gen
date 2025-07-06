package dao

import (
	"time"
)

// User represents a user model. This struct is used as an input for gorm.io/gen
// to define the schema from which code will be generated.
type User struct {
	ID        uint `gorm:"primary"`
	Name      string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []Product // One-to-many relationship with Product
}
