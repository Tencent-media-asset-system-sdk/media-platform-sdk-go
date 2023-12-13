package filetransfer

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Tencent-Ti/ti-sign-go/tisign"
)

// UploadPartResponse ...
type UploadPartResponse struct {
	Response struct {
		CommonResponse
		ETag string `json:"ETag"`
	} `json:"Response"`
}

// JsonPut http put 请求
func JsonPut(ctx context.Context, uri string, header map[string]string, req, rsp interface{}) (err error) {
	body, ok := req.([]byte)
	if !ok {
		body, _ = json.Marshal(req)
	}
	reqCtx, err := http.NewRequestWithContext(ctx, http.MethodPut, uri, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("JsonPut make NewRequest error: %v", err)
	}
	if header == nil {
		reqCtx.Header.Set("Content-Type", "application/octet-stream")
	} else {
		for head, value := range header {
			reqCtx.Header.Set(head, value)
		}
	}
	response, err := http.DefaultClient.Do(reqCtx)
	if err != nil {
		return fmt.Errorf("JsonPut do request error: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("JsonPut response status error, statuscode: %d, reason: %s",
			response.StatusCode, response.Status)
	}
	data, _ := ioutil.ReadAll(response.Body)
	// log.Printf("JsonPut req: %v, rsp: %v \n", string(body), string(data))
	if rsp != nil {
		err = json.Unmarshal(data, rsp)
		if err != nil {
			return fmt.Errorf("JsonPut response protocol error: %v, data: %s", err, string(data))
		}
	}
	return nil
}

// UploadPart 上传分片
func (c Client) UploadPart(ctx context.Context, bucket, key, uploadID string,
	partNumber int, filebuf []byte) (etag string, err error) {
	h := md5.New()
	h.Write(filebuf)
	md5sum := hex.EncodeToString(h.Sum(nil))

	rsp := &UploadPartResponse{}
	uri := fmt.Sprintf(
		"http://%s:%d/FileManager/UploadPart?useJson=true&Bucket=%s&Key=%s&uploadId=%s&partNumber=%d&Content-MD5=%s",
		c.opt.Host, c.opt.Port, bucket, url.QueryEscape(key), uploadID, partNumber, md5sum)
	headerContent := tisign.HttpHeaderContent{
		XTCAction:   "UploadPart",               // 请求接口
		XTCService:  "app-cdn4aowk",             // 接口所属服务名
		XTCVersion:  "2021-02-26",               // 接口版本
		ContentType: "application/octet-stream", // http请求的content-type, 当前网关只支持: application/json  multipart/form-data
		HttpMethod:  "PUT",                      // http请求方法，当前网关只支持: POST GET
		Host:        c.opt.Host,                 // 访问网关的host
	}
	ts := tisign.NewTiSign(headerContent, c.opt.SecretId, c.opt.SecretKey)
	header, _ := ts.CreateSignatureInfo()
	err = JsonPut(ctx, uri, header, filebuf, &rsp.Response)
	if err != nil {
		return key, fmt.Errorf("[%s]UploadPart failed, err: %v", uri, err)
	}

	if rsp.Response.Error != nil {
		respBody, _ := json.Marshal(rsp)
		return key, fmt.Errorf("[%s]UploadPart response failed, response: %s", uri, string(respBody))
	}
	return rsp.Response.ETag, err
}
