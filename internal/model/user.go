package model

import "time"

type User struct {
	Id                int64      `gorm:"column:id;primary_key"`
	GroupId           int64      `gorm:"column:group_id;not null;default:0;comment:'组别ID'"`
	TypeId            int64      `gorm:"column:type_id;not null;comment:'类型ID'"`
	Username          string     `gorm:"column:username;type:varchar(32);charset:utf8mb4;not null;default:'';comment:'用户名'"`
	Nickname          string     `gorm:"column:nickname;type:varchar(50);charset:utf8mb4;not null;default:'';comment:'昵称'"`
	Password          string     `gorm:"column:password;type:varchar(32);charset:utf8mb4;not null;default:'';comment:'密码'"`
	Salt              string     `gorm:"column:salt;type:varchar(30);charset:utf8mb4;not null;default:'';comment:'密码盐'"`
	Email             string     `gorm:"column:email;type:varchar(100);charset:utf8mb4;not null;default:'';comment:'电子邮箱'"`
	Mobile            string     `gorm:"column:mobile;type:varchar(15);charset:utf8mb4;not null;default:'';comment:'手机号'"`
	Avatar            string     `gorm:"column:avatar;type:varchar(255);charset:utf8mb4;not null;default:'';comment:'头像'"`
	LevelId           int64      `gorm:"column:level_id;not null;default:0;comment:'等级'"`
	Gender            int64      `gorm:"column:gender;not null;default:0;comment:'性别'"`
	Birthday          *time.Time `gorm:"column:birthday;type:date;comment:'生日'"`
	Bio               string     `gorm:"column:bio;type:varchar(100);charset:utf8mb4;not null;default:'';comment:'格言'"`
	Money             float64    `gorm:"column:money;type:decimal(12,2);not null;default:0.00;comment:'余额'"`
	Cybermoney        float64    `gorm:"column:cybermoney;type:decimal(12,2);not null;default:0.00;comment:'虚拟货币'"`
	Prevtime          int64      `gorm:"column:prevtime;type:int;comment:'上次登录时间'"`
	Logintime         int64      `gorm:"column:logintime;type:int;comment:'登录时间'"`
	Loginip           string     `gorm:"column:loginip;type:varchar(50);charset:utf8mb4;not null;default:'';comment:'登录IP'"`
	Joinip            string     `gorm:"column:joinip;type:varchar(50);charset:utf8mb4;not null;default:'';comment:'加入IP'"`
	Jointime          int64      `gorm:"column:jointime;type:int;comment:'加入时间'"`
	Createtime        int64      `gorm:"column:createtime;type:int;comment:'创建时间'"`
	Updatetime        int64      `gorm:"column:updatetime;type:int;comment:'更新时间'"`
	Token             string     `gorm:"column:token;type:varchar(50);charset:utf8mb4;not null;default:'';comment:'Token'"`
	CategoryId        string     `gorm:"column:category_id;type:varchar(255);charset:utf8mb4;not null;default:'';comment:'类别ID'"`
	Status            uint8      `gorm:"column:status;not null;default:0;comment:'状态 0 正常 1 禁用'"`
	Verification      string     `gorm:"column:verification;type:varchar(255);charset:utf8mb4;not null;default:'';comment:'验证'"`
	TuserId           int64      `gorm:"column:tuser_id;not null;default:0;comment:'推荐人ID'"`
	IsCompany         int64      `gorm:"column:is_company;not null;default:0;comment:'是否企业：0 否，1 认证中，2 已认证'"`
	IsPerson          int64      `gorm:"column:is_person;not null;default:0;comment:'0 个人未认证 1 个人认证中 2 个人已认证'"`
	Remark            string     `gorm:"column:remark;type:varchar(255);charset:utf8mb4;comment:'备注'"`
	FlowOver          string     `gorm:"column:flow_over;type:varchar(50);charset:utf8mb4;not null;default:'0';comment:'超出流量'"`
	LoginTimes        int64      `gorm:"column:login_times;not null;default:0;comment:'登录失败次数'"`
	LoginBindTime     string     `gorm:"column:login_bind_time;type:varchar(20);charset:utf8mb4;not null;default:'0';comment:'账号锁定时间'"`
	LoginLastTime     string     `gorm:"column:login_last_time;type:varchar(20);charset:utf8mb4;not null;default:'0';comment:'最新操作失败时间'"`
	MonthOrFlow       int64      `gorm:"column:month_or_flow;not null;default:0;comment:'0都不是 1流量 2包月带宽'"`
	DiscountEcs       int64      `gorm:"column:discount_ecs;not null;default:100;comment:'ecs折扣'"`
	DiscountCdn       int64      `gorm:"column:discount_cdn;not null;default:100;comment:'cdn折扣'"`
	DiscountOth       int64      `gorm:"column:discount_oth;not null;default:100;comment:'其他折扣'"`
	DoubleCheckSwitch int64      `gorm:"column:double_check_switch;not null;default:0;comment:'登录是否二重认证，默认0 ，0关闭 1开启'"`
	DoubleCheckType   int64      `gorm:"column:double_check_type;not null;default:0;comment:'0未选择 1邮箱 2手机'"`
	MobilePre         string     `gorm:"column:mobile_pre;type:varchar(10);charset:utf8mb4;comment:'手机号前缀'"`
	ChatId            string     `gorm:"column:chat_id;type:varchar(20);charset:utf8mb4;comment:'telegram会话id'"`
}

func (User) TableName() string {
	return "se_user"
}

const (
	// 用户状态
	USER_STATUS_NORMAL = 0 // 正常
	USER_STATUS_FORBID = 1 // 禁用
)
