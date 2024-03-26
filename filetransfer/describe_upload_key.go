package filetransfer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Tencent-Ti/ti-sign-go/tisign"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// Error ...
type Error struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

// CommonResponse ...
type CommonResponse struct {
	RequestId string `json:"RequestId"`
	Error     *Error `json:"Error"`
}

// DescribeUploadKeyResponse ...
type DescribeUploadKeyResponse struct {
	Response struct {
		CommonResponse
		Key string `json:"Key"`
	} `json:"Response"`
}

// DescribeUploadKey 获取上传文件 key
func (c Client) DescribeUploadKey(ctx context.Context, filename string, fileSize uint64) (key string, err error) {
	action := "DescribeUploadKey"
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
	req := map[string]interface{}{
		"DescribeUploadKey": map[string]interface{}{
			"FileName": filename,
			"Size":     strconv.FormatUint(fileSize, 10),
		},
	}

	uri := fmt.Sprintf("http://%s:%d/FileManager/DescribeUploadKey?useJson=true", c.opt.Host, c.opt.Port)
	ts := tisign.NewTiSign(headerContent, c.opt.SecretId, c.opt.SecretKey)
	header, _ := ts.CreateSignatureInfo()
	maxTry := 3
	timeSleep := 50 * time.Millisecond
	rsp := &DescribeUploadKeyResponse{}
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
	}
	return rsp.Response.Key, err
}
