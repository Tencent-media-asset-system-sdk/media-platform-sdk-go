package filetransfer

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/Tencent-Ti/ti-sign-go/tisign"
)

// PutObjectResponse ...
type PutObjectResponse struct {
	Response struct {
		CommonResponse
		Key  string `json:"Key"`
		ETag string `json:"ETag"`
	} `json:"Response"`
}

// PutObject 直接上传
func (c Client) PutObject(ctx context.Context, key string, filebuf []byte) (
	r *PutObjectResponse, err error) {
	h := md5.New()
	h.Write(filebuf)
	md5sum := hex.EncodeToString(h.Sum(nil))
	canonicalQueryString := fmt.Sprintf("useJson=true&Bucket=&Key=%s&Content-MD5=%s", key, md5sum)
	uri := fmt.Sprintf("http://%s:%d/FileManager/PutObject?%s", c.opt.Host, c.opt.Port, canonicalQueryString)
	headerContent := tisign.HttpHeaderContent{
		XTCAction:   "PutObject",                                  // 请求接口
		XTCService:  "app-cdn4aowk",                               // 接口所属服务名
		XTCVersion:  "2021-02-26",                                 // 接口版本
		ContentType: "application/octet-stream",                   // http请求的content-type, 当前网关只支持: application/json  multipart/form-data
		HttpMethod:  "PUT",                                        // http请求方法，当前网关只支持: POST GET
		Host:        fmt.Sprintf("%s:%d", c.opt.Host, c.opt.Port), // 访问网关的host
	}
	ts := tisign.NewTiSign(headerContent, c.opt.SecretId, c.opt.SecretKey)
	header, _ := ts.CreateSignatureInfo()
	rsp := &PutObjectResponse{}
	err = JsonPut(ctx, uri, header, filebuf, &rsp.Response)
	if err != nil {
		return nil, fmt.Errorf("[%s]PutObject make NewRequest error: %v", uri, err)
	}
	if rsp.Response.Error != nil {
		respBody, _ := json.Marshal(rsp)
		return nil, fmt.Errorf("[%s]PutObject response error %s", uri, string(respBody))
	}
	return rsp, nil
}
