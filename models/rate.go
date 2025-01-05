package models

import "time"

type Rate struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	Code    string    `gorm:"size:10;not null" json:"code"`
	Name    string    `gorm:"size:100;not null" json:"name"`
	Nominal int       `gorm:"not null" json:"nominal"`
	Rate    float64   `gorm:"type:decimal(10,4);not null" json:"rate"`
	Date    time.Time `gorm:"not null" json:"date"`
}
