package model

type UserCompany struct {
	Id      int64  `gorm:"primary_key"`
	UserID  int64  `gorm:"not null;comment:'用户ID'"`
	Name    string `gorm:"type:varchar(50);default:'';charset:utf8mb4;collation:utf8mb4_unicode_ci;comment:'组名'"`
	Code    string `gorm:"type:varchar(255);charset:utf8mb4;collation:utf8mb4_unicode_ci;comment:'统一信用代码'"`
	Epreson string `gorm:"type:varchar(64);not null;charset:utf8mb4;collation:utf8mb4_unicode_ci;comment:'企业法人'"`
	Ecard   string `gorm:"type:varchar(64);not null;charset:utf8mb4;collation:utf8mb4_unicode_ci;comment:'法人身份证'"`
	Phone   string `gorm:"type:varchar(255);not null;charset:utf8mb4;collation:utf8mb4_unicode_ci;comment:'联系电话'"`
	Address string `gorm:"type:varchar(255);not null;charset:utf8mb4;collation:utf8mb4_unicode_ci;comment:'联系地址'"`
	Image   string `gorm:"type:varchar(255);not null;charset:utf8mb4;collation:utf8mb4_unicode_ci;comment:'营业执照'"`
	Status  int64  `gorm:"type:tinyint(1);not null;default:0;comment:'0 待审核 1 已审核'"`
}

func (UserCompany) TableName() string {
	return "se_user_company"
}

const (
	// 企业认证状态
	USER_COMPANY_STATUS_WAIT   = 0 // 待认证
	USER_COMPANY_STATUS_NORMAL = 1 // 已认证
)
