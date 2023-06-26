package facade

import (
	"context"

	"cdncloud/internal/model"
)

type WorkOrderDetailRepo interface {
	Save(ctx *context.Context, u *model.WorkOrderDetail) (id int64, err error)
	GetWorkOrderDetailByWorkOrderId(ctx *context.Context, workOrderId int64) (workOrderDetail *model.WorkOrderDetail, err error)
}
