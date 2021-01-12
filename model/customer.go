package model

import(
	"time"
	"gorm.io/gorm"
)

type Customer struct {
	Id int `json:"id" gorm:"autoIncrement;primaryKey"`
	Name string `json:"name"`
	Address string `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
