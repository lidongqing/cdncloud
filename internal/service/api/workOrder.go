package api

import (
	wo "cdncloud/api/v1/workOrder"
	"cdncloud/internal/biz/logic"
	"context"
)

type WorkOrderService struct {
	wo.UnimplementedWorkOrderServer
	ul *logic.UserLogic
}

func NewWorkOrderService(ul *logic.UserLogic) *WorkOrderService {
	return &WorkOrderService{
		ul: ul,
	}
}

// 工单提交
func (s *WorkOrderService) AddWorkOrder(ctx context.Context, in *wo.AddWorkOrderRequest) (*wo.EmptyReply, error) {
	return &wo.EmptyReply{}, nil
}

// 工单列表
func (s *WorkOrderService) GetWorkOrderList(ctx context.Context, in *wo.GetWorkOrderListRequest) (*wo.GetWorkOrderListReply, error) {
	return &wo.GetWorkOrderListReply{}, nil
}
