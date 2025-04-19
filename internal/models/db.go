package models

import "time"

type DbModel struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type DbModelWithID struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
