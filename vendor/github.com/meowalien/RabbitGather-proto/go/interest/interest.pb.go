// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: interest/interest.proto

package interest

import (
	share "github.com/meowalien/RabbitGather-proto/go/share"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CrawlerType int32

const (
	CrawlerType_SimpleCrawler CrawlerType = 0
)

// Enum value maps for CrawlerType.
var (
	CrawlerType_name = map[int32]string{
		0: "SimpleCrawler",
	}
	CrawlerType_value = map[string]int32{
		"SimpleCrawler": 0,
	}
)

func (x CrawlerType) Enum() *CrawlerType {
	p := new(CrawlerType)
	*p = x
	return p
}

func (x CrawlerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CrawlerType) Descriptor() protoreflect.EnumDescriptor {
	return file_interest_interest_proto_enumTypes[0].Descriptor()
}

func (CrawlerType) Type() protoreflect.EnumType {
	return &file_interest_interest_proto_enumTypes[0]
}

func (x CrawlerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CrawlerType.Descriptor instead.
func (CrawlerType) EnumDescriptor() ([]byte, []int) {
	return file_interest_interest_proto_rawDescGZIP(), []int{0}
}

type CrawlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data [][]byte `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CrawlResponse) Reset() {
	*x = CrawlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interest_interest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlResponse) ProtoMessage() {}

func (x *CrawlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interest_interest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlResponse.ProtoReflect.Descriptor instead.
func (*CrawlResponse) Descriptor() ([]byte, []int) {
	return file_interest_interest_proto_rawDescGZIP(), []int{0}
}

func (x *CrawlResponse) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CrawlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    CrawlerType           `protobuf:"varint,1,opt,name=Type,proto3,enum=CrawlerType" json:"Type,omitempty"`
	Message *share.EncodedMessage `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *CrawlRequest) Reset() {
	*x = CrawlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interest_interest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlRequest) ProtoMessage() {}

func (x *CrawlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interest_interest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlRequest.ProtoReflect.Descriptor instead.
func (*CrawlRequest) Descriptor() ([]byte, []int) {
	return file_interest_interest_proto_rawDescGZIP(), []int{1}
}

func (x *CrawlRequest) GetType() CrawlerType {
	if x != nil {
		return x.Type
	}
	return CrawlerType_SimpleCrawler
}

func (x *CrawlRequest) GetMessage() *share.EncodedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_interest_interest_proto protoreflect.FileDescriptor

var file_interest_interest_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x23, 0x0a, 0x0d, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x0c, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2a, 0x20, 0x0a, 0x0b, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x10, 0x00, 0x32, 0x39, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x12,
	0x0d, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x6f,
	0x77, 0x61, 0x6c, 0x69, 0x65, 0x6e, 0x2f, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x47, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interest_interest_proto_rawDescOnce sync.Once
	file_interest_interest_proto_rawDescData = file_interest_interest_proto_rawDesc
)

func file_interest_interest_proto_rawDescGZIP() []byte {
	file_interest_interest_proto_rawDescOnce.Do(func() {
		file_interest_interest_proto_rawDescData = protoimpl.X.CompressGZIP(file_interest_interest_proto_rawDescData)
	})
	return file_interest_interest_proto_rawDescData
}

var file_interest_interest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_interest_interest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interest_interest_proto_goTypes = []interface{}{
	(CrawlerType)(0),             // 0: CrawlerType
	(*CrawlResponse)(nil),        // 1: CrawlResponse
	(*CrawlRequest)(nil),         // 2: CrawlRequest
	(*share.EncodedMessage)(nil), // 3: EncodedMessage
}
var file_interest_interest_proto_depIdxs = []int32{
	0, // 0: CrawlRequest.Type:type_name -> CrawlerType
	3, // 1: CrawlRequest.Message:type_name -> EncodedMessage
	2, // 2: InterestCrawler.Crawl:input_type -> CrawlRequest
	1, // 3: InterestCrawler.Crawl:output_type -> CrawlResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_interest_interest_proto_init() }
func file_interest_interest_proto_init() {
	if File_interest_interest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interest_interest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlResponse); i {
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
		file_interest_interest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlRequest); i {
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
			RawDescriptor: file_interest_interest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interest_interest_proto_goTypes,
		DependencyIndexes: file_interest_interest_proto_depIdxs,
		EnumInfos:         file_interest_interest_proto_enumTypes,
		MessageInfos:      file_interest_interest_proto_msgTypes,
	}.Build()
	File_interest_interest_proto = out.File
	file_interest_interest_proto_rawDesc = nil
	file_interest_interest_proto_goTypes = nil
	file_interest_interest_proto_depIdxs = nil
}
