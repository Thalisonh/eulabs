package models

import (
	"time"
)

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       float32
	Active      bool
	CreatedAt   time.Time
	UpdateAt    *time.Time
}
