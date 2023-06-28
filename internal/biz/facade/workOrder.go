package facade

import (
	"context"

	"cdncloud/internal/model"
)

type WorkOrderRepo interface {
	Save(ctx *context.Context, u *model.WorkOrder) (id int64, err error)
	GetWorkOrderList(ctx *context.Context, userId int64, status int64, workOrderType string, page int64, pageSize int64) (workOrderList []*model.WorkOrder, count int64, err error)
	GetWorkOrderById(ctx *context.Context, id int64) (workOrder *model.WorkOrder, err error)
}
