package filetransfer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Tencent-Ti/ti-sign-go/tisign"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// CompleteMultipartUploadAsyncResponse ...
type CompleteMultipartUploadAsyncResponse struct {
	Response struct {
		CommonResponse
		TaskId string `json:"TaskId"`
	} `json:"Response"`
}

// CompleteMultipartUploadAsync 异步完成文件上传
func (c Client) CompleteMultipartUploadAsync(ctx context.Context, key, bucket, uploadID string) (
	taskId string, err error) {

	action := "CompleteMultipartUploadAsync"
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
	req := map[string]interface{}{
		"CompleteMultipartUpload": nil,
	}

	uri := fmt.Sprintf("http://%s/FileManager/CompleteMultipartUploadAsync?useJson=true&Bucket=%s&Key=%s&uploadId=%s",
		c.opt.GetEndpoint(), bucket, key, uploadID)
	ts := tisign.NewTiSign(headerContent, c.opt.SecretId, c.opt.SecretKey)
	header, _ := ts.CreateSignatureInfo()

	maxTry := 3
	rsp := &CompleteMultipartUploadAsyncResponse{}
	timeSleep := 50 * time.Millisecond
	for i := 0; i < maxTry; i++ {
		err = client.JsonPost(ctx, uri, header, req, &rsp.Response)
		if rsp.Response.Error != nil {
			bys, _ := json.Marshal(rsp)
			err = errors.New("Response error: " + string(bys))
		}
		if err == nil {
			break
		}
		time.Sleep(timeSleep)
		timeSleep *= 2
		fmt.Println("Complete try ", i+1, " error: ", err.Error())
	}
	if err != nil {
		return "", err
	}
	return rsp.Response.TaskId, nil
}
