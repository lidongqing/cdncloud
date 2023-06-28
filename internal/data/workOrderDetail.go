package data

import (
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type workOrderDetailRepo struct {
	data *Data
	log  *log.Helper
}

func NewWorkOrderDetailRepo(data *Data, logger log.Logger) facade.WorkOrderDetailRepo {
	return &workOrderDetailRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *workOrderDetailRepo) Save(ctx *context.Context, u *model.WorkOrderDetail) (id int64, err error) {
	db := r.data.DataBase
	err = db.Create(&u).Error
	return u.Id, err
}

// 根据工单id获取工单详情列表
func (r *workOrderDetailRepo) GetWorkOrderDetailListByWorkOrderId(ctx *context.Context, workOrderId int64) (workOrderDetailList []*model.WorkOrderDetail, err error) {
	db := r.data.DataBase
	err = db.Where("w_id = ?", workOrderId).Find(&workOrderDetailList).Order("create_time ASC").Error
	return workOrderDetailList, err
}
