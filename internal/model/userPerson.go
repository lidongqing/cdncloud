package model

import "time"

type UserPerson struct {
	Id         int64     `gorm:"primary_key;auto_increment"`
	Name       string    `gorm:"type:varchar(50);charset:utf8mb4;collation:utf8mb4_unicode_ci;default:'';comment:'真实姓名'"`
	UserID     int64     `gorm:"not null;default:0;comment:'用户'"`
	Card       string    `gorm:"type:varchar(50);charset:utf8mb4;collation:utf8mb4_unicode_ci;default:null;comment:'真实身份证'"`
	Image      string    `gorm:"type:varchar(255);charset:utf8mb4;collation:utf8mb4_unicode_ci;not null;default:''"`
	Status     int64     `gorm:"not null;default:0;comment:'0 待认证 1 已认证中'"`
	Mobile     string    `gorm:"type:varchar(15);charset:utf8mb4;collation:utf8mb4_unicode_ci;default:null;comment:'手机号码'"`
	MobilePre  string    `gorm:"type:varchar(11);charset:utf8mb4;collation:utf8mb4_unicode_ci;default:null"`
	Createdate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (UserPerson) TableName() string {
	return "se_user_person"
}

const (
	// 用户认证状态
	USER_PERSON_STATUS_WAIT   = 0 // 待认证
	USER_PERSON_STATUS_NORMAL = 1 // 已认证
)
