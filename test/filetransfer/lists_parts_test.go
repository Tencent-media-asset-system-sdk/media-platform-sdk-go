package filetransfer

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/filetransfer"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

func TestListParts(t *testing.T) {
	host := os.Getenv("ENV_HOST")
	secretId := os.Getenv("ENV_SECRET_ID")
	secretKey := os.Getenv("ENV_SECRET_KEY")
	projectId, _ := strconv.ParseUint(os.Getenv("ENV_PROJECT_ID"), 10, 32)
	businessId, _ := strconv.ParseUint(os.Getenv("ENV_BUSINESS_ID"), 10, 32)
	key := os.Getenv("ENV_KEY")
	t.Logf("%s %s %s %d %d %s", host, secretId, secretKey, projectId, businessId, key)

	proxy := filetransfer.NewFiletransferClientProxy(client.Option{
		Host:         host,
		SecretId:     secretId,
		SecretKey:    secretKey,
		TIBusinessId: uint32(businessId),
		TIProjectId:  uint32(projectId),
	})

	if _, err := proxy.ListParts(context.Background(), key, "", "xxxx"); err != nil {
		t.Fatalf("upload file error %v", err)
	}
}
