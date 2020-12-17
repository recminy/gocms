package article

import (
	"cms/app/models"
	"database/sql"
)

type Article struct {
	//ID uint `json:"id" gorm:"size:32;primaryKey"`
	models.PrimaryKeyField

	AuthorId uint `json:"author_id" gorm:"size:32;not null;default:0"`
	CategoryId uint16 `json:"category_id" gorm:"size:32;not null;default:0"`
	Title string `json:"title" gorm:"type:varchar(150);not null;default:"""`
	Content sql.NullString `json:"content" gorm:"type:text;not null;default:"""`
	IsActive uint `json:"is_active" gorm:"size:8;not null;default:0"`
	//gorm.Model

	models.TimestampField
}


