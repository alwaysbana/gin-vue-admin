package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Username    string         `gorm:"type:varchar(100);not null;uniqueIndex;comment:用户名"`
	Password    string         `gorm:"type:varchar(255);not null;comment:密码"`
	NickName    string         `gorm:"type:varchar(255);not null;comment:昵称"`
	HeaderImg   string         `gorm:"type:varchar(255);not null;comment:头像"`
	AuthorityId uint           `gorm:"type:int(255);not null;comment:角色权限"`
}
