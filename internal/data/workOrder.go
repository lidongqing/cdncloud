package data

import (
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type workOrderRepo struct {
	data *Data
	log  *log.Helper
}

func NewWorkOrderRepo(data *Data, logger log.Logger) facade.WorkOrderRepo {
	return &workOrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *workOrderRepo) Save(ctx *context.Context, u *model.WorkOrder) (id int64, err error) {
	db := r.data.DataBase
	err = db.Create(&u).Error
	return u.Id, err
}

// 获取工单列表
func (r *workOrderRepo) GetWorkOrderList(ctx *context.Context, userId int64, status int64, workOrderType string, page int64, pageSize int64) (workOrderList []*model.WorkOrder, count int64, err error) {
	db := r.data.DataBase
	db = db.Where("user_id = ?", userId)
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if workOrderType != "" {
		db = db.Where("type = ?", workOrderType)
	}
	offset := (int(page) - 1) * int(pageSize)
	err = db.Offset(offset).Limit(int(pageSize)).Find(&workOrderList).Error
	if err != nil {
		return nil, 0, err
	}
	// 获取总数
	err = db.Count(&count).Error
	return
}

// 根据id获取工单基础信息
func (r *workOrderRepo) GetWorkOrderById(ctx *context.Context, id int64) (workOrder *model.WorkOrder, err error) {
	db := r.data.DataBase
	err = db.Where("id = ?", id).First(&workOrder).Error
	return
}
