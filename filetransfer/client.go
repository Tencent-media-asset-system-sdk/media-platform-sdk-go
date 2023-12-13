package filetransfer

import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// Client ...
type Client struct {
	opt client.Option
}

// NewFiletransferClientProxy ...
func NewFiletransferClientProxy(opt client.Option) *Client {
	if opt.Port == 0 {
		opt.Port = 80
	}
	return &Client{opt: opt}
}
