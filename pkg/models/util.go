package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

var db *gorm.DB

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
   }