// Code generated by protoc-gen-go. DO NOT EDIT.
// source: verrpc/verrpc.proto

package verrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VersionRequest struct {
	User_Id              string   `protobuf:"bytes,1,opt,name=User_Id,json=UserId,proto3" json:"User_Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_494312204cefa0e6, []int{0}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func (m *VersionRequest) GetUser_Id() string {
	if m != nil {
		return m.User_Id
	}
	return ""
}

type Version struct {
	// A verbose description of the daemon's commit.
	Commit string `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	// The SHA1 commit hash that the daemon is compiled with.
	CommitHash string `protobuf:"bytes,2,opt,name=commit_hash,json=commitHash,proto3" json:"commit_hash,omitempty"`
	// The semantic version.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// The major application version.
	AppMajor uint32 `protobuf:"varint,4,opt,name=app_major,json=appMajor,proto3" json:"app_major,omitempty"`
	// The minor application version.
	AppMinor uint32 `protobuf:"varint,5,opt,name=app_minor,json=appMinor,proto3" json:"app_minor,omitempty"`
	// The application patch number.
	AppPatch uint32 `protobuf:"varint,6,opt,name=app_patch,json=appPatch,proto3" json:"app_patch,omitempty"`
	// The application pre-release modifier, possibly empty.
	AppPreRelease string `protobuf:"bytes,7,opt,name=app_pre_release,json=appPreRelease,proto3" json:"app_pre_release,omitempty"`
	// The list of build tags that were supplied during compilation.
	BuildTags []string `protobuf:"bytes,8,rep,name=build_tags,json=buildTags,proto3" json:"build_tags,omitempty"`
	// The version of go that compiled the executable.
	GoVersion            string   `protobuf:"bytes,9,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_494312204cefa0e6, []int{1}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *Version) GetCommitHash() string {
	if m != nil {
		return m.CommitHash
	}
	return ""
}

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetAppMajor() uint32 {
	if m != nil {
		return m.AppMajor
	}
	return 0
}

func (m *Version) GetAppMinor() uint32 {
	if m != nil {
		return m.AppMinor
	}
	return 0
}

func (m *Version) GetAppPatch() uint32 {
	if m != nil {
		return m.AppPatch
	}
	return 0
}

func (m *Version) GetAppPreRelease() string {
	if m != nil {
		return m.AppPreRelease
	}
	return ""
}

func (m *Version) GetBuildTags() []string {
	if m != nil {
		return m.BuildTags
	}
	return nil
}

func (m *Version) GetGoVersion() string {
	if m != nil {
		return m.GoVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "verrpc.VersionRequest")
	proto.RegisterType((*Version)(nil), "verrpc.Version")
}

func init() { proto.RegisterFile("verrpc/verrpc.proto", fileDescriptor_494312204cefa0e6) }

var fileDescriptor_494312204cefa0e6 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0x86, 0xd9, 0xf6, 0xff, 0xed, 0xf2, 0xc9, 0x1c, 0x44, 0xd0, 0xa0, 0x88, 0x63, 0x07, 0x32,
	0x41, 0x5a, 0x50, 0xbc, 0x81, 0x9d, 0xe8, 0x0e, 0x04, 0x29, 0xea, 0x81, 0x27, 0x25, 0x6b, 0x43,
	0x1a, 0x6d, 0x93, 0x98, 0x64, 0xf3, 0x5a, 0xbc, 0x5b, 0x69, 0xd2, 0x6e, 0xe8, 0x51, 0xf3, 0xbc,
	0xcf, 0x47, 0x1a, 0xde, 0x0f, 0x8e, 0xb6, 0xcc, 0x18, 0x5d, 0xa4, 0xe1, 0x93, 0x68, 0xa3, 0x9c,
	0xc2, 0x51, 0xa0, 0xf9, 0x15, 0x1c, 0xbe, 0x32, 0x63, 0x85, 0x92, 0x19, 0xfb, 0xdc, 0x30, 0xeb,
	0xf0, 0x09, 0xc4, 0x2f, 0x96, 0x99, 0x7c, 0x55, 0x92, 0xc1, 0x6c, 0xb0, 0x40, 0x59, 0xd4, 0xe2,
	0xaa, 0x9c, 0x7f, 0x0f, 0x21, 0xee, 0x66, 0xf1, 0x31, 0x44, 0x85, 0x6a, 0x1a, 0xe1, 0xfa, 0x99,
	0x40, 0xf8, 0x02, 0x0e, 0xc2, 0x29, 0xaf, 0xa8, 0xad, 0xc8, 0xd0, 0x4b, 0x08, 0xd1, 0x03, 0xb5,
	0x15, 0x26, 0x10, 0x6f, 0xc3, 0x1d, 0x64, 0xe4, 0x65, 0x8f, 0xf8, 0x0c, 0x10, 0xd5, 0x3a, 0x6f,
	0xe8, 0xbb, 0x32, 0xe4, 0xdf, 0x6c, 0xb0, 0x98, 0x64, 0x63, 0xaa, 0xf5, 0x63, 0xcb, 0x3b, 0x29,
	0xa4, 0x32, 0xe4, 0xff, 0x5e, 0xb6, 0xdc, 0x4b, 0x4d, 0x5d, 0x51, 0x91, 0x68, 0x27, 0x9f, 0x5a,
	0xc6, 0x97, 0x30, 0xf5, 0xd2, 0xb0, 0xdc, 0xb0, 0x9a, 0x51, 0xcb, 0x48, 0xec, 0x7f, 0x3c, 0x69,
	0x47, 0x0c, 0xcb, 0x42, 0x88, 0xcf, 0x01, 0xd6, 0x1b, 0x51, 0x97, 0xb9, 0xa3, 0xdc, 0x92, 0xf1,
	0x6c, 0xb4, 0x40, 0x19, 0xf2, 0xc9, 0x33, 0xe5, 0xb6, 0xd5, 0x5c, 0xe5, 0xfd, 0xd3, 0x91, 0xbf,
	0x01, 0x71, 0xd5, 0xf5, 0x71, 0xb3, 0x04, 0xd4, 0x1d, 0x99, 0xc1, 0x77, 0x00, 0xf7, 0xcc, 0xed,
	0xaa, 0x4a, 0xba, 0xe2, 0x7f, 0xf7, 0x7c, 0x3a, 0xfd, 0x93, 0x2f, 0x93, 0xb7, 0x6b, 0x2e, 0x5c,
	0xb5, 0x59, 0x27, 0x85, 0x6a, 0xd2, 0x5a, 0xf0, 0xca, 0x49, 0x21, 0xb9, 0x64, 0xee, 0x4b, 0x99,
	0x8f, 0xb4, 0x96, 0x65, 0x5a, 0xcb, 0xfd, 0x22, 0xd7, 0x91, 0xdf, 0xe4, 0xed, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0x8f, 0xda, 0xd6, 0xe0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VersionerClient is the client API for Versioner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VersionerClient interface {
	// lncli: `version`
	//GetVersion returns the current version and build information of the running
	//daemon.
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error)
}

type versionerClient struct {
	cc *grpc.ClientConn
}

func NewVersionerClient(cc *grpc.ClientConn) VersionerClient {
	return &versionerClient{cc}
}

func (c *versionerClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/verrpc.Versioner/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionerServer is the server API for Versioner service.
type VersionerServer interface {
	// lncli: `version`
	//GetVersion returns the current version and build information of the running
	//daemon.
	GetVersion(context.Context, *VersionRequest) (*Version, error)
}

func RegisterVersionerServer(s *grpc.Server, srv VersionerServer) {
	s.RegisterService(&_Versioner_serviceDesc, srv)
}

func _Versioner_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionerServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verrpc.Versioner/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionerServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Versioner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "verrpc.Versioner",
	HandlerType: (*VersionerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Versioner_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verrpc/verrpc.proto",
}
