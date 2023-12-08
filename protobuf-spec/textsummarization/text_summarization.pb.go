// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: text_summarization.proto

package textsummarization

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sentence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text                string    `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Keywords            []string  `protobuf:"bytes,2,rep,name=Keywords,proto3" json:"Keywords,omitempty"`
	PositionPercentages []float32 `protobuf:"fixed32,3,rep,packed,name=PositionPercentages,proto3" json:"PositionPercentages,omitempty"`
}

func (x *Sentence) Reset() {
	*x = Sentence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_summarization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sentence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sentence) ProtoMessage() {}

func (x *Sentence) ProtoReflect() protoreflect.Message {
	mi := &file_text_summarization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sentence.ProtoReflect.Descriptor instead.
func (*Sentence) Descriptor() ([]byte, []int) {
	return file_text_summarization_proto_rawDescGZIP(), []int{0}
}

func (x *Sentence) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Sentence) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *Sentence) GetPositionPercentages() []float32 {
	if x != nil {
		return x.PositionPercentages
	}
	return nil
}

type TextSummarizationTaskData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary   string      `protobuf:"bytes,1,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Title     string      `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Sentences []*Sentence `protobuf:"bytes,3,rep,name=Sentences,proto3" json:"Sentences,omitempty"`
	Keywords  []string    `protobuf:"bytes,4,rep,name=Keywords,proto3" json:"Keywords,omitempty"`
}

func (x *TextSummarizationTaskData) Reset() {
	*x = TextSummarizationTaskData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_summarization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSummarizationTaskData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSummarizationTaskData) ProtoMessage() {}

func (x *TextSummarizationTaskData) ProtoReflect() protoreflect.Message {
	mi := &file_text_summarization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSummarizationTaskData.ProtoReflect.Descriptor instead.
func (*TextSummarizationTaskData) Descriptor() ([]byte, []int) {
	return file_text_summarization_proto_rawDescGZIP(), []int{1}
}

func (x *TextSummarizationTaskData) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *TextSummarizationTaskData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TextSummarizationTaskData) GetSentences() []*Sentence {
	if x != nil {
		return x.Sentences
	}
	return nil
}

func (x *TextSummarizationTaskData) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

var File_text_summarization_proto protoreflect.FileDescriptor

var file_text_summarization_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x13, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x19, 0x54, 0x65, 0x78, 0x74, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x09, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x61, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_text_summarization_proto_rawDescOnce sync.Once
	file_text_summarization_proto_rawDescData = file_text_summarization_proto_rawDesc
)

func file_text_summarization_proto_rawDescGZIP() []byte {
	file_text_summarization_proto_rawDescOnce.Do(func() {
		file_text_summarization_proto_rawDescData = protoimpl.X.CompressGZIP(file_text_summarization_proto_rawDescData)
	})
	return file_text_summarization_proto_rawDescData
}

var file_text_summarization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_text_summarization_proto_goTypes = []interface{}{
	(*Sentence)(nil),                  // 0: trpc.media.textsummarization.Sentence
	(*TextSummarizationTaskData)(nil), // 1: trpc.media.textsummarization.TextSummarizationTaskData
}
var file_text_summarization_proto_depIdxs = []int32{
	0, // 0: trpc.media.textsummarization.TextSummarizationTaskData.Sentences:type_name -> trpc.media.textsummarization.Sentence
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_text_summarization_proto_init() }
func file_text_summarization_proto_init() {
	if File_text_summarization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_text_summarization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sentence); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_text_summarization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSummarizationTaskData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_text_summarization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_text_summarization_proto_goTypes,
		DependencyIndexes: file_text_summarization_proto_depIdxs,
		MessageInfos:      file_text_summarization_proto_msgTypes,
	}.Build()
	File_text_summarization_proto = out.File
	file_text_summarization_proto_rawDesc = nil
	file_text_summarization_proto_goTypes = nil
	file_text_summarization_proto_depIdxs = nil
}
