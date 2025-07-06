package dao

import (
	"time"
)

// Product represents a product model. This struct is used as an input for gorm.io/gen
// to define the schema from which code will be generated.
type Product struct {
	ID        uint `gorm:"primary"`
	UserID    uint // Foreign key to User
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
