// Code generated by trpc-go/trpc-go-cmdline. DO NOT EDIT.
// source: person_user_define.proto

package personuserdefine

import (
	"context"
	"fmt"

	_ "github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock"
	_ "github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/http"

	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/codec"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/server"
)

/* ************************************ Service Definition ************************************ */

// UserDefinePersonService defines service
type UserDefinePersonService interface {

	// CreateUserDefineGroup 自定义人物库
	CreateUserDefineGroup(ctx context.Context, req *CreateUserDefineGroupRequest) (*CreateUserDefineGroupResponse, error) // @alias=/CreateUserDefineGroup

	DescribeUserDefineGroup(ctx context.Context, req *DescribeUserDefineGroupRequest) (*DescribeUserDefineGroupResponse, error) // @alias=/DescribeUserDefineGroup

	RemoveUserDefineGroup(ctx context.Context, req *RemoveUserDefineGroupRequest) (*RemoveUserDefineGroupResponse, error) // @alias=/RemoveUserDefineGroup

	UpdateUserDefineGroup(ctx context.Context, req *UpdateUserDefineGroupRequest) (*UpdateUserDefineGroupResponse, error) // @alias=/UpdateUserDefineGroup

	// CreateUserDefinePerson 自定义人脸
	CreateUserDefinePerson(ctx context.Context, req *CreateUserDefinePersonRequest) (*CreateUserDefinePersonResponse, error) // @alias=/CreateUserDefinePerson

	RemoveUserDefinePerson(ctx context.Context, req *RemoveUserDefinePersonRequest) (*RemoveUserDefinePersonResponse, error) // @alias=/RemoveUserDefinePerson

	UpdateUserDefinePerson(ctx context.Context, req *UpdateUserDefinePersonRequest) (*UpdateUserDefinePersonResponse, error) // @alias=/UpdateUserDefinePerson

	DescribeUserDefinePerson(ctx context.Context, req *DescribeUserDefinePersonRequest) (*DescribeUserDefinePersonResponse, error) // @alias=/DescribeUserDefinePerson

	DescribeUserDefinePersonDetail(ctx context.Context, req *DescribeUserDefinePersonDetailRequest) (*DescribeUserDefinePersonDetailResponse, error) // @alias=/DescribeUserDefinePersonDetail

	DescribeUserDefinePersonWithFeature(ctx context.Context, req *DescribeUserDefinePersonWithFeatureRequest) (*DescribeUserDefinePersonWithFeatureResponse, error) // @alias=/DescribeUserDefinePersonWithFeature

	// CreatePersonCategory 人物分类
	CreatePersonCategory(ctx context.Context, req *CreatePersonCategoryRequest) (*CreatePersonCategoryResponse, error) // @alias=/CreatePersonCategory

	RemovePersonCategory(ctx context.Context, req *RemovePersonCategoryRequest) (*RemovePersonCategoryResponse, error) // @alias=/RemovePersonCategory

	DescribePersonCategory(ctx context.Context, req *DescribePersonCategoryRequest) (*DescribePersonCategoryResponse, error) // @alias=/DescribePersonCategory

	// RevertPerson 翻库
	RevertPerson(ctx context.Context, req *PersonRevertRequest) (*PersonRevertResponse, error) // @alias=/RevertPerson

	DescribeRevertTasks(ctx context.Context, req *DescribeRevertTasksRequest) (*DescribeRevertTasksResponse, error) // @alias=/DescribeRevertTasks
}

func UserDefinePersonService_CreateUserDefineGroup_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &CreateUserDefineGroupRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).CreateUserDefineGroup(ctx, reqbody.(*CreateUserDefineGroupRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_DescribeUserDefineGroup_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &DescribeUserDefineGroupRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).DescribeUserDefineGroup(ctx, reqbody.(*DescribeUserDefineGroupRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_RemoveUserDefineGroup_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &RemoveUserDefineGroupRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).RemoveUserDefineGroup(ctx, reqbody.(*RemoveUserDefineGroupRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_UpdateUserDefineGroup_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &UpdateUserDefineGroupRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).UpdateUserDefineGroup(ctx, reqbody.(*UpdateUserDefineGroupRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_CreateUserDefinePerson_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &CreateUserDefinePersonRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).CreateUserDefinePerson(ctx, reqbody.(*CreateUserDefinePersonRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_RemoveUserDefinePerson_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &RemoveUserDefinePersonRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).RemoveUserDefinePerson(ctx, reqbody.(*RemoveUserDefinePersonRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_UpdateUserDefinePerson_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &UpdateUserDefinePersonRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).UpdateUserDefinePerson(ctx, reqbody.(*UpdateUserDefinePersonRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_DescribeUserDefinePerson_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &DescribeUserDefinePersonRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).DescribeUserDefinePerson(ctx, reqbody.(*DescribeUserDefinePersonRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_DescribeUserDefinePersonDetail_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &DescribeUserDefinePersonDetailRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).DescribeUserDefinePersonDetail(ctx, reqbody.(*DescribeUserDefinePersonDetailRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_DescribeUserDefinePersonWithFeature_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &DescribeUserDefinePersonWithFeatureRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).DescribeUserDefinePersonWithFeature(ctx, reqbody.(*DescribeUserDefinePersonWithFeatureRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_CreatePersonCategory_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &CreatePersonCategoryRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).CreatePersonCategory(ctx, reqbody.(*CreatePersonCategoryRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_RemovePersonCategory_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &RemovePersonCategoryRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).RemovePersonCategory(ctx, reqbody.(*RemovePersonCategoryRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_DescribePersonCategory_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &DescribePersonCategoryRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).DescribePersonCategory(ctx, reqbody.(*DescribePersonCategoryRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_RevertPerson_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &PersonRevertRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).RevertPerson(ctx, reqbody.(*PersonRevertRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserDefinePersonService_DescribeRevertTasks_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &DescribeRevertTasksRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserDefinePersonService).DescribeRevertTasks(ctx, reqbody.(*DescribeRevertTasksRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

// UserDefinePersonServer_ServiceDesc descriptor for server.RegisterService
var UserDefinePersonServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.media.personuserdefine.UserDefinePerson",
	HandlerType: ((*UserDefinePersonService)(nil)),
	Methods: []server.Method{
		{
			Name: "/CreateUserDefineGroup",
			Func: UserDefinePersonService_CreateUserDefineGroup_Handler,
		},
		{
			Name: "/DescribeUserDefineGroup",
			Func: UserDefinePersonService_DescribeUserDefineGroup_Handler,
		},
		{
			Name: "/RemoveUserDefineGroup",
			Func: UserDefinePersonService_RemoveUserDefineGroup_Handler,
		},
		{
			Name: "/UpdateUserDefineGroup",
			Func: UserDefinePersonService_UpdateUserDefineGroup_Handler,
		},
		{
			Name: "/CreateUserDefinePerson",
			Func: UserDefinePersonService_CreateUserDefinePerson_Handler,
		},
		{
			Name: "/RemoveUserDefinePerson",
			Func: UserDefinePersonService_RemoveUserDefinePerson_Handler,
		},
		{
			Name: "/UpdateUserDefinePerson",
			Func: UserDefinePersonService_UpdateUserDefinePerson_Handler,
		},
		{
			Name: "/DescribeUserDefinePerson",
			Func: UserDefinePersonService_DescribeUserDefinePerson_Handler,
		},
		{
			Name: "/DescribeUserDefinePersonDetail",
			Func: UserDefinePersonService_DescribeUserDefinePersonDetail_Handler,
		},
		{
			Name: "/DescribeUserDefinePersonWithFeature",
			Func: UserDefinePersonService_DescribeUserDefinePersonWithFeature_Handler,
		},
		{
			Name: "/CreatePersonCategory",
			Func: UserDefinePersonService_CreatePersonCategory_Handler,
		},
		{
			Name: "/RemovePersonCategory",
			Func: UserDefinePersonService_RemovePersonCategory_Handler,
		},
		{
			Name: "/DescribePersonCategory",
			Func: UserDefinePersonService_DescribePersonCategory_Handler,
		},
		{
			Name: "/RevertPerson",
			Func: UserDefinePersonService_RevertPerson_Handler,
		},
		{
			Name: "/DescribeRevertTasks",
			Func: UserDefinePersonService_DescribeRevertTasks_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/CreateUserDefineGroup",
			Func: UserDefinePersonService_CreateUserDefineGroup_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/DescribeUserDefineGroup",
			Func: UserDefinePersonService_DescribeUserDefineGroup_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/RemoveUserDefineGroup",
			Func: UserDefinePersonService_RemoveUserDefineGroup_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/UpdateUserDefineGroup",
			Func: UserDefinePersonService_UpdateUserDefineGroup_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/CreateUserDefinePerson",
			Func: UserDefinePersonService_CreateUserDefinePerson_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/RemoveUserDefinePerson",
			Func: UserDefinePersonService_RemoveUserDefinePerson_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/UpdateUserDefinePerson",
			Func: UserDefinePersonService_UpdateUserDefinePerson_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/DescribeUserDefinePerson",
			Func: UserDefinePersonService_DescribeUserDefinePerson_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/DescribeUserDefinePersonDetail",
			Func: UserDefinePersonService_DescribeUserDefinePersonDetail_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/DescribeUserDefinePersonWithFeature",
			Func: UserDefinePersonService_DescribeUserDefinePersonWithFeature_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/CreatePersonCategory",
			Func: UserDefinePersonService_CreatePersonCategory_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/RemovePersonCategory",
			Func: UserDefinePersonService_RemovePersonCategory_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/DescribePersonCategory",
			Func: UserDefinePersonService_DescribePersonCategory_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/RevertPerson",
			Func: UserDefinePersonService_RevertPerson_Handler,
		},
		{
			Name: "/trpc.media.personuserdefine.UserDefinePerson/DescribeRevertTasks",
			Func: UserDefinePersonService_DescribeRevertTasks_Handler,
		},
	},
}

// RegisterUserDefinePersonService register service
func RegisterUserDefinePersonService(s server.Service, svr UserDefinePersonService) {
	if err := s.Register(&UserDefinePersonServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("UserDefinePerson register error:%v", err))
	}

}

/* ************************************ Client Definition ************************************ */

// UserDefinePersonClientProxy defines service client proxy
type UserDefinePersonClientProxy interface {

	// CreateUserDefineGroup 自定义人物库
	CreateUserDefineGroup(ctx context.Context, req *CreateUserDefineGroupRequest, opts ...client.Option) (rsp *CreateUserDefineGroupResponse, err error) // @alias=/CreateUserDefineGroup

	DescribeUserDefineGroup(ctx context.Context, req *DescribeUserDefineGroupRequest, opts ...client.Option) (rsp *DescribeUserDefineGroupResponse, err error) // @alias=/DescribeUserDefineGroup

	RemoveUserDefineGroup(ctx context.Context, req *RemoveUserDefineGroupRequest, opts ...client.Option) (rsp *RemoveUserDefineGroupResponse, err error) // @alias=/RemoveUserDefineGroup

	UpdateUserDefineGroup(ctx context.Context, req *UpdateUserDefineGroupRequest, opts ...client.Option) (rsp *UpdateUserDefineGroupResponse, err error) // @alias=/UpdateUserDefineGroup

	// CreateUserDefinePerson 自定义人脸
	CreateUserDefinePerson(ctx context.Context, req *CreateUserDefinePersonRequest, opts ...client.Option) (rsp *CreateUserDefinePersonResponse, err error) // @alias=/CreateUserDefinePerson

	RemoveUserDefinePerson(ctx context.Context, req *RemoveUserDefinePersonRequest, opts ...client.Option) (rsp *RemoveUserDefinePersonResponse, err error) // @alias=/RemoveUserDefinePerson

	UpdateUserDefinePerson(ctx context.Context, req *UpdateUserDefinePersonRequest, opts ...client.Option) (rsp *UpdateUserDefinePersonResponse, err error) // @alias=/UpdateUserDefinePerson

	DescribeUserDefinePerson(ctx context.Context, req *DescribeUserDefinePersonRequest, opts ...client.Option) (rsp *DescribeUserDefinePersonResponse, err error) // @alias=/DescribeUserDefinePerson

	DescribeUserDefinePersonDetail(ctx context.Context, req *DescribeUserDefinePersonDetailRequest, opts ...client.Option) (rsp *DescribeUserDefinePersonDetailResponse, err error) // @alias=/DescribeUserDefinePersonDetail

	DescribeUserDefinePersonWithFeature(ctx context.Context, req *DescribeUserDefinePersonWithFeatureRequest, opts ...client.Option) (rsp *DescribeUserDefinePersonWithFeatureResponse, err error) // @alias=/DescribeUserDefinePersonWithFeature

	// CreatePersonCategory 人物分类
	CreatePersonCategory(ctx context.Context, req *CreatePersonCategoryRequest, opts ...client.Option) (rsp *CreatePersonCategoryResponse, err error) // @alias=/CreatePersonCategory

	RemovePersonCategory(ctx context.Context, req *RemovePersonCategoryRequest, opts ...client.Option) (rsp *RemovePersonCategoryResponse, err error) // @alias=/RemovePersonCategory

	DescribePersonCategory(ctx context.Context, req *DescribePersonCategoryRequest, opts ...client.Option) (rsp *DescribePersonCategoryResponse, err error) // @alias=/DescribePersonCategory

	// RevertPerson 翻库
	RevertPerson(ctx context.Context, req *PersonRevertRequest, opts ...client.Option) (rsp *PersonRevertResponse, err error) // @alias=/RevertPerson

	DescribeRevertTasks(ctx context.Context, req *DescribeRevertTasksRequest, opts ...client.Option) (rsp *DescribeRevertTasksResponse, err error) // @alias=/DescribeRevertTasks
}

type UserDefinePersonClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewUserDefinePersonClientProxy = func(opts ...client.Option) UserDefinePersonClientProxy {
	return &UserDefinePersonClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *UserDefinePersonClientProxyImpl) CreateUserDefineGroup(ctx context.Context, req *CreateUserDefineGroupRequest, opts ...client.Option) (*CreateUserDefineGroupResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/CreateUserDefineGroup")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("CreateUserDefineGroup")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &CreateUserDefineGroupResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) DescribeUserDefineGroup(ctx context.Context, req *DescribeUserDefineGroupRequest, opts ...client.Option) (*DescribeUserDefineGroupResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/DescribeUserDefineGroup")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("DescribeUserDefineGroup")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &DescribeUserDefineGroupResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) RemoveUserDefineGroup(ctx context.Context, req *RemoveUserDefineGroupRequest, opts ...client.Option) (*RemoveUserDefineGroupResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/RemoveUserDefineGroup")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("RemoveUserDefineGroup")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &RemoveUserDefineGroupResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) UpdateUserDefineGroup(ctx context.Context, req *UpdateUserDefineGroupRequest, opts ...client.Option) (*UpdateUserDefineGroupResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/UpdateUserDefineGroup")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("UpdateUserDefineGroup")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &UpdateUserDefineGroupResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) CreateUserDefinePerson(ctx context.Context, req *CreateUserDefinePersonRequest, opts ...client.Option) (*CreateUserDefinePersonResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/CreateUserDefinePerson")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("CreateUserDefinePerson")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &CreateUserDefinePersonResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) RemoveUserDefinePerson(ctx context.Context, req *RemoveUserDefinePersonRequest, opts ...client.Option) (*RemoveUserDefinePersonResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/RemoveUserDefinePerson")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("RemoveUserDefinePerson")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &RemoveUserDefinePersonResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) UpdateUserDefinePerson(ctx context.Context, req *UpdateUserDefinePersonRequest, opts ...client.Option) (*UpdateUserDefinePersonResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/UpdateUserDefinePerson")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("UpdateUserDefinePerson")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &UpdateUserDefinePersonResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) DescribeUserDefinePerson(ctx context.Context, req *DescribeUserDefinePersonRequest, opts ...client.Option) (*DescribeUserDefinePersonResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/DescribeUserDefinePerson")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("DescribeUserDefinePerson")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &DescribeUserDefinePersonResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) DescribeUserDefinePersonDetail(ctx context.Context, req *DescribeUserDefinePersonDetailRequest, opts ...client.Option) (*DescribeUserDefinePersonDetailResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/DescribeUserDefinePersonDetail")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("DescribeUserDefinePersonDetail")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &DescribeUserDefinePersonDetailResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) DescribeUserDefinePersonWithFeature(ctx context.Context, req *DescribeUserDefinePersonWithFeatureRequest, opts ...client.Option) (*DescribeUserDefinePersonWithFeatureResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/DescribeUserDefinePersonWithFeature")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("DescribeUserDefinePersonWithFeature")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &DescribeUserDefinePersonWithFeatureResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) CreatePersonCategory(ctx context.Context, req *CreatePersonCategoryRequest, opts ...client.Option) (*CreatePersonCategoryResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/CreatePersonCategory")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("CreatePersonCategory")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &CreatePersonCategoryResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) RemovePersonCategory(ctx context.Context, req *RemovePersonCategoryRequest, opts ...client.Option) (*RemovePersonCategoryResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/RemovePersonCategory")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("RemovePersonCategory")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &RemovePersonCategoryResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) DescribePersonCategory(ctx context.Context, req *DescribePersonCategoryRequest, opts ...client.Option) (*DescribePersonCategoryResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/DescribePersonCategory")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("DescribePersonCategory")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &DescribePersonCategoryResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) RevertPerson(ctx context.Context, req *PersonRevertRequest, opts ...client.Option) (*PersonRevertResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/RevertPerson")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("RevertPerson")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &PersonRevertResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserDefinePersonClientProxyImpl) DescribeRevertTasks(ctx context.Context, req *DescribeRevertTasksRequest, opts ...client.Option) (*DescribeRevertTasksResponse, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/DescribeRevertTasks")
	msg.WithCalleeServiceName(UserDefinePersonServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("personuserdefine")
	msg.WithCalleeService("UserDefinePerson")
	msg.WithCalleeMethod("DescribeRevertTasks")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &DescribeRevertTasksResponse{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}