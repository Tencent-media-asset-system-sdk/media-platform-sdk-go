package mediamanage_test

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/apicommon"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/media"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/codec"
)

func TestCreateMedias(t *testing.T) {
	host := os.Getenv("ENV_HOST")
	secretId := os.Getenv("ENV_SECRET_ID")
	secretKey := os.Getenv("ENV_SECRET_KEY")
	projectId, _ := strconv.ParseUint(os.Getenv("ENV_PROJECT_ID"), 10, 32)
	businessId, _ := strconv.ParseUint(os.Getenv("ENV_BUSINESS_ID"), 10, 32)
	t.Logf("%s %s %s %d %d", host, secretId, secretKey, projectId, businessId)
	infos := []*media.MediaUnionInfo{
		{
			MediaInfo: &media.MediaInfo{
				MediaType: apicommon.MediaType_MEDIA_TYPE_VIDEO,
				MediaName: "测试",
				MediaTag:  apicommon.CategoryTagType(apicommon.MediaSourceType_MEDIA_SOURCE_TYPE_OTHER),
			},
			DomainGroupInfo: &apicommon.DomainGroupInfo{
				DomainGroupType: apicommon.DomainGroupType_DOMAIN_GROUP_TYPE_PRIVATE,
				DomainGroupId:   "0",
			},
			TranscodeInfo: &media.TranscodeInfo{
				TranscodeFormat:  "mp4",
				TranscodeBitRate: 600,
				TranscodeWidth:   480,
				TranscodeHeight:  360,
			},
			FileInfo: &media.FileInfo{
				FileName: "test.mp4",
				FileType: "video/mp4",
				FileSize: 1024,
				Bucket:   "bucket123",
				UploadId: "uploadid123",
				Key:      "media/test/test_video/test-beijingninzao-6mins.mp4",
			},
			MediaWorkflowTemplateInfo: &media.MediaWorkflowTemplateInfo{
				WorkflowMode: apicommon.WorkflowMode_WORKFLOW_MODE_CUSTOMER,
			},
		},
	}
	req := &media.CreateMediasRequest{
		MediaUnionInfoSet: infos,
	}
	proxy := media.NewMediaClientProxy(client.Option{
		Host:         host,
		SecretId:     secretId,
		SecretKey:    secretKey,
		TIBusinessId: uint32(businessId),
		TIProjectId:  uint32(projectId),
	})
	ctx := context.Background()
	rsp, err := proxy.CreateMedias(ctx, req)
	if err != nil {
		t.Fatalf("CreateMedias error: %v", err)
	}
	requestId := ctx.Value(codec.ContextKeyRequestId) // 获取 RequestId
	t.Fatalf("rsp: %v, RequestId: %v", rsp, requestId)

}
