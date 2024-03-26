package filetransfer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JsonGet http json get 请求
func JsonGet(ctx context.Context, uri string, header map[string]string, rsp interface{}) (err error) {
	reqCtx, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return fmt.Errorf("JsonGet make NewRequest error: %v", err)
	}

	for head, value := range header {
		reqCtx.Header.Set(head, value)
	}

	response, err := http.DefaultClient.Do(reqCtx)
	if err != nil {
		return fmt.Errorf("JsonGet do request error: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("JsonGet response status error, statuscode: %d, reason: %s",
			response.StatusCode, response.Status)
	}
	data, _ := ioutil.ReadAll(response.Body)
	if rsp != nil {
		err = json.Unmarshal(data, rsp)
		if err != nil {
			return fmt.Errorf("JsonGet response protocol error: %v, data: %s", err, string(data))
		}
	}
	return nil
}

// ListPartsResponse ...
type ListPartsResponse struct {
	Response struct {
		CommonResponse
		ListPartsResult struct {
			Bucket               string `json:"Bucket"`
			Key                  string `json:"Key"`
			UploadId             string `json:"UploadId"`
			PartNumberMarker     int32  `json:"PartNumberMarker"`
			NextPartNumberMarker int32  `json:"NextPartNumberMarker"`
			MaxParts             int32  `json:"MaxParts"`
			IsTruncated          bool   `json:"IsTruncated"`
			Part                 []struct {
				PartNumber   int32  `json:"PartNumber"`
				Size         int32  `json:"Size"`
				ETag         string `json:"ETag"`
				LastModified string `json:"LastModified"`
			} `json:"Part"`
		} `json:"ListPartsResult"`
	} `json:"Response"`
}

// ListParts 获取上传分片列表
func (c Client) ListParts(ctx context.Context,
	key, bucket, uploadID string) (response *ListPartsResponse, err error) {

	rsp := &ListPartsResponse{}
	uri := fmt.Sprintf("http://%s/FileManager/ListParts?useJson=true&Bucket=%s&Key=%s&uploadId=%s",
		c.opt.GetEndpoint(), bucket, key, uploadID)

	err = JsonGet(ctx, uri, nil, &rsp.Response)
	if err != nil {

		return nil, fmt.Errorf("ListParts failed, err: %v", err)
	}

	if rsp.Response.Error != nil {
		respBody, _ := json.Marshal(rsp)
		return nil, fmt.Errorf("ListParts response failed, response: %s", string(respBody))
	}
	return rsp, err
}
