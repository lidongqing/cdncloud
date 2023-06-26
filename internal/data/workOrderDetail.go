package data

import "time"

type WorkOrderDetail struct {
	ID           uint      `gorm:"primary_key;column:id;type:int(10) unsigned;not null;auto_increment;comment:'ID'"`
	UserID       int       `gorm:"column:user_id;type:int(10);comment:'发送人'"`
	WorkOrderID  int       `gorm:"column:w_id;type:int(10);comment:'工单ID'"`
	Content      string    `gorm:"column:content;type:mediumtext;character set:utf8;collation:utf8_general_ci;comment:'内容'"`
	Type         int       `gorm:"column:type;type:tinyint(4);comment:'类型：0 用户 1 客服 2 运维'"`
	Look         int       `gorm:"column:look;type:tinyint(1);not null;default:0;comment:'0 未查看 1 已查看'"`
	Room         int       `gorm:"column:room;type:tinyint(1);not null;default:0;comment:'0 用户客服 1 客服运维'"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;comment:'创建时间'"`
	ReactionTime int       `gorm:"column:reaction_time;type:int(10);not null;default:0;comment:'反应时长（跟最新数据的时间差）'"`
	ImgURL       string    `gorm:"column:img_url;type:longtext;character set:utf8;collation:utf8_general_ci;comment:'上传的图片路径'"`
	FileData     string    `gorm:"column:file_data;type:longtext;character set:utf8;collation:utf8_general_ci;comment:'上传文件数据'"`
}
