package models

import "time"

type Item struct {
    ID          uint      `gorm:"primaryKey;autoIncrement"`
    Name        string    `gorm:"type:varchar(255);not null"`
    Description string    `gorm:"type:varchar(255)"`
    Quantity    int       `gorm:"not null"`
    OrderID     uint      `gorm:"not null"`
    CreatedAt   *time.Time
    UpdatedAt   *time.Time
}