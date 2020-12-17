package clerk

import (
	"cms/app/models"
	"database/sql"
)

type Clerk struct {
	//pk
	models.PrimaryKeyField

	Username string `json:"username" gorm:"type:varchar(50);not null;default:"""`
	Password string `json:"password" gorm:"type:char(32);not null;default ""`
	Encrypt string `json:"encrypt" gorm:"type:char(10);not null;default """`
	Email string `json:"email" gorm:"type:varchar(50);not null;default """`
	IsVerifyEmail uint `json:"is_verify_email" gorm:"size:8;not null;default:0"`
	IsActive uint `json:"is_active" gorm:"size:8;not null default:0"`
	IsResetPassword uint `json:"is_reset_password" gorm:"size:8;not null;default:0"`
	Intro string `json:"intro" gorm:"type:varchar(50);not null;default:"";after:is_reset_password"`
	LastLoginAt sql.NullTime `json:"last_login_at"`

	models.TimestampField
}
