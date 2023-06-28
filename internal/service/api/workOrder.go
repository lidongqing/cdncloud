package api

import (
	wo "cdncloud/api/v1/workOrder"
	"cdncloud/internal/biz/logic"
	"context"
	"errors"
)

type WorkOrderService struct {
	wo.UnimplementedWorkOrderServer
	wol *logic.WorkOrderLogic
}

func NewWorkOrderService(wol *logic.WorkOrderLogic) *WorkOrderService {
	return &WorkOrderService{
		wol: wol,
	}
}

// 工单提交
func (s *WorkOrderService) AddWorkOrder(ctx context.Context, in *wo.AddWorkOrderRequest) (*wo.EmptyReply, error) {
	if in.Title == "" {
		return nil, errors.New("标题不能为空")
	}
	if in.Content == "" {
		return nil, errors.New("内容不能为空")
	}
	if in.Type == "" {
		return nil, errors.New("类型不能为空")
	}
	if in.Weigh == 0 {
		return nil, errors.New("优先级不能为空")
	}
	_, err := s.wol.CreateWorkOrder(&ctx, in)
	return &wo.EmptyReply{}, err
}

// 工单列表
func (s *WorkOrderService) GetWorkOrderList(ctx context.Context, in *wo.GetWorkOrderListRequest) (*wo.GetWorkOrderListReply, error) {
	return s.wol.GetWorkOrderList(&ctx, in)
}

// 工单详情
func (s *WorkOrderService) GetWorkOrderDetail(ctx context.Context, in *wo.GetWorkOrderDetailRequest) (*wo.GetWorkOrderDetailReply, error) {
	return s.wol.GetWorkOrderDetail(&ctx, in)
}
