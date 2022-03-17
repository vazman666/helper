// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0--rc1
// source: proto/fromsql.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Firm   string `protobuf:"bytes,2,opt,name=firm,proto3" json:"firm,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fromsql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fromsql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_fromsql_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Request) GetFirm() string {
	if x != nil {
		return x.Firm
	}
	return ""
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid           string `protobuf:"bytes,1,opt,name=Oid,proto3" json:"Oid,omitempty"`
	FirmSql       string `protobuf:"bytes,2,opt,name=FirmSql,proto3" json:"FirmSql,omitempty"`
	PresencePrice string `protobuf:"bytes,3,opt,name=PresencePrice,proto3" json:"PresencePrice,omitempty"`
	SalesPrice    string `protobuf:"bytes,4,opt,name=SalesPrice,proto3" json:"SalesPrice,omitempty"`
	Caption       string `protobuf:"bytes,5,opt,name=Caption,proto3" json:"Caption,omitempty"`
	Number        string `protobuf:"bytes,6,opt,name=Number,proto3" json:"Number,omitempty"`
	Cellm         string `protobuf:"bytes,7,opt,name=Cellm,proto3" json:"Cellm,omitempty"`
	Cellt         string `protobuf:"bytes,8,opt,name=Cellt,proto3" json:"Cellt,omitempty"`
	Name          string `protobuf:"bytes,9,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fromsql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fromsql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_proto_fromsql_proto_rawDescGZIP(), []int{1}
}

func (x *Answer) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *Answer) GetFirmSql() string {
	if x != nil {
		return x.FirmSql
	}
	return ""
}

func (x *Answer) GetPresencePrice() string {
	if x != nil {
		return x.PresencePrice
	}
	return ""
}

func (x *Answer) GetSalesPrice() string {
	if x != nil {
		return x.SalesPrice
	}
	return ""
}

func (x *Answer) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *Answer) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Answer) GetCellm() string {
	if x != nil {
		return x.Cellm
	}
	return ""
}

func (x *Answer) GetCellt() string {
	if x != nil {
		return x.Cellt
	}
	return ""
}

func (x *Answer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_fromsql_proto protoreflect.FileDescriptor

var file_proto_fromsql_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x72, 0x6f, 0x6d, 0x73, 0x71, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x73, 0x71, 0x6c, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x69, 0x72, 0x6d, 0x22, 0xec, 0x01, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x4f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4f,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x72, 0x6d, 0x53, 0x71, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x69, 0x72, 0x6d, 0x53, 0x71, 0x6c, 0x12, 0x24, 0x0a, 0x0d,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x65, 0x6c, 0x6c, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x65, 0x6c, 0x6c, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x65,
	0x6c, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x65, 0x6c, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0x75, 0x0a, 0x0a, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x71, 0x6c, 0x12,
	0x10, 0x2e, 0x66, 0x72, 0x6f, 0x6d, 0x73, 0x71, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x66, 0x72, 0x6f, 0x6d, 0x73, 0x71, 0x6c, 0x2e, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x66,
	0x72, 0x6f, 0x6d, 0x73, 0x71, 0x6c, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x1a, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x28, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fromsql_proto_rawDescOnce sync.Once
	file_proto_fromsql_proto_rawDescData = file_proto_fromsql_proto_rawDesc
)

func file_proto_fromsql_proto_rawDescGZIP() []byte {
	file_proto_fromsql_proto_rawDescOnce.Do(func() {
		file_proto_fromsql_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fromsql_proto_rawDescData)
	})
	return file_proto_fromsql_proto_rawDescData
}

var file_proto_fromsql_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_fromsql_proto_goTypes = []interface{}{
	(*Request)(nil),              // 0: fromsql.Request
	(*Answer)(nil),               // 1: fromsql.Answer
	(*wrapperspb.BoolValue)(nil), // 2: google.protobuf.BoolValue
}
var file_proto_fromsql_proto_depIdxs = []int32{
	0, // 0: fromsql.SqlRequest.StreamSql:input_type -> fromsql.Request
	1, // 1: fromsql.SqlRequest.Change:input_type -> fromsql.Answer
	1, // 2: fromsql.SqlRequest.StreamSql:output_type -> fromsql.Answer
	2, // 3: fromsql.SqlRequest.Change:output_type -> google.protobuf.BoolValue
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_fromsql_proto_init() }
func file_proto_fromsql_proto_init() {
	if File_proto_fromsql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fromsql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_fromsql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
			RawDescriptor: file_proto_fromsql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fromsql_proto_goTypes,
		DependencyIndexes: file_proto_fromsql_proto_depIdxs,
		MessageInfos:      file_proto_fromsql_proto_msgTypes,
	}.Build()
	File_proto_fromsql_proto = out.File
	file_proto_fromsql_proto_rawDesc = nil
	file_proto_fromsql_proto_goTypes = nil
	file_proto_fromsql_proto_depIdxs = nil
}