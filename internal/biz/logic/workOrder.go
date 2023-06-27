package logic

import (
	v1WorkOrder "cdncloud/api/v1/workOrder"
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"cdncloud/sessions"
	"context"
	"errors"
	"strconv"
	"time"
)

type WorkOrderLogic struct {
	woRepo        facade.WorkOrderRepo
	wodRepo       facade.WorkOrderDetailRepo
	sessionHandle *sessions.SessionHandle
}

func NewWorkOrderLogic(woRepo facade.WorkOrderRepo, wodRepo facade.WorkOrderDetailRepo, sessionHandle *sessions.SessionHandle) *WorkOrderLogic {
	return &WorkOrderLogic{
		woRepo:        woRepo,
		wodRepo:       wodRepo,
		sessionHandle: sessionHandle,
	}
}

// 创建工单
func (l *WorkOrderLogic) CreateWorkOrder(ctx *context.Context, req *v1WorkOrder.AddWorkOrderRequest) (workOrderId int64, err error) {
	userId, err := l.GetUserIdBySession(ctx)
	if err != nil {
		return 0, err
	}
	// 创建工单
	// @todo:生成工单编号
	workOrderCode := ""
	// @todo:保存图片
	imageUrl := ""
	workOrder := &model.WorkOrder{
		Code:        workOrderCode,
		UserID:      userId,
		Title:       req.Title,
		Status:      1,
		Type:        req.Type,
		Weigh:       req.Weigh,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		EndTime:     time.Now(),
		ReceiveTime: time.Now(),
		ImgURL:      imageUrl,
		FileData:    req.Image,
	}
	workOrderId, err = l.woRepo.Save(ctx, workOrder)
	if err != nil {
		return 0, err
	}
	// 创建工单详情
	workOrderDetail := &model.WorkOrderDetail{
		WorkOrderID: workOrderId,
		UserID:      userId,
		Content:     req.Content,
		CreateTime:  time.Now(),
		ImgURL:      imageUrl,
		FileData:    req.Image,
	}
	_, err = l.wodRepo.Save(ctx, workOrderDetail)
	if err != nil {
		return 0, err
	}
	return workOrderId, nil
}

// 获取工单列表
func (l *WorkOrderLogic) GetWorkOrderList(ctx *context.Context, req *v1WorkOrder.GetWorkOrderListRequest) (workOrderRes *v1WorkOrder.GetWorkOrderListReply, err error) {
	userId, err := l.GetUserIdBySession(ctx)
	if err != nil {
		return nil, err
	}

	workOrderList, count, err := l.woRepo.GetWorkOrderList(ctx, userId, req.Status, req.Type, req.Page, 10)
	if err != nil {
		return nil, err
	}
	workOrderRes = &v1WorkOrder.GetWorkOrderListReply{
		Total: count,
		List:  make([]*v1WorkOrder.GetWorkOrderListReply_WorkOrderListItem, 0),
		Page:  req.Page,
	}
	for _, workOrder := range workOrderList {
		workOrderRes.List = append(workOrderRes.List, &v1WorkOrder.GetWorkOrderListReply_WorkOrderListItem{
			Code:       workOrder.Code,
			Status:     workOrder.Status,
			Type:       workOrder.Type,
			CreateTime: workOrder.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

// session读取用户id
func (wol *WorkOrderLogic) GetUserIdBySession(ctx *context.Context) (userId int64, err error) {
	userIdStr, err := wol.sessionHandle.GetSession(*ctx, "user_id")
	if err != nil {
		return 0, err
	}
	if userIdStr == "" {
		return 0, errors.New("用户未登录")
	}
	userId, _ = strconv.ParseInt(userIdStr, 10, 64)
	return
}
