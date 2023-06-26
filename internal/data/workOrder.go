package data

import "time"

type WorkOrder struct {
	Id               int64     `gorm:"primary_key;auto_increment;comment:'ID'" json:"id"`
	Code             string    `gorm:"type:varchar(64);charset:utf8;comment:'工单编号'" json:"code"`
	Title            string    `gorm:"type:varchar(64);charset:utf8;comment:'标题'" json:"title"`
	UserID           int64     `gorm:"comment:'发单人'" json:"user_id"`
	CUserID          int64     `gorm:"comment:'结单人'" json:"cuser_id"`
	Type             string    `gorm:"type:varchar(64);charset:utf8;default:'';comment:'类型'" json:"type"`
	IsY              bool      `gorm:"type:tinyint(1);not null;default:0;comment:'处理状态：0 客服 1 运维'" json:"is_y"`
	Status           int64     `gorm:"type:tinyint;not null;default:0;comment:'0 待接单 1 处理中   2 已完成 3 已评价 '" json:"status"`
	Star             int64     `gorm:"type:tinyint(1);not null;default:1;comment:'评价星'" json:"star"`
	Weigh            int64     `gorm:"type:tinyint;not null;default:0;comment:'优先级'" json:"weigh"`
	CreateTime       time.Time `gorm:"comment:'创建时间'" json:"create_time"`
	UpdateTime       time.Time `gorm:"comment:'接单时间'" json:"update_time"`
	EndTime          time.Time `gorm:"comment:'结单时间'" json:"end_time"`
	Info             string    `gorm:"type:text;charset:utf8;comment:'评价'" json:"info"`
	AllTime          int64     `gorm:"not null;default:0;comment:'总时长'" json:"all_time"`
	VmInstanceUuid   string    `gorm:"type:varchar(64);charset:utf8;comment:'关联问题 云主机uuid'" json:"vm_instance_uuid"`
	ReceiveTime      time.Time `gorm:"comment:'接单时间'" json:"receive_time"`
	LastUserID       int64     `gorm:"comment:'最后一位处理人ID'" json:"last_user_id"`
	ToNocTime        string    `gorm:"type:varchar(20);charset:utf8;comment:'转运维时间'" json:"to_noc_time"`
	NocReceiveStatus int64     `gorm:"not null;default:0;comment:'1运维未接单 2运维已接单 0没有转运维'" json:"noc_receive_status"`
	NocType          int64     `gorm:"comment:'转求助对象: 3运维  14cdn'" json:"noc_type"`
	ImgURL           string    `gorm:"type:longtext;charset:utf8;comment:'上传的图片路径'" json:"img_url"`
	FileData         string    `gorm:"type:longtext;charset:utf8;comment:'上传的文件数据'" json:"file_data"`
}
