package model

import "time"

type Device struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	CreationTime time.Time `json:"creationTime"`
}

type PartialDevice struct {
	Name  *string `json:"name,omitempty" binding:"omitempty"`
	Brand *string `json:"brand,omitempty" binding:"omitempty"`
}
