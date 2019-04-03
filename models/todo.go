package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Todo struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50)"`
	Description      string `gorm:"type:varchar(250)"`
}

func (model *Todo) BeforeUpdate() (err error) {

	model.CreatedAt = time.Now()
	return
}
