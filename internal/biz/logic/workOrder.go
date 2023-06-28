package logic

import (
	v1WorkOrder "cdncloud/api/v1/workOrder"
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"cdncloud/sessions"
	"context"
	"errors"
	"fmt"
	"math/rand"
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
	workOrderCode := l.CreateWorkOrderCode()
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
			Id:         workOrder.Id,
			Code:       workOrder.Code,
			Status:     workOrder.Status,
			StatusName: l.GetWorkOrderStatus(workOrder.Status),
			Type:       workOrder.Type,
			CreateTime: workOrder.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

// 获取工单详情
func (l *WorkOrderLogic) GetWorkOrderDetail(ctx *context.Context, req *v1WorkOrder.GetWorkOrderDetailRequest) (workOrderDetailRes *v1WorkOrder.GetWorkOrderDetailReply, err error) {
	// 获取工单基础信息
	workOrder, err := l.woRepo.GetWorkOrderById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 工单详情列表
	workOrderDetailList, err := l.wodRepo.GetWorkOrderDetailListByWorkOrderId(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	workOrderDetailRes = &v1WorkOrder.GetWorkOrderDetailReply{
		Title:      workOrder.Title,
		Code:       workOrder.Code,
		Status:     workOrder.Status,
		StatusName: l.GetWorkOrderStatus(workOrder.Status),
		Type:       workOrder.Type,
		Weigh:      workOrder.Weigh,
		WeighName:  l.GetWorkOrderWeigh(workOrder.Weigh),
		CreateTime: workOrder.CreateTime.Format("2006-01-02 15:04:05"),
		ReplyList:  make([]*v1WorkOrder.WorkOrderDetailListItem, 0),
	}
	for _, workOrderDetail := range workOrderDetailList {
		workOrderDetailRes.ReplyList = append(workOrderDetailRes.ReplyList, &v1WorkOrder.WorkOrderDetailListItem{
			Content:    workOrderDetail.Content,
			CreateTime: workOrderDetail.CreateTime.Format("2006-01-02 15:04:05"),
			Type:       workOrderDetail.Type,
			Look:       workOrderDetail.Look,
		})
	}
	return
}

// 生成工单编号
func (l *WorkOrderLogic) CreateWorkOrderCode() string {
	// 生成当前时间
	now := time.Now()
	// 格式化时间为字符串
	timeStr := now.Format("20060102150405")
	// 生成3位随机数字
	rand.Seed(now.UnixNano())
	randNum := rand.Intn(1000)
	// 拼接字符串
	code := fmt.Sprintf("%s%03d", timeStr, randNum)
	return code
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

// 读取工单状态
func (wol *WorkOrderLogic) GetWorkOrderStatus(status int64) string {
	switch status {
	case model.WORK_ORDER_STATUS_WAIT:
		return "待接单"
	case model.WORK_ORDER_STATUS_ING:
		return "处理中"
	case model.WORK_ORDER_STATUS_END:
		return "已完成"
	case model.WORK_ORDER_STATUS_EVAL:
		return "已评价"
	case model.WORK_ORDER_STATUS_CANCEL:
		return "已取消"
	default:
		return ""
	}
}

// 读取工单优先级
func (wol *WorkOrderLogic) GetWorkOrderWeigh(weigh int64) string {
	switch weigh {
	case model.WORK_ORDER_WEIGH_LOW:
		return "普通"
	case model.WORK_ORDER_WEIGH_HIGH:
		return "紧急"
	default:
		return ""
	}
}

// 获取所有工单类型
func (wol *WorkOrderLogic) GetWorkOrderType() []string {
	return []string{
		"售前",
		"云主机",
		"CDN",
		"财务",
		"CDNex",
		"其他",
	}
}
