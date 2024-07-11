package filetransfer_test

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/filetransfer"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

func TestUploadFile(t *testing.T) {
	host := os.Getenv("ENV_HOST")
	secretId := os.Getenv("ENV_SECRET_ID")
	secretKey := os.Getenv("ENV_SECRET_KEY")
	projectId, _ := strconv.ParseUint(os.Getenv("ENV_PROJECT_ID"), 10, 32)
	businessId, _ := strconv.ParseUint(os.Getenv("ENV_BUSINESS_ID"), 10, 32)
	filePath := os.Getenv("ENV_FILE")
	t.Logf("%s %s %s %d %d %s", host, secretId, secretKey, projectId, businessId, filePath)

	proxy := filetransfer.NewFiletransferClientProxy(client.Option{
		Host:         host,
		SecretId:     secretId,
		SecretKey:    secretKey,
		TIBusinessId: uint32(businessId),
		TIProjectId:  uint32(projectId),
	})

	ctx, _ := context.WithTimeout(context.Background(), time.Hour)
	start := time.Now()
	if key, _, _, err := proxy.UploadFile(ctx, filePath, "测试文件1"); err != nil {
		t.Fatalf("upload file error %v", err)
	} else {
		t.Logf("success upload file key %s", key)
	}
	t.Logf("upload time cost: %fs", time.Now().Sub(start).Seconds())
}

func TestUploadBuf(t *testing.T) {
	host := os.Getenv("ENV_HOST")
	secretId := os.Getenv("ENV_SECRET_ID")
	secretKey := os.Getenv("ENV_SECRET_KEY")
	projectId, _ := strconv.ParseUint(os.Getenv("ENV_PROJECT_ID"), 10, 32)
	businessId, _ := strconv.ParseUint(os.Getenv("ENV_BUSINESS_ID"), 10, 32)
	filePath := os.Getenv("ENV_FILE")
	t.Logf("%s %s %s %d %d %s", host, secretId, secretKey, projectId, businessId, filePath)

	proxy := filetransfer.NewFiletransferClientProxy(client.Option{
		Host:         host,
		SecretId:     secretId,
		SecretKey:    secretKey,
		TIBusinessId: uint32(businessId),
		TIProjectId:  uint32(projectId),
	})

	buf, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("file not found")
	}
	if key, err := proxy.UploadBuf(context.Background(), buf, "测试文件2"); err != nil {
		t.Fatalf("upload buf error %v", err)
	} else {
		t.Logf("success upload buf key %s", key)
	}
}
