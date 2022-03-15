// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: captcha.proto

package captcha

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Difficulty int32  `protobuf:"varint,2,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
}

func (x *TaskPayload) Reset() {
	*x = TaskPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskPayload) ProtoMessage() {}

func (x *TaskPayload) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskPayload.ProtoReflect.Descriptor instead.
func (*TaskPayload) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *TaskPayload) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TaskPayload) GetDifficulty() int32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

type GenerateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Math        string       `protobuf:"bytes,1,opt,name=math,proto3" json:"math,omitempty"`
	TaskPayload *TaskPayload `protobuf:"bytes,2,opt,name=task_payload,json=taskPayload,proto3" json:"task_payload,omitempty"`
}

func (x *GenerateTaskResponse) Reset() {
	*x = GenerateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTaskResponse) ProtoMessage() {}

func (x *GenerateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTaskResponse.ProtoReflect.Descriptor instead.
func (*GenerateTaskResponse) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTaskResponse) GetMath() string {
	if x != nil {
		return x.Math
	}
	return ""
}

func (x *GenerateTaskResponse) GetTaskPayload() *TaskPayload {
	if x != nil {
		return x.TaskPayload
	}
	return nil
}

type ValidateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Answer int32  `protobuf:"varint,2,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *ValidateTaskRequest) Reset() {
	*x = ValidateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTaskRequest) ProtoMessage() {}

func (x *ValidateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTaskRequest.ProtoReflect.Descriptor instead.
func (*ValidateTaskRequest) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateTaskRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ValidateTaskRequest) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type ValidateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAnswerCorrect bool `protobuf:"varint,1,opt,name=is_answer_correct,json=isAnswerCorrect,proto3" json:"is_answer_correct,omitempty"`
}

func (x *ValidateTaskResponse) Reset() {
	*x = ValidateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTaskResponse) ProtoMessage() {}

func (x *ValidateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTaskResponse.ProtoReflect.Descriptor instead.
func (*ValidateTaskResponse) Descriptor() ([]byte, []int) {
	return file_captcha_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateTaskResponse) GetIsAnswerCorrect() bool {
	if x != nil {
		return x.IsAnswerCorrect
	}
	return false
}

var File_captcha_proto protoreflect.FileDescriptor

var file_captcha_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x22, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x61, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x41, 0x0a,
	0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x22, 0x42, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x72,
	0x72, 0x65, 0x63, 0x74, 0x32, 0xa4, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1d, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1c,
	0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x61, 0x6e,
	0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x68, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x3b, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_captcha_proto_rawDescOnce sync.Once
	file_captcha_proto_rawDescData = file_captcha_proto_rawDesc
)

func file_captcha_proto_rawDescGZIP() []byte {
	file_captcha_proto_rawDescOnce.Do(func() {
		file_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_captcha_proto_rawDescData)
	})
	return file_captcha_proto_rawDescData
}

var file_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_captcha_proto_goTypes = []interface{}{
	(*TaskPayload)(nil),          // 0: captcha.TaskPayload
	(*GenerateTaskResponse)(nil), // 1: captcha.GenerateTaskResponse
	(*ValidateTaskRequest)(nil),  // 2: captcha.ValidateTaskRequest
	(*ValidateTaskResponse)(nil), // 3: captcha.ValidateTaskResponse
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_captcha_proto_depIdxs = []int32{
	0, // 0: captcha.GenerateTaskResponse.task_payload:type_name -> captcha.TaskPayload
	4, // 1: captcha.CaptchaService.generateTask:input_type -> google.protobuf.Empty
	2, // 2: captcha.CaptchaService.validateTask:input_type -> captcha.ValidateTaskRequest
	1, // 3: captcha.CaptchaService.generateTask:output_type -> captcha.GenerateTaskResponse
	3, // 4: captcha.CaptchaService.validateTask:output_type -> captcha.ValidateTaskResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_captcha_proto_init() }
func file_captcha_proto_init() {
	if File_captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_captcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskPayload); i {
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
		file_captcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTaskResponse); i {
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
		file_captcha_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTaskRequest); i {
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
		file_captcha_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTaskResponse); i {
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
			RawDescriptor: file_captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_captcha_proto_goTypes,
		DependencyIndexes: file_captcha_proto_depIdxs,
		MessageInfos:      file_captcha_proto_msgTypes,
	}.Build()
	File_captcha_proto = out.File
	file_captcha_proto_rawDesc = nil
	file_captcha_proto_goTypes = nil
	file_captcha_proto_depIdxs = nil
}
