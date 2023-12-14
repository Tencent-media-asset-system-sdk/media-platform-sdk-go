// Code generated by trpc-go/trpc-go-cmdline v2.2.18. DO NOT EDIT.
// source: custom_feature.proto

package customfeature

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

// CustomFeatureService defines service.
type CustomFeatureService interface {
	// DescribeTextCategories 自定义文本标签相关接口
	DescribeTextCategories(ctx context.Context, req *DescribeTextCategoriesRequest, rsp *DescribeTextCategoriesResponse) (err error) // @alias=/DescribeTextCategories

	CreateCustomText(ctx context.Context, req *CreateCustomTextRequest, rsp *CreateCustomTextResponse) (err error) // @alias=/CreateCustomText

	ModifyCustomText(ctx context.Context, req *ModifyCustomTextRequest, rsp *ModifyCustomTextResponse) (err error) // @alias=/ModifyCustomText

	DescribeCustomTexts(ctx context.Context, req *DescribeCustomTextsRequest, rsp *DescribeCustomTextsResponse) (err error) // @alias=/DescribeCustomTexts

	DescribeLastUpdateTime(ctx context.Context, req *DescribeLastUpdateTimeRequest, rsp *DescribeLastUpdateTimeResponse) (err error) // @alias=/DescribeLastUpdateTime

	DeleteCustomText(ctx context.Context, req *DeleteCustomTextRequest, rsp *DeleteCustomTextResponse) (err error) // @alias=/DeleteCustomText
	// CreateFeature 特征相关接口
	CreateFeature(ctx context.Context, req *CreateFeatureRequest, rsp *CreateFeatureResponse) (err error) // @alias=/CreateFeature

	ModifyFeature(ctx context.Context, req *ModifyFeatureRequest, rsp *ModifyFeatureResponse) (err error) // @alias=/ModifyFeature

	DescribeFeatures(ctx context.Context, req *DescribeFeaturesRequest, rsp *DescribeFeaturesResponse) (err error) // @alias=/DescribeFeatures

	DeleteFeature(ctx context.Context, req *DeleteFeatureRequest, rsp *DeleteFeatureResponse) (err error) // @alias=/DeleteFeature
}

func CustomFeatureService_DescribeTextCategories_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeTextCategoriesRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).DescribeTextCategories(ctx, reqbody.(*DescribeTextCategoriesRequest), rspbody.(*DescribeTextCategoriesResponse))
	}

	rsp := &DescribeTextCategoriesResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_CreateCustomText_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &CreateCustomTextRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).CreateCustomText(ctx, reqbody.(*CreateCustomTextRequest), rspbody.(*CreateCustomTextResponse))
	}

	rsp := &CreateCustomTextResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_ModifyCustomText_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &ModifyCustomTextRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).ModifyCustomText(ctx, reqbody.(*ModifyCustomTextRequest), rspbody.(*ModifyCustomTextResponse))
	}

	rsp := &ModifyCustomTextResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_DescribeCustomTexts_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeCustomTextsRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).DescribeCustomTexts(ctx, reqbody.(*DescribeCustomTextsRequest), rspbody.(*DescribeCustomTextsResponse))
	}

	rsp := &DescribeCustomTextsResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_DescribeLastUpdateTime_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeLastUpdateTimeRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).DescribeLastUpdateTime(ctx, reqbody.(*DescribeLastUpdateTimeRequest), rspbody.(*DescribeLastUpdateTimeResponse))
	}

	rsp := &DescribeLastUpdateTimeResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_DeleteCustomText_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DeleteCustomTextRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).DeleteCustomText(ctx, reqbody.(*DeleteCustomTextRequest), rspbody.(*DeleteCustomTextResponse))
	}

	rsp := &DeleteCustomTextResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_CreateFeature_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &CreateFeatureRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).CreateFeature(ctx, reqbody.(*CreateFeatureRequest), rspbody.(*CreateFeatureResponse))
	}

	rsp := &CreateFeatureResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_ModifyFeature_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &ModifyFeatureRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).ModifyFeature(ctx, reqbody.(*ModifyFeatureRequest), rspbody.(*ModifyFeatureResponse))
	}

	rsp := &ModifyFeatureResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_DescribeFeatures_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DescribeFeaturesRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).DescribeFeatures(ctx, reqbody.(*DescribeFeaturesRequest), rspbody.(*DescribeFeaturesResponse))
	}

	rsp := &DescribeFeaturesResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func CustomFeatureService_DeleteFeature_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &DeleteFeatureRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(CustomFeatureService).DeleteFeature(ctx, reqbody.(*DeleteFeatureRequest), rspbody.(*DeleteFeatureResponse))
	}

	rsp := &DeleteFeatureResponse{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// CustomFeatureServer_ServiceDesc descriptor for server.RegisterService.
var CustomFeatureServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.media.customfeature.CustomFeature",
	HandlerType: ((*CustomFeatureService)(nil)),
	Methods: []server.Method{
		{
			Name: "/DescribeTextCategories",
			Func: CustomFeatureService_DescribeTextCategories_Handler,
		},
		{
			Name: "/CreateCustomText",
			Func: CustomFeatureService_CreateCustomText_Handler,
		},
		{
			Name: "/ModifyCustomText",
			Func: CustomFeatureService_ModifyCustomText_Handler,
		},
		{
			Name: "/DescribeCustomTexts",
			Func: CustomFeatureService_DescribeCustomTexts_Handler,
		},
		{
			Name: "/DescribeLastUpdateTime",
			Func: CustomFeatureService_DescribeLastUpdateTime_Handler,
		},
		{
			Name: "/DeleteCustomText",
			Func: CustomFeatureService_DeleteCustomText_Handler,
		},
		{
			Name: "/CreateFeature",
			Func: CustomFeatureService_CreateFeature_Handler,
		},
		{
			Name: "/ModifyFeature",
			Func: CustomFeatureService_ModifyFeature_Handler,
		},
		{
			Name: "/DescribeFeatures",
			Func: CustomFeatureService_DescribeFeatures_Handler,
		},
		{
			Name: "/DeleteFeature",
			Func: CustomFeatureService_DeleteFeature_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/DescribeTextCategories",
			Func: CustomFeatureService_DescribeTextCategories_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/CreateCustomText",
			Func: CustomFeatureService_CreateCustomText_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/ModifyCustomText",
			Func: CustomFeatureService_ModifyCustomText_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/DescribeCustomTexts",
			Func: CustomFeatureService_DescribeCustomTexts_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/DescribeLastUpdateTime",
			Func: CustomFeatureService_DescribeLastUpdateTime_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/DeleteCustomText",
			Func: CustomFeatureService_DeleteCustomText_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/CreateFeature",
			Func: CustomFeatureService_CreateFeature_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/ModifyFeature",
			Func: CustomFeatureService_ModifyFeature_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/DescribeFeatures",
			Func: CustomFeatureService_DescribeFeatures_Handler,
		},
		{
			Name: "/trpc.media.customfeature.CustomFeature/DeleteFeature",
			Func: CustomFeatureService_DeleteFeature_Handler,
		},
	},
}

// RegisterCustomFeatureService registers service.
func RegisterCustomFeatureService(s server.Service, svr CustomFeatureService) {
	if err := s.Register(&CustomFeatureServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("CustomFeature register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedCustomFeature struct{}

// DescribeTextCategories 自定义文本标签相关接口
func (s *UnimplementedCustomFeature) DescribeTextCategories(ctx context.Context, req *DescribeTextCategoriesRequest, rsp *DescribeTextCategoriesResponse) error {
	return errors.New("rpc DescribeTextCategories of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) CreateCustomText(ctx context.Context, req *CreateCustomTextRequest, rsp *CreateCustomTextResponse) error {
	return errors.New("rpc CreateCustomText of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) ModifyCustomText(ctx context.Context, req *ModifyCustomTextRequest, rsp *ModifyCustomTextResponse) error {
	return errors.New("rpc ModifyCustomText of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) DescribeCustomTexts(ctx context.Context, req *DescribeCustomTextsRequest, rsp *DescribeCustomTextsResponse) error {
	return errors.New("rpc DescribeCustomTexts of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) DescribeLastUpdateTime(ctx context.Context, req *DescribeLastUpdateTimeRequest, rsp *DescribeLastUpdateTimeResponse) error {
	return errors.New("rpc DescribeLastUpdateTime of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) DeleteCustomText(ctx context.Context, req *DeleteCustomTextRequest, rsp *DeleteCustomTextResponse) error {
	return errors.New("rpc DeleteCustomText of service CustomFeature is not implemented")
}

// CreateFeature 特征相关接口
func (s *UnimplementedCustomFeature) CreateFeature(ctx context.Context, req *CreateFeatureRequest, rsp *CreateFeatureResponse) error {
	return errors.New("rpc CreateFeature of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) ModifyFeature(ctx context.Context, req *ModifyFeatureRequest, rsp *ModifyFeatureResponse) error {
	return errors.New("rpc ModifyFeature of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) DescribeFeatures(ctx context.Context, req *DescribeFeaturesRequest, rsp *DescribeFeaturesResponse) error {
	return errors.New("rpc DescribeFeatures of service CustomFeature is not implemented")
}
func (s *UnimplementedCustomFeature) DeleteFeature(ctx context.Context, req *DeleteFeatureRequest, rsp *DeleteFeatureResponse) error {
	return errors.New("rpc DeleteFeature of service CustomFeature is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// CustomFeatureClientProxy defines service client proxy
type CustomFeatureClientProxy interface {
	// DescribeTextCategories 自定义文本标签相关接口
	DescribeTextCategories(ctx context.Context, req *DescribeTextCategoriesRequest, opts ...client.Option) (rsp *DescribeTextCategoriesResponse, err error) // @alias=/DescribeTextCategories

	CreateCustomText(ctx context.Context, req *CreateCustomTextRequest, opts ...client.Option) (rsp *CreateCustomTextResponse, err error) // @alias=/CreateCustomText

	ModifyCustomText(ctx context.Context, req *ModifyCustomTextRequest, opts ...client.Option) (rsp *ModifyCustomTextResponse, err error) // @alias=/ModifyCustomText

	DescribeCustomTexts(ctx context.Context, req *DescribeCustomTextsRequest, opts ...client.Option) (rsp *DescribeCustomTextsResponse, err error) // @alias=/DescribeCustomTexts

	DescribeLastUpdateTime(ctx context.Context, req *DescribeLastUpdateTimeRequest, opts ...client.Option) (rsp *DescribeLastUpdateTimeResponse, err error) // @alias=/DescribeLastUpdateTime

	DeleteCustomText(ctx context.Context, req *DeleteCustomTextRequest, opts ...client.Option) (rsp *DeleteCustomTextResponse, err error) // @alias=/DeleteCustomText
	// CreateFeature 特征相关接口
	CreateFeature(ctx context.Context, req *CreateFeatureRequest, opts ...client.Option) (rsp *CreateFeatureResponse, err error) // @alias=/CreateFeature

	ModifyFeature(ctx context.Context, req *ModifyFeatureRequest, opts ...client.Option) (rsp *ModifyFeatureResponse, err error) // @alias=/ModifyFeature

	DescribeFeatures(ctx context.Context, req *DescribeFeaturesRequest, opts ...client.Option) (rsp *DescribeFeaturesResponse, err error) // @alias=/DescribeFeatures

	DeleteFeature(ctx context.Context, req *DeleteFeatureRequest, opts ...client.Option) (rsp *DeleteFeatureResponse, err error) // @alias=/DeleteFeature
}

type CustomFeatureClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewCustomFeatureClientProxy = func(opts ...client.Option) CustomFeatureClientProxy {
	return &CustomFeatureClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *CustomFeatureClientProxyImpl) DescribeTextCategories(ctx context.Context, req *DescribeTextCategoriesRequest, opts ...client.Option) (*DescribeTextCategoriesResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeTextCategories")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("DescribeTextCategories")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeTextCategoriesResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) CreateCustomText(ctx context.Context, req *CreateCustomTextRequest, opts ...client.Option) (*CreateCustomTextResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/CreateCustomText")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("CreateCustomText")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CreateCustomTextResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) ModifyCustomText(ctx context.Context, req *ModifyCustomTextRequest, opts ...client.Option) (*ModifyCustomTextResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/ModifyCustomText")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("ModifyCustomText")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &ModifyCustomTextResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) DescribeCustomTexts(ctx context.Context, req *DescribeCustomTextsRequest, opts ...client.Option) (*DescribeCustomTextsResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeCustomTexts")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("DescribeCustomTexts")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeCustomTextsResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) DescribeLastUpdateTime(ctx context.Context, req *DescribeLastUpdateTimeRequest, opts ...client.Option) (*DescribeLastUpdateTimeResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeLastUpdateTime")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("DescribeLastUpdateTime")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeLastUpdateTimeResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) DeleteCustomText(ctx context.Context, req *DeleteCustomTextRequest, opts ...client.Option) (*DeleteCustomTextResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DeleteCustomText")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("DeleteCustomText")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DeleteCustomTextResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) CreateFeature(ctx context.Context, req *CreateFeatureRequest, opts ...client.Option) (*CreateFeatureResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/CreateFeature")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("CreateFeature")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CreateFeatureResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) ModifyFeature(ctx context.Context, req *ModifyFeatureRequest, opts ...client.Option) (*ModifyFeatureResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/ModifyFeature")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("ModifyFeature")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &ModifyFeatureResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) DescribeFeatures(ctx context.Context, req *DescribeFeaturesRequest, opts ...client.Option) (*DescribeFeaturesResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DescribeFeatures")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("DescribeFeatures")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DescribeFeaturesResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *CustomFeatureClientProxyImpl) DeleteFeature(ctx context.Context, req *DeleteFeatureRequest, opts ...client.Option) (*DeleteFeatureResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/DeleteFeature")
	msg.WithCalleeServiceName(CustomFeatureServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("media")
	msg.WithCalleeServer("customfeature")
	msg.WithCalleeService("CustomFeature")
	msg.WithCalleeMethod("DeleteFeature")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &DeleteFeatureResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
