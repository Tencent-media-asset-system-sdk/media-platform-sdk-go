// Code generated by trpc-go/trpc-go-cmdline v2.4.0. DO NOT EDIT.
// source: ai_video_process.proto

package aivideoprocess

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/codec"
	_ "github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/http"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/server"
)

// START ======================================= Server Service Definition ======================================= START

// AIVideoProcessService defines service.
type AIVideoProcessService interface {
	// CreateAIVideoProcessTask 创建AI视频处理任务
	CreateAIVideoProcessTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, rsp *CreateAIVideoProcessResponse) (err error) // @alias=/CreateAIVideoProcessTask
	// DescribeAIVideoProcessDetail 获取任务详情
	DescribeAIVideoProcessDetail(ctx context.Context, req *DescribeAIVideoProcessDetailRequest, rsp *DescribeAIVideoProcessDetailResponse) (err error) // @alias=/DescribeAIVideoProcessDetail
	// DescribeAIVideoProcessTasks 获取任务列表 (TODO: 前端切换调用后清理接口)
	//  NOTICE: deprecated, use DescribeTasks instead
	DescribeAIVideoProcessTasks(ctx context.Context, req *DescribeAIVideoProcessTasksRequest, rsp *DescribeAIVideoProcessTasksResponse) (err error) // @alias=/DescribeAIVideoProcessTasks
	// UpdateAIVideoProcessTask 更新AI视频处理任务
	UpdateAIVideoProcessTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, rsp *UpdateAIVideoProcessResponse) (err error) // @alias=/UpdateAIVideoProcessTask
	// DescribeAIVideoProcessConfig 获取视频处理配置
	DescribeAIVideoProcessConfig(ctx context.Context, req *DescribeAIVideoProcessConfigRequest, rsp *DescribeAIVideoProcessConfigResponse) (err error) // @alias=/DescribeAIVideoProcessConfig
	// CreateVideoCropTask 创建横竖屏任务 (NOTICE: 已合并到CreateAIVideoProcessTask)
	//  NOTICE: deprecated, use CreateAIVideoProcessTask instead
	CreateVideoCropTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, rsp *CreateAIVideoProcessResponse) (err error) // @alias=/CreateVideoCropTask
	// UpdateVideoCropTask 更新横竖屏任务 (NOTICE: 已合并到UpdateAIVideoProcessTask)
	//  NOTICE: deprecated, use UpdateAIVideoProcessTask instead
	UpdateVideoCropTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, rsp *UpdateAIVideoProcessResponse) (err error) // @alias=/UpdateVideoCropTask
	// CreateM3UIndexTask 创建M3U Index任务
	CreateM3UIndexTask(ctx context.Context, req *CreateM3UIndexTaskRequest, rsp *CreateM3UIndexTaskResponse) (err error) // @alias=/CreateM3UIndexTask
	// DescribeM3UIndexDetail 获取M3U Index任务详情
	DescribeM3UIndexDetail(ctx context.Context, req *DescribeM3UIndexDetailRequest, rsp *DescribeM3UIndexDetailResponse) (err error) // @alias=/DescribeM3UIndexDetail
	// StopAIVideoProcessTask 停止任务，内部使用对外不暴露
	StopAIVideoProcessTask(ctx context.Context, req *StopAIVideoProcessTaskReq, rsp *StopAIVideoProcessTaskRsp) (err error) // @alias=/StopAIVideoProcessTask
	// RemoveAIVideoProcessTask 删除任务，内部使用对外不暴露
	RemoveAIVideoProcessTask(ctx context.Context, req *RemoveAIVideoProcessTaskReq, rsp *RemoveAIVideoProcessTaskRsp) (err error) // @alias=/RemoveAIVideoProcessTask
}

func AIVideoProcessService_CreateAIVideoProcessTask_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &CreateAIVideoProcessTaskRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).CreateAIVideoProcessTask(ctx, reqbody.(*CreateAIVideoProcessTaskRequest), rspbody.(*CreateAIVideoProcessResponse))
	}

	rsp := &CreateAIVideoProcessResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_DescribeAIVideoProcessDetail_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeAIVideoProcessDetailRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).DescribeAIVideoProcessDetail(ctx, reqbody.(*DescribeAIVideoProcessDetailRequest), rspbody.(*DescribeAIVideoProcessDetailResponse))
	}

	rsp := &DescribeAIVideoProcessDetailResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_DescribeAIVideoProcessTasks_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeAIVideoProcessTasksRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).DescribeAIVideoProcessTasks(ctx, reqbody.(*DescribeAIVideoProcessTasksRequest), rspbody.(*DescribeAIVideoProcessTasksResponse))
	}

	rsp := &DescribeAIVideoProcessTasksResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_UpdateAIVideoProcessTask_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &UpdateAIVideoProcessTaskRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).UpdateAIVideoProcessTask(ctx, reqbody.(*UpdateAIVideoProcessTaskRequest), rspbody.(*UpdateAIVideoProcessResponse))
	}

	rsp := &UpdateAIVideoProcessResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_DescribeAIVideoProcessConfig_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeAIVideoProcessConfigRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).DescribeAIVideoProcessConfig(ctx, reqbody.(*DescribeAIVideoProcessConfigRequest), rspbody.(*DescribeAIVideoProcessConfigResponse))
	}

	rsp := &DescribeAIVideoProcessConfigResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_CreateVideoCropTask_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &CreateAIVideoProcessTaskRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).CreateVideoCropTask(ctx, reqbody.(*CreateAIVideoProcessTaskRequest), rspbody.(*CreateAIVideoProcessResponse))
	}

	rsp := &CreateAIVideoProcessResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_UpdateVideoCropTask_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &UpdateAIVideoProcessTaskRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).UpdateVideoCropTask(ctx, reqbody.(*UpdateAIVideoProcessTaskRequest), rspbody.(*UpdateAIVideoProcessResponse))
	}

	rsp := &UpdateAIVideoProcessResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_CreateM3UIndexTask_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &CreateM3UIndexTaskRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).CreateM3UIndexTask(ctx, reqbody.(*CreateM3UIndexTaskRequest), rspbody.(*CreateM3UIndexTaskResponse))
	}

	rsp := &CreateM3UIndexTaskResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_DescribeM3UIndexDetail_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeM3UIndexDetailRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).DescribeM3UIndexDetail(ctx, reqbody.(*DescribeM3UIndexDetailRequest), rspbody.(*DescribeM3UIndexDetailResponse))
	}

	rsp := &DescribeM3UIndexDetailResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_StopAIVideoProcessTask_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &StopAIVideoProcessTaskReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).StopAIVideoProcessTask(ctx, reqbody.(*StopAIVideoProcessTaskReq), rspbody.(*StopAIVideoProcessTaskRsp))
	}

	rsp := &StopAIVideoProcessTaskRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func AIVideoProcessService_RemoveAIVideoProcessTask_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &RemoveAIVideoProcessTaskReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(AIVideoProcessService).RemoveAIVideoProcessTask(ctx, reqbody.(*RemoveAIVideoProcessTaskReq), rspbody.(*RemoveAIVideoProcessTaskRsp))
	}

	rsp := &RemoveAIVideoProcessTaskRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// AIVideoProcessServer_ServiceDesc descriptor for server.RegisterService.
var AIVideoProcessServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.media.aivideoprocess.AIVideoProcess",
	HandlerType: ((*AIVideoProcessService)(nil)),
	Methods: []server.Method{
		{
			Name: "/CreateAIVideoProcessTask",
			Func: AIVideoProcessService_CreateAIVideoProcessTask_Handler,
		},
		{
			Name: "/DescribeAIVideoProcessDetail",
			Func: AIVideoProcessService_DescribeAIVideoProcessDetail_Handler,
		},
		{
			Name: "/DescribeAIVideoProcessTasks",
			Func: AIVideoProcessService_DescribeAIVideoProcessTasks_Handler,
		},
		{
			Name: "/UpdateAIVideoProcessTask",
			Func: AIVideoProcessService_UpdateAIVideoProcessTask_Handler,
		},
		{
			Name: "/DescribeAIVideoProcessConfig",
			Func: AIVideoProcessService_DescribeAIVideoProcessConfig_Handler,
		},
		{
			Name: "/CreateVideoCropTask",
			Func: AIVideoProcessService_CreateVideoCropTask_Handler,
		},
		{
			Name: "/UpdateVideoCropTask",
			Func: AIVideoProcessService_UpdateVideoCropTask_Handler,
		},
		{
			Name: "/CreateM3UIndexTask",
			Func: AIVideoProcessService_CreateM3UIndexTask_Handler,
		},
		{
			Name: "/DescribeM3UIndexDetail",
			Func: AIVideoProcessService_DescribeM3UIndexDetail_Handler,
		},
		{
			Name: "/StopAIVideoProcessTask",
			Func: AIVideoProcessService_StopAIVideoProcessTask_Handler,
		},
		{
			Name: "/RemoveAIVideoProcessTask",
			Func: AIVideoProcessService_RemoveAIVideoProcessTask_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/CreateAIVideoProcessTask",
			Func: AIVideoProcessService_CreateAIVideoProcessTask_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/DescribeAIVideoProcessDetail",
			Func: AIVideoProcessService_DescribeAIVideoProcessDetail_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/DescribeAIVideoProcessTasks",
			Func: AIVideoProcessService_DescribeAIVideoProcessTasks_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/UpdateAIVideoProcessTask",
			Func: AIVideoProcessService_UpdateAIVideoProcessTask_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/DescribeAIVideoProcessConfig",
			Func: AIVideoProcessService_DescribeAIVideoProcessConfig_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/CreateVideoCropTask",
			Func: AIVideoProcessService_CreateVideoCropTask_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/UpdateVideoCropTask",
			Func: AIVideoProcessService_UpdateVideoCropTask_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/CreateM3UIndexTask",
			Func: AIVideoProcessService_CreateM3UIndexTask_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/DescribeM3UIndexDetail",
			Func: AIVideoProcessService_DescribeM3UIndexDetail_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/StopAIVideoProcessTask",
			Func: AIVideoProcessService_StopAIVideoProcessTask_Handler,
		},
		{
			Name: "/trpc.media.aivideoprocess.AIVideoProcess/RemoveAIVideoProcessTask",
			Func: AIVideoProcessService_RemoveAIVideoProcessTask_Handler,
		},
	},
}

// RegisterAIVideoProcessService registers service.
func RegisterAIVideoProcessService(s server.Service, svr AIVideoProcessService) {
	if err := s.Register(&AIVideoProcessServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("AIVideoProcess register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedAIVideoProcess struct{}

// CreateAIVideoProcessTask 创建AI视频处理任务
func (s *UnimplementedAIVideoProcess) CreateAIVideoProcessTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, rsp *CreateAIVideoProcessResponse) error {
	return errors.New("rpc CreateAIVideoProcessTask of service AIVideoProcess is not implemented")
}

// DescribeAIVideoProcessDetail 获取任务详情
func (s *UnimplementedAIVideoProcess) DescribeAIVideoProcessDetail(ctx context.Context, req *DescribeAIVideoProcessDetailRequest, rsp *DescribeAIVideoProcessDetailResponse) error {
	return errors.New("rpc DescribeAIVideoProcessDetail of service AIVideoProcess is not implemented")
}

// DescribeAIVideoProcessTasks 获取任务列表 (TODO: 前端切换调用后清理接口)
//  NOTICE: deprecated, use DescribeTasks instead
func (s *UnimplementedAIVideoProcess) DescribeAIVideoProcessTasks(ctx context.Context, req *DescribeAIVideoProcessTasksRequest, rsp *DescribeAIVideoProcessTasksResponse) error {
	return errors.New("rpc DescribeAIVideoProcessTasks of service AIVideoProcess is not implemented")
}

// UpdateAIVideoProcessTask 更新AI视频处理任务
func (s *UnimplementedAIVideoProcess) UpdateAIVideoProcessTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, rsp *UpdateAIVideoProcessResponse) error {
	return errors.New("rpc UpdateAIVideoProcessTask of service AIVideoProcess is not implemented")
}

// DescribeAIVideoProcessConfig 获取视频处理配置
func (s *UnimplementedAIVideoProcess) DescribeAIVideoProcessConfig(ctx context.Context, req *DescribeAIVideoProcessConfigRequest, rsp *DescribeAIVideoProcessConfigResponse) error {
	return errors.New("rpc DescribeAIVideoProcessConfig of service AIVideoProcess is not implemented")
}

// CreateVideoCropTask 创建横竖屏任务 (NOTICE: 已合并到CreateAIVideoProcessTask)
//  NOTICE: deprecated, use CreateAIVideoProcessTask instead
func (s *UnimplementedAIVideoProcess) CreateVideoCropTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, rsp *CreateAIVideoProcessResponse) error {
	return errors.New("rpc CreateVideoCropTask of service AIVideoProcess is not implemented")
}

// UpdateVideoCropTask 更新横竖屏任务 (NOTICE: 已合并到UpdateAIVideoProcessTask)
//  NOTICE: deprecated, use UpdateAIVideoProcessTask instead
func (s *UnimplementedAIVideoProcess) UpdateVideoCropTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, rsp *UpdateAIVideoProcessResponse) error {
	return errors.New("rpc UpdateVideoCropTask of service AIVideoProcess is not implemented")
}

// CreateM3UIndexTask 创建M3U Index任务
func (s *UnimplementedAIVideoProcess) CreateM3UIndexTask(ctx context.Context, req *CreateM3UIndexTaskRequest, rsp *CreateM3UIndexTaskResponse) error {
	return errors.New("rpc CreateM3UIndexTask of service AIVideoProcess is not implemented")
}

// DescribeM3UIndexDetail 获取M3U Index任务详情
func (s *UnimplementedAIVideoProcess) DescribeM3UIndexDetail(ctx context.Context, req *DescribeM3UIndexDetailRequest, rsp *DescribeM3UIndexDetailResponse) error {
	return errors.New("rpc DescribeM3UIndexDetail of service AIVideoProcess is not implemented")
}

// StopAIVideoProcessTask 停止任务，内部使用对外不暴露
func (s *UnimplementedAIVideoProcess) StopAIVideoProcessTask(ctx context.Context, req *StopAIVideoProcessTaskReq, rsp *StopAIVideoProcessTaskRsp) error {
	return errors.New("rpc StopAIVideoProcessTask of service AIVideoProcess is not implemented")
}

// RemoveAIVideoProcessTask 删除任务，内部使用对外不暴露
func (s *UnimplementedAIVideoProcess) RemoveAIVideoProcessTask(ctx context.Context, req *RemoveAIVideoProcessTaskReq, rsp *RemoveAIVideoProcessTaskRsp) error {
	return errors.New("rpc RemoveAIVideoProcessTask of service AIVideoProcess is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// AIVideoProcessClientProxy defines service client proxy
type AIVideoProcessClientProxy interface {
	// CreateAIVideoProcessTask 创建AI视频处理任务
	CreateAIVideoProcessTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, opts ...client.Option) (rsp *CreateAIVideoProcessResponse, err error) // @alias=/CreateAIVideoProcessTask
	// DescribeAIVideoProcessDetail 获取任务详情
	DescribeAIVideoProcessDetail(ctx context.Context, req *DescribeAIVideoProcessDetailRequest, opts ...client.Option) (rsp *DescribeAIVideoProcessDetailResponse, err error) // @alias=/DescribeAIVideoProcessDetail
	// DescribeAIVideoProcessTasks 获取任务列表 (TODO: 前端切换调用后清理接口)
	//  NOTICE: deprecated, use DescribeTasks instead
	DescribeAIVideoProcessTasks(ctx context.Context, req *DescribeAIVideoProcessTasksRequest, opts ...client.Option) (rsp *DescribeAIVideoProcessTasksResponse, err error) // @alias=/DescribeAIVideoProcessTasks
	// UpdateAIVideoProcessTask 更新AI视频处理任务
	UpdateAIVideoProcessTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, opts ...client.Option) (rsp *UpdateAIVideoProcessResponse, err error) // @alias=/UpdateAIVideoProcessTask
	// DescribeAIVideoProcessConfig 获取视频处理配置
	DescribeAIVideoProcessConfig(ctx context.Context, req *DescribeAIVideoProcessConfigRequest, opts ...client.Option) (rsp *DescribeAIVideoProcessConfigResponse, err error) // @alias=/DescribeAIVideoProcessConfig
	// CreateVideoCropTask 创建横竖屏任务 (NOTICE: 已合并到CreateAIVideoProcessTask)
	//  NOTICE: deprecated, use CreateAIVideoProcessTask instead
	CreateVideoCropTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, opts ...client.Option) (rsp *CreateAIVideoProcessResponse, err error) // @alias=/CreateVideoCropTask
	// UpdateVideoCropTask 更新横竖屏任务 (NOTICE: 已合并到UpdateAIVideoProcessTask)
	//  NOTICE: deprecated, use UpdateAIVideoProcessTask instead
	UpdateVideoCropTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, opts ...client.Option) (rsp *UpdateAIVideoProcessResponse, err error) // @alias=/UpdateVideoCropTask
	// CreateM3UIndexTask 创建M3U Index任务
	CreateM3UIndexTask(ctx context.Context, req *CreateM3UIndexTaskRequest, opts ...client.Option) (rsp *CreateM3UIndexTaskResponse, err error) // @alias=/CreateM3UIndexTask
	// DescribeM3UIndexDetail 获取M3U Index任务详情
	DescribeM3UIndexDetail(ctx context.Context, req *DescribeM3UIndexDetailRequest, opts ...client.Option) (rsp *DescribeM3UIndexDetailResponse, err error) // @alias=/DescribeM3UIndexDetail
	// StopAIVideoProcessTask 停止任务，内部使用对外不暴露
	StopAIVideoProcessTask(ctx context.Context, req *StopAIVideoProcessTaskReq, opts ...client.Option) (rsp *StopAIVideoProcessTaskRsp, err error) // @alias=/StopAIVideoProcessTask
	// RemoveAIVideoProcessTask 删除任务，内部使用对外不暴露
	RemoveAIVideoProcessTask(ctx context.Context, req *RemoveAIVideoProcessTaskReq, opts ...client.Option) (rsp *RemoveAIVideoProcessTaskRsp, err error) // @alias=/RemoveAIVideoProcessTask
}

type AIVideoProcessClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewAIVideoProcessClientProxy = func(opts ...client.Option) AIVideoProcessClientProxy {
	return &AIVideoProcessClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *AIVideoProcessClientProxyImpl) CreateAIVideoProcessTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, opts ...client.Option) (*CreateAIVideoProcessResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/CreateAIVideoProcessTask")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("CreateAIVideoProcessTask")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CreateAIVideoProcessResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) DescribeAIVideoProcessDetail(ctx context.Context, req *DescribeAIVideoProcessDetailRequest, opts ...client.Option) (*DescribeAIVideoProcessDetailResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeAIVideoProcessDetail")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("DescribeAIVideoProcessDetail")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeAIVideoProcessDetailResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) DescribeAIVideoProcessTasks(ctx context.Context, req *DescribeAIVideoProcessTasksRequest, opts ...client.Option) (*DescribeAIVideoProcessTasksResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeAIVideoProcessTasks")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("DescribeAIVideoProcessTasks")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeAIVideoProcessTasksResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) UpdateAIVideoProcessTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, opts ...client.Option) (*UpdateAIVideoProcessResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/UpdateAIVideoProcessTask")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("UpdateAIVideoProcessTask")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &UpdateAIVideoProcessResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) DescribeAIVideoProcessConfig(ctx context.Context, req *DescribeAIVideoProcessConfigRequest, opts ...client.Option) (*DescribeAIVideoProcessConfigResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeAIVideoProcessConfig")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("DescribeAIVideoProcessConfig")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeAIVideoProcessConfigResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) CreateVideoCropTask(ctx context.Context, req *CreateAIVideoProcessTaskRequest, opts ...client.Option) (*CreateAIVideoProcessResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/CreateVideoCropTask")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("CreateVideoCropTask")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CreateAIVideoProcessResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) UpdateVideoCropTask(ctx context.Context, req *UpdateAIVideoProcessTaskRequest, opts ...client.Option) (*UpdateAIVideoProcessResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/UpdateVideoCropTask")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("UpdateVideoCropTask")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &UpdateAIVideoProcessResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) CreateM3UIndexTask(ctx context.Context, req *CreateM3UIndexTaskRequest, opts ...client.Option) (*CreateM3UIndexTaskResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/CreateM3UIndexTask")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("CreateM3UIndexTask")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CreateM3UIndexTaskResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) DescribeM3UIndexDetail(ctx context.Context, req *DescribeM3UIndexDetailRequest, opts ...client.Option) (*DescribeM3UIndexDetailResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeM3UIndexDetail")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("DescribeM3UIndexDetail")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeM3UIndexDetailResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) StopAIVideoProcessTask(ctx context.Context, req *StopAIVideoProcessTaskReq, opts ...client.Option) (*StopAIVideoProcessTaskRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/StopAIVideoProcessTask")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("StopAIVideoProcessTask")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &StopAIVideoProcessTaskRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *AIVideoProcessClientProxyImpl) RemoveAIVideoProcessTask(ctx context.Context, req *RemoveAIVideoProcessTaskReq, opts ...client.Option) (*RemoveAIVideoProcessTaskRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/RemoveAIVideoProcessTask")
	msg.WithCalleeServiceName(AIVideoProcessServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("aivideoprocess")
	msg.WithCalleeService("AIVideoProcess")
	msg.WithCalleeMethod("RemoveAIVideoProcessTask")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &RemoveAIVideoProcessTaskRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
