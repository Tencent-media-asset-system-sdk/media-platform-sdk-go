package filetransfer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Tencent-Ti/ti-sign-go/tisign"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// CreateDeleteFileTaskRequest ...
type CreateDeleteFileTaskRequest struct {
	FileKeys []string `json:"FileKeys"`
	Async    bool     `json:"Async"`
}

// CreateDeleteFileTaskResponse ...
type CreateDeleteFileTaskResponse struct {
	Response struct {
		CommonResponse
		DeleteFileTask struct {
			TaskId         string `json:"TaskId"`
			Async          bool   `json:"Async"`
			CreateTime     string `json:"CreateTime"`
			FinishTime     string `json:"FinishTime"`
			Total          int    `json:"Total"`
			ProcessedCount int    `json:"ProcessedCount"`
			Status         string `json:"Status"`
			FailedFileKeys []struct {
				Key        string `json:"Key"`
				Reason     string `json:"Reason"`
				FailedTime string `json:"FailedTime"`
			} `json:"FailedFileKeys"`
		} `json:"DeleteFileTask"`
	} `json:"Response"`
}

// NewCreateDeleteFileTaskRequest ...
func NewCreateDeleteFileTaskRequest(keys []string, isAsync bool) *CreateDeleteFileTaskRequest {
	r := &CreateDeleteFileTaskRequest{}
	r.FileKeys = keys
	r.Async = isAsync
	return r
}

// CreateDeleteFileTask 删除媒体
func (c Client) CreateDeleteFileTask(ctx context.Context,
	keys []string, isAsync bool) (failedKeys, failedReadon []string, err error) {

	action := "CreateDeleteFileTask"
	service := "app-cdn4aowk"
	version := "2021-02-26"
	headerContent := tisign.HttpHeaderContent{
		XTCAction:   action,              // 请求接口
		XTCService:  service,             // 接口所属服务名
		XTCVersion:  version,             // 接口版本
		ContentType: "application/json",  // http请求的content-type, 当前网关只支持: application/json  multipart/form-data
		HttpMethod:  "POST",              // http请求方法，当前网关只支持: POST GET
		Host:        c.opt.GetEndpoint(), // 访问网关的host
	}

	req := NewCreateDeleteFileTaskRequest(keys, isAsync)
	rsp := &CreateDeleteFileTaskResponse{}
	uri := fmt.Sprintf("http://%s/FileManager/CreateDeleteFileTask?useJson=true", c.opt.GetEndpoint())
	ts := tisign.NewTiSign(headerContent, c.opt.SecretId, c.opt.SecretKey)
	header, _ := ts.CreateSignatureInfo()
	err = client.JsonPost(ctx, uri, header, req, &rsp.Response)
	if rsp.Response.Error != nil {
		bs, _ := json.Marshal(rsp)
		return nil, nil, fmt.Errorf("CreateDeleteFileTask response failed, response: %s", string(bs))
	}
	for _, item := range rsp.Response.DeleteFileTask.FailedFileKeys {
		failedKeys = append(failedKeys, item.Key)
		failedReadon = append(failedReadon, item.Reason)
	}
	return failedKeys, failedReadon, err
}
