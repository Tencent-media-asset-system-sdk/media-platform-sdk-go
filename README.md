# 腾讯云媒体AI中台私有化对接 SDK-golang

## 调用单个接口
以 CreateMedias 接口为例

分三步
1. 构建输入 request;
2. 创建 client proxy，需要传入 secretId, secretKey, projectId, businessId;
3. 调用接口。
```go
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
	infos := []*media.MediaUnionInfo{
		{
			MediaInfo: &media.MediaInfo{
				MediaType: apicommon.MediaType_MEDIA_TYPE_VIDEO,
				MediaName: "测试",
				MediaTag:  apicommon.CategoryTagType(apicommon.MediaSourceType_MEDIA_SOURCE_TYPE_OTHER),
			},
			DomainGroupInfo: &apicommon.DomainGroupInfo{
				DomainGroupType: apicommon.DomainGroupType_DOMAIN_GROUP_TYPE_PUBLIC,
			},
			FileInfo: &media.FileInfo{
				FileName: "test.mp4",
				FileType: "video/mp4",
				FileSize: 102400,
				Key:      "filekey-xxx", // 通过文件管理上传的 key
			},
			MediaWorkflowTemplateInfo: &media.MediaWorkflowTemplateInfo{
				WorkflowMode: apicommon.WorkflowMode_WORKFLOW_MODE_CUSTOMER, // 用户资源
        AutoMatchTemplate: true, // 自动匹配入库模版
			},
		},
	}
	req := &media.CreateMediasRequest{
		MediaUnionInfoSet: infos,
	}
	// 创建ClientProxy
	proxy := media.NewMediaClientProxy(client.Option{
		Host:         host,
		SecretId:     secretId, // ti 平台个人中心获取
		SecretKey:    secretKey, // ti 平台个人中心获取
		TIBusinessId: uint32(businessId), // ti 平台管理中心获取
		TIProjectId:  uint32(projectId), // ti 平台管理中心获取
	})
	ctx := context.Background()
	// 调用接口
	rsp, err := proxy.CreateMedias(ctx, req)
	if err != nil {
		t.Fatalf("CreateMedias error: %v", err)
	}
	requestId := ctx.Value(codec.ContextKeyRequestId) // 获取 RequestId
	t.Fatalf("rsp: %v, RequestId: %v", rsp, requestId)
}

```

更多用例查看 [test](https://github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/tree/master/test) 目录

## 上传文件
文件上传模块处理提供了单个接口以外，还提供一个一键上传文件的方法

```go
package filetransfer_test

import (
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/filetransfer"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// 上传那本地文件
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

	if key, _, _, err := proxy.UploadFile(context.Background(), filePath, "测试文件1"); err != nil {
		t.Fatalf("upload file error %v", err)
	} else {
		t.Logf("success upload file key %s", key)
	}
}

// 上传内存 buf
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

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("file not found")
	}
	if key, err := proxy.UploadBuf(context.Background(), buf, "测试文件2"); err != nil {
		t.Fatalf("upload buf error %v", err)
	} else {
		t.Logf("success upload buf key %s", key)
	}
}
```

## 每个模块的 proxy 创建方法
### 文件管理
```go
import (
  "github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/filetransfer"
  "github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := filetransfer.NewFiletransferClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 资源和资源包管理
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/media"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := media.NewMediaClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 任务管理
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/task"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := media.NewTaskClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 智能标签
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/videostructure"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := videostructure.NewAITagClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 智能拆条
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/videostructure"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := videostructure.NewAICutClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 智能编目
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/videostructure"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := videostructure.NewAICatalogClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 人脸集锦
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/personretrieval"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := personretrieval.NewPersonRetrievalClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 自定义人脸库
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/personuserdefine"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := personuserdefine.NewUserDefinePersonClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 自定义标签
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/customfeature"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := videostructure.NewCustomFeatureClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```


### 内容中台
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/media"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/task"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/toolkit"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// 检索
retrieveproxy := media.NewMediaClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})

// 工具箱
tookitproxy := toolkit.NewToolkitClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 视频处理
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/aivideoprocess"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := aivideoprocess.NewAIVideoProcessClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 视频质检
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/videoquality"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

proxy := videoquality.NewVideoQualityEvaluationClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

### 视频擦除
```go
import (
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/protobuf-spec/videoquality"
	"github.com/Tencent-media-asset-system-sdk/media-platform-sdk-go/trpc_mock/client"
)

// 去标题
captionproxy := videoerase.NewRemoveCaptionClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})

// 去字幕
logoproxy := videoerase.NewRemoveLogoClientProxy(client.Option{
	Host:         host,
	SecretId:     secretId,
	SecretKey:    secretKey,
	TIBusinessId: uint32(businessId),
	TIProjectId:  uint32(projectId),
})
```

