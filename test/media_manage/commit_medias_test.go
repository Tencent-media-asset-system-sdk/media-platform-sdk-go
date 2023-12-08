package mediamanage_test

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/media"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/codec"
)

func TestCommitMedias(t *testing.T) {
	host := os.Getenv("ENV_HOST")
	secretId := os.Getenv("ENV_SECRET_ID")
	secretKey := os.Getenv("ENV_SECRET_KEY")
	projectId, _ := strconv.ParseUint(os.Getenv("ENV_PROJECT_ID"), 10, 32)
	businessId, _ := strconv.ParseUint(os.Getenv("ENV_BUSINESS_ID"), 10, 32)
	t.Logf("%s %s %s %d %d", host, secretId, secretKey, projectId, businessId)
	req := &media.CommitMediasRequest{
		MediaCommitInfoSet: []*media.MediaCommitInfo{
			{
				MediaId:       "media-f2gzwp70-0",
				UploadStatus:  "success",
				UploadMessage: "upload success",
			},
		},
	}
	proxy := media.NewMediaClientProxy(client.Option{
		Host:         host,
		SecretId:     secretId,
		SecretKey:    secretKey,
		TIBusinessId: uint32(businessId),
		TIProjectId:  uint32(projectId),
	})
	ctx := context.Background()
	rsp, err := proxy.CommitMedias(ctx, req)
	if err != nil {
		t.Fatalf("CommitMedias error: %v", err)
	}
	requestId := ctx.Value(codec.ContextKeyRequestId) // 获取 RequestId
	t.Fatalf("rsp: %v, RequestId: %v", rsp, requestId)

}
