package filetransfer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Tencent-Ti/ti-sign-go/tisign"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// CreateMultipartUploadResponse ...
type CreateMultipartUploadResponse struct {
	Response struct {
		CommonResponse
		InitiateMultipartUploadResult struct {
			UploadId   string `json:"UploadId"`
			Bucket     string `json:"Bucket"`
			Key        string `json:"Key"`
			CreateTime string `json:"CreateTime"`
			UpdateTime string `json:"UpdateTime"`
			ExpireTime string `json:"ExpireTime"`
		} `json:"InitiateMultipartUploadResult"`
	} `json:"Response"`
}

// CreateMultipartUpload 创建分片上传任务
func (c Client) CreateMultipartUpload(ctx context.Context, bucket, key string) (uploadId string, err error) {
	action := "CreateMultipartUpload"
	service := "app-cdn4aowk"
	version := "2021-02-26"
	headerContent := tisign.HttpHeaderContent{
		XTCAction:   action,                                       // 请求接口
		XTCService:  service,                                      // 接口所属服务名
		XTCVersion:  version,                                      // 接口版本
		ContentType: "application/json",                           // http请求的content-type, 当前网关只支持: application/json  multipart/form-data
		HttpMethod:  "POST",                                       // http请求方法，当前网关只支持: POST GET
		Host:        fmt.Sprintf("%s:%d", c.opt.Host, c.opt.Port), // 访问网关的host
	}

	req := map[string]interface{}{}
	rsp := &CreateMultipartUploadResponse{}
	uri := fmt.Sprintf("http://%s:%d/FileManager/CreateMultipartUpload?useJson=true&Bucket=%s&Key=%s",
		c.opt.Host, c.opt.Port, bucket, key)
	ts := tisign.NewTiSign(headerContent, c.opt.SecretId, c.opt.SecretKey)
	header, _ := ts.CreateSignatureInfo()
	err = client.JsonPost(ctx, uri, header, req, &rsp.Response)
	if err != nil {
		return uploadId, fmt.Errorf("CreateMultipartUpload failed, err: %v", err)
	}
	if rsp.Response.Error != nil {
		respBody, _ := json.Marshal(rsp)
		return uploadId, fmt.Errorf("CreateMultipartUpload response failed, response: %s", string(respBody))
	}
	return rsp.Response.InitiateMultipartUploadResult.UploadId, err
}
