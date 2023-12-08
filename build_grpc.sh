#!/bin/bash

set -e

module=`cat go.mod | grep module | awk {'print $2'}`

find ./proto/media -type f -name "*.proto" -exec sed -i "s#git.code.oa.com/yt-media-ai-videounderstanding/yt-ai-media-middle-platform-api#${module}#g" {} \;


function gen() {
  dir=$1
  proto=$2
  mkdir -p ${dir}
  protoc --go_out=${dir} --go_opt=paths=source_relative  --go-grpc_out=${dir} --go-grpc_opt=paths=source_relative  -I ./proto/media ${proto}.proto
}

gen "protobuf-spec/apicommon" "common"

gen "protobuf-spec/media" "media"

gen "protobuf-spec/videostructure" "video_structure"

gen "protobuf-spec/aitaganalyse" "ai_tag_analyse"

gen "protobuf-spec/personretrieval" "person_retrieval"

gen "protobuf-spec/task" "task"

gen "protobuf-spec/videocut" "video_cut"

gen "protobuf-spec/taskdata" "task_data"

gen "protobuf-spec/videoquality" "video_quality_evaluation"

gen "protobuf-spec/snapshot" "snapshot"

gen "protobuf-spec/personuserdefine" "person_user_define"

gen "protobuf-spec/videoerase" "video_erase"

gen "protobuf-spec/customfeature" "custom_feature"

gen "protobuf-spec/toolkit" "toolkit"

gen "protobuf-spec/mediacrop" "media_crop"

gen "protobuf-spec/domaingroup" "domain_group"

gen "protobuf-spec/workflowtemplate" "workflow_template"

gen "protobuf-spec/ytvideoprocess" "yt_video_process"

gen "protobuf-spec/aivideoprocess" "ai_video_process"

gen "protobuf-spec/textsummarization" "text_summarization"

gen "protobuf-spec/shotmatch" "shot_match"

