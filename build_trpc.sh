#!/bin/bash

set -e

module=`cat go.mod | grep module | awk {'print $2'}`

find ./proto/media -type f -name "*.proto" -exec sed -i "s#git.code.oa.com/yt-media-ai-videounderstanding/yt-ai-media-middle-platform-api#${module}#g" {} \;

# 兼容 sdk-prometheus-trpc-go
trpc create --api-version=1 --protofile=proto/media/common.proto --rpconly -o protobuf-spec/apicommon --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/video_structure.proto --rpconly -o protobuf-spec/videostructure --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/ai_tag_analyse.proto --rpconly -o protobuf-spec/aitaganalyse --mock=false --alias

trpc create --api-version=1 --protofile=proto/media/person_retrieval.proto --rpconly -o protobuf-spec/personretrieval --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/media.proto --rpconly -o protobuf-spec/media --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/task.proto --rpconly -o protobuf-spec/task --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/video_cut.proto --rpconly -o protobuf-spec/videocut --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/task_data.proto --rpconly -o protobuf-spec/taskdata --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/video_quality_evaluation.proto --rpconly -o protobuf-spec/videoquality --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/snapshot.proto --rpconly -o protobuf-spec/snapshot --mock=false --alias
trpc create --api-version=2 --protofile=proto/media/person_user_define.proto --rpconly -o protobuf-spec/personuserdefine --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/video_erase.proto --rpconly -o protobuf-spec/videoerase --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/custom_feature.proto --rpconly -o protobuf-spec/customfeature --mock=false --alias

trpc create --api-version=1 --protofile=proto/media/toolkit.proto --rpconly -o protobuf-spec/toolkit --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/media_crop.proto --rpconly -o protobuf-spec/mediacrop --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/domain_group.proto --rpconly -o protobuf-spec/domaingroup --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/workflow_template.proto --rpconly -o protobuf-spec/workflowtemplate --mock=false --alias

trpc create --api-version=1 --protofile=proto/media/yt_video_process.proto --rpconly -o protobuf-spec/ytvideoprocess --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/ai_video_process.proto --rpconly -o protobuf-spec/aivideoprocess --mock=false --alias

trpc create --api-version=1 --protofile=proto/media/text_summarization.proto --rpconly -o protobuf-spec/textsummarization --mock=false --alias
trpc create --api-version=1 --protofile=proto/media/shot_match.proto --rpconly -o protobuf-spec/shotmatch --mock=false --alias

find ./protobuf-spec -type f -name "*.trpc.go" -exec sed -i "s#git.code.oa.com/trpc-go/trpc-go#${module}/trpc_mock#g" {} \;