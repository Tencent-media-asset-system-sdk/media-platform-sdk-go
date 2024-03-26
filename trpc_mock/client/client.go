package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Tencent-Ti/ti-sign-go/tisign"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/codec"
	"github.com/tidwall/gjson"
)

// JsonPost http json post 请求
func JsonPost(ctx context.Context, uri string, header map[string]string, req, rsp interface{}) (err error) {
	body, _ := json.Marshal(req)
	reqCtx, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("HttpPost make NewRequest error: %v", err)
	}
	if header == nil {
		reqCtx.Header.Set("Content-Type", "application/json")
	} else {
		for head, value := range header {
			reqCtx.Header.Set(head, value)
		}
	}
	response, err := http.DefaultClient.Do(reqCtx)
	if err != nil {
		return fmt.Errorf("HttpPost do request error: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HttpPost response status error, statuscode: %d, reason: %s",
			response.StatusCode, response.Status)
	}
	data, _ := ioutil.ReadAll(response.Body)
	// log.Printf("JsonPost req: %v, rsp: %v \n", string(body), string(data))
	g := gjson.ParseBytes(data)
	if g.Get("Response.Error").Exists() {
		return fmt.Errorf("response error: %v", string(data))
	}
	if rsp != nil {
		err = json.Unmarshal([]byte(g.Get("Response").String()), rsp)
		if err != nil {
			return fmt.Errorf("HttpPost response protocol error: %v, data: %s", err, string(data))
		}
	}
	ctx = context.WithValue(ctx, codec.ContextKeyRequestId, g.Get("Response.RequestId").String())
	return nil
}

type Client struct {
}

var DefaultClient Client = Client{}

type Option struct {
	Host                      string
	Port                      int
	SecretId, SecretKey       string
	TIBusinessId, TIProjectId uint32
	Inside                    bool
	Uin, SubAccountUin        string
}

const (
	defaultservice = "fusion-media-service"
	defaultversion = "2022-03-02"
	defaultgateway = "gateway"
)

var serviceMap = map[string]string{}
var versionMap = map[string]string{}
var gatewayMap = map[string]string{}

func (Client) Invoke(ctx context.Context, reqBody interface{}, rspBody interface{}, opt ...Option) (err error) {
	if len(opt) == 0 {
		return fmt.Errorf("必须要 host option")
	}
	if opt[0].Port == 0 {
		opt[0].Port = 80
	}
	// 走网关
	msg := ctx.Value(codec.ContextKeyMessage).(codec.Msg)
	action := msg.GetAction()
	service := defaultservice
	version := defaultversion
	gateway := defaultgateway
	if v, ok := serviceMap[action]; ok {
		service = v
	}
	if v, ok := versionMap[action]; ok {
		version = v
	}
	if v, ok := gatewayMap[action]; ok {
		gateway = v
	}
	var header map[string]string
	var uri string
	var m map[string]interface{}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("json Marshal req error: %v", err)
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return fmt.Errorf("json UnMarshal req error: %v", err)
	}
	if opt[0].Inside {
		uri = fmt.Sprintf("http://%s:%d/%s", opt[0].Host, opt[0].Port, action)
		m["Uin"] = opt[0].Uin
		m["SubAccountUin"] = opt[0].SubAccountUin
	} else {
		headerContent := tisign.HttpHeaderContent{
			XTCAction:   action,                                         // 请求接口
			XTCService:  service,                                        // 接口所属服务名
			XTCVersion:  version,                                        // 接口版本
			ContentType: "application/json",                             // http请求的content-type, 当前网关只支持: application/json  multipart/form-data
			HttpMethod:  "POST",                                         // http请求方法，当前网关只支持: POST GET
			Host:        fmt.Sprintf("%s:%d", opt[0].Host, opt[0].Port), // 访问网关的host
		}
		m["TIProjectId"] = opt[0].TIProjectId
		m["TIBusinessId"] = opt[0].TIBusinessId
		uri = fmt.Sprintf("http://%s:%d/%s", opt[0].Host, opt[0].Port, gateway)
		ts := tisign.NewTiSign(headerContent, opt[0].SecretId, opt[0].SecretKey)
		header, _ = ts.CreateSignatureInfo()
	}

	// fmt.Printf("%v", header)

	err = JsonPost(ctx, uri, header, m, rspBody)
	if err != nil {
		return fmt.Errorf("JsonPost failed, err: %v", err)
	}
	return nil
}
