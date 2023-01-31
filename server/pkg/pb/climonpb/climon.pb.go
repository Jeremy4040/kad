// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.6
// source: climon.proto

package climonpb

import (
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

type DeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	RepoUrl     string `protobuf:"bytes,2,opt,name=RepoUrl,proto3" json:"RepoUrl,omitempty"`
	RepoName    string `protobuf:"bytes,3,opt,name=RepoName,proto3" json:"RepoName,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	ChartName   string `protobuf:"bytes,5,opt,name=ChartName,proto3" json:"ChartName,omitempty"`
	ReleaseName string `protobuf:"bytes,6,opt,name=ReleaseName,proto3" json:"ReleaseName,omitempty"`
	ReferenceID string `protobuf:"bytes,7,opt,name=ReferenceID,proto3" json:"ReferenceID,omitempty"`
	Plugin      string `protobuf:"bytes,8,opt,name=Plugin,proto3" json:"Plugin,omitempty"`
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_climon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_climon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_climon_proto_rawDescGZIP(), []int{0}
}

func (x *DeployRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DeployRequest) GetRepoUrl() string {
	if x != nil {
		return x.RepoUrl
	}
	return ""
}

func (x *DeployRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *DeployRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeployRequest) GetChartName() string {
	if x != nil {
		return x.ChartName
	}
	return ""
}

func (x *DeployRequest) GetReleaseName() string {
	if x != nil {
		return x.ReleaseName
	}
	return ""
}

func (x *DeployRequest) GetReferenceID() string {
	if x != nil {
		return x.ReferenceID
	}
	return ""
}

func (x *DeployRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

type DeployResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                 string            `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	WorkFlowResponseStatus *WorkFlowResponse `protobuf:"bytes,2,opt,name=WorkFlowResponseStatus,proto3" json:"WorkFlowResponseStatus,omitempty"`
}

func (x *DeployResponse) Reset() {
	*x = DeployResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_climon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployResponse) ProtoMessage() {}

func (x *DeployResponse) ProtoReflect() protoreflect.Message {
	mi := &file_climon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployResponse.ProtoReflect.Descriptor instead.
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return file_climon_proto_rawDescGZIP(), []int{1}
}

func (x *DeployResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeployResponse) GetWorkFlowResponseStatus() *WorkFlowResponse {
	if x != nil {
		return x.WorkFlowResponseStatus
	}
	return nil
}

type WorkFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowID    string `protobuf:"bytes,1,opt,name=WorkflowID,proto3" json:"WorkflowID,omitempty"`
	WorkFlowRunID string `protobuf:"bytes,2,opt,name=WorkFlowRunID,proto3" json:"WorkFlowRunID,omitempty"`
}

func (x *WorkFlowResponse) Reset() {
	*x = WorkFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_climon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkFlowResponse) ProtoMessage() {}

func (x *WorkFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_climon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkFlowResponse.ProtoReflect.Descriptor instead.
func (*WorkFlowResponse) Descriptor() ([]byte, []int) {
	return file_climon_proto_rawDescGZIP(), []int{2}
}

func (x *WorkFlowResponse) GetWorkflowID() string {
	if x != nil {
		return x.WorkflowID
	}
	return ""
}

func (x *WorkFlowResponse) GetWorkFlowRunID() string {
	if x != nil {
		return x.WorkFlowRunID
	}
	return ""
}

var File_climon_proto protoreflect.FileDescriptor

var file_climon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x63, 0x6c, 0x69, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x22, 0xf7, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x22, 0x7c, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x52, 0x0a, 0x16,
	0x57, 0x6f, 0x72, 0x6b, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6c, 0x69, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x46, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x16, 0x57, 0x6f, 0x72, 0x6b, 0x46, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x58, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x46, 0x6c, 0x6f, 0x77,
	0x52, 0x75, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x57, 0x6f, 0x72,
	0x6b, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x49, 0x44, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x63,
	0x6c, 0x69, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_climon_proto_rawDescOnce sync.Once
	file_climon_proto_rawDescData = file_climon_proto_rawDesc
)

func file_climon_proto_rawDescGZIP() []byte {
	file_climon_proto_rawDescOnce.Do(func() {
		file_climon_proto_rawDescData = protoimpl.X.CompressGZIP(file_climon_proto_rawDescData)
	})
	return file_climon_proto_rawDescData
}

var file_climon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_climon_proto_goTypes = []interface{}{
	(*DeployRequest)(nil),    // 0: climonpb.DeployRequest
	(*DeployResponse)(nil),   // 1: climonpb.DeployResponse
	(*WorkFlowResponse)(nil), // 2: climonpb.WorkFlowResponse
}
var file_climon_proto_depIdxs = []int32{
	2, // 0: climonpb.DeployResponse.WorkFlowResponseStatus:type_name -> climonpb.WorkFlowResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_climon_proto_init() }
func file_climon_proto_init() {
	if File_climon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_climon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployRequest); i {
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
		file_climon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployResponse); i {
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
		file_climon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkFlowResponse); i {
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
			RawDescriptor: file_climon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_climon_proto_goTypes,
		DependencyIndexes: file_climon_proto_depIdxs,
		MessageInfos:      file_climon_proto_msgTypes,
	}.Build()
	File_climon_proto = out.File
	file_climon_proto_rawDesc = nil
	file_climon_proto_goTypes = nil
	file_climon_proto_depIdxs = nil
}
