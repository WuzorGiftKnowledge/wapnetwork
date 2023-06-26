package models

import (
	"time"
	"github.com/jinzhu/gorm"
	// "github.com/google/uuid"
	 "github.com/satori/go.uuid"
	
)

var db *gorm.DB

type Base struct {
	ID        uuid.UUID 	`json:"id"`
	CreatedAt time.Time		`json:"createdAt"`
	UpdatedAt time.Time		`json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`
   }
  // BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid:= uuid.NewV4()
	
	return scope.SetColumn("ID", uuid)
   }

   type CurrentUser struct {
    Id    uuid.UUID
   Username      string
   Authorized bool

}