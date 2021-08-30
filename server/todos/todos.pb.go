// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: todos.proto

package todos

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

type Todos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *Todos) Reset() {
	*x = Todos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todos) ProtoMessage() {}

func (x *Todos) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todos.ProtoReflect.Descriptor instead.
func (*Todos) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{0}
}

func (x *Todos) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Done  bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{1}
}

func (x *Todo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_todos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_todos_proto_rawDescGZIP(), []int{2}
}

var File_todos_proto protoreflect.FileDescriptor

var file_todos_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x22, 0x2a, 0x0a, 0x05, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x21, 0x0a,
	0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x22, 0x40, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f,
	0x6e, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x5c, 0x0a, 0x07, 0x54,
	0x6f, 0x64, 0x6f, 0x41, 0x70, 0x70, 0x12, 0x26, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x1a, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x29,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x0c, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x64, 0x79, 0x64,
	0x78, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todos_proto_rawDescOnce sync.Once
	file_todos_proto_rawDescData = file_todos_proto_rawDesc
)

func file_todos_proto_rawDescGZIP() []byte {
	file_todos_proto_rawDescOnce.Do(func() {
		file_todos_proto_rawDescData = protoimpl.X.CompressGZIP(file_todos_proto_rawDescData)
	})
	return file_todos_proto_rawDescData
}

var file_todos_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_todos_proto_goTypes = []interface{}{
	(*Todos)(nil), // 0: todos.Todos
	(*Todo)(nil),  // 1: todos.Todo
	(*Empty)(nil), // 2: todos.Empty
}
var file_todos_proto_depIdxs = []int32{
	1, // 0: todos.Todos.todos:type_name -> todos.Todo
	1, // 1: todos.TodoApp.CreateTodo:input_type -> todos.Todo
	2, // 2: todos.TodoApp.GetAllTodos:input_type -> todos.Empty
	1, // 3: todos.TodoApp.CreateTodo:output_type -> todos.Todo
	0, // 4: todos.TodoApp.GetAllTodos:output_type -> todos.Todos
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todos_proto_init() }
func file_todos_proto_init() {
	if File_todos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todos); i {
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
		file_todos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_todos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_todos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todos_proto_goTypes,
		DependencyIndexes: file_todos_proto_depIdxs,
		MessageInfos:      file_todos_proto_msgTypes,
	}.Build()
	File_todos_proto = out.File
	file_todos_proto_rawDesc = nil
	file_todos_proto_goTypes = nil
	file_todos_proto_depIdxs = nil
}
