// Code generated by protoc-gen-go.
// source: project.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProjectListRequest struct {
	WithMember bool     `protobuf:"varint,1,opt,name=with_member,json=withMember" json:"with_member,omitempty"`
	Members    []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *ProjectListRequest) Reset()                    { *m = ProjectListRequest{} }
func (m *ProjectListRequest) String() string            { return proto.CompactTextString(m) }
func (*ProjectListRequest) ProtoMessage()               {}
func (*ProjectListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ProjectListRequest) GetWithMember() bool {
	if m != nil {
		return m.WithMember
	}
	return false
}

func (m *ProjectListRequest) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ProjectListResponse struct {
	Status  *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Projets []*Project              `protobuf:"bytes,2,rep,name=projets" json:"projets,omitempty"`
}

func (m *ProjectListResponse) Reset()                    { *m = ProjectListResponse{} }
func (m *ProjectListResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectListResponse) ProtoMessage()               {}
func (*ProjectListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ProjectListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ProjectListResponse) GetProjets() []*Project {
	if m != nil {
		return m.Projets
	}
	return nil
}

type ProjectMemberListRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *ProjectMemberListRequest) Reset()                    { *m = ProjectMemberListRequest{} }
func (m *ProjectMemberListRequest) String() string            { return proto.CompactTextString(m) }
func (*ProjectMemberListRequest) ProtoMessage()               {}
func (*ProjectMemberListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ProjectMemberListRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ProjectMemberListResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Projet *Project                `protobuf:"bytes,2,opt,name=projet" json:"projet,omitempty"`
}

func (m *ProjectMemberListResponse) Reset()                    { *m = ProjectMemberListResponse{} }
func (m *ProjectMemberListResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectMemberListResponse) ProtoMessage()               {}
func (*ProjectMemberListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ProjectMemberListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ProjectMemberListResponse) GetProjet() *Project {
	if m != nil {
		return m.Projet
	}
	return nil
}

type Project struct {
	Name             string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Phid             string    `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	Type             string    `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Status           string    `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ViewPolicy       string    `protobuf:"bytes,5,opt,name=view_policy,json=viewPolicy" json:"view_policy,omitempty"`
	EditPolicy       string    `protobuf:"bytes,6,opt,name=edit_policy,json=editPolicy" json:"edit_policy,omitempty"`
	JoinPolicy       string    `protobuf:"bytes,7,opt,name=join_policy,json=joinPolicy" json:"join_policy,omitempty"`
	MembershipLocked bool      `protobuf:"varint,8,opt,name=membership_locked,json=membershipLocked" json:"membership_locked,omitempty"`
	HasSubprojects   bool      `protobuf:"varint,9,opt,name=has_subprojects,json=hasSubprojects" json:"has_subprojects,omitempty"`
	Members          []*Member `protobuf:"bytes,10,rep,name=members" json:"members,omitempty"`
	CreatedAt        int64     `protobuf:"varint,11,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Project) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Project) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Project) GetViewPolicy() string {
	if m != nil {
		return m.ViewPolicy
	}
	return ""
}

func (m *Project) GetEditPolicy() string {
	if m != nil {
		return m.EditPolicy
	}
	return ""
}

func (m *Project) GetJoinPolicy() string {
	if m != nil {
		return m.JoinPolicy
	}
	return ""
}

func (m *Project) GetMembershipLocked() bool {
	if m != nil {
		return m.MembershipLocked
	}
	return false
}

func (m *Project) GetHasSubprojects() bool {
	if m != nil {
		return m.HasSubprojects
	}
	return false
}

func (m *Project) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Project) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Member struct {
	Phid     string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	RealName string `protobuf:"bytes,3,opt,name=real_name,json=realName" json:"real_name,omitempty"`
	IsAdmin  bool   `protobuf:"varint,4,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Member) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Member) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Member) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *Member) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func init() {
	proto.RegisterType((*ProjectListRequest)(nil), "appscode.auth.v1beta1.ProjectListRequest")
	proto.RegisterType((*ProjectListResponse)(nil), "appscode.auth.v1beta1.ProjectListResponse")
	proto.RegisterType((*ProjectMemberListRequest)(nil), "appscode.auth.v1beta1.ProjectMemberListRequest")
	proto.RegisterType((*ProjectMemberListResponse)(nil), "appscode.auth.v1beta1.ProjectMemberListResponse")
	proto.RegisterType((*Project)(nil), "appscode.auth.v1beta1.Project")
	proto.RegisterType((*Member)(nil), "appscode.auth.v1beta1.Member")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Projects service

type ProjectsClient interface {
	List(ctx context.Context, in *ProjectListRequest, opts ...grpc.CallOption) (*ProjectListResponse, error)
	Members(ctx context.Context, in *ProjectMemberListRequest, opts ...grpc.CallOption) (*ProjectMemberListResponse, error)
}

type projectsClient struct {
	cc *grpc.ClientConn
}

func NewProjectsClient(cc *grpc.ClientConn) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) List(ctx context.Context, in *ProjectListRequest, opts ...grpc.CallOption) (*ProjectListResponse, error) {
	out := new(ProjectListResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Projects/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) Members(ctx context.Context, in *ProjectMemberListRequest, opts ...grpc.CallOption) (*ProjectMemberListResponse, error) {
	out := new(ProjectMemberListResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Projects/Members", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Projects service

type ProjectsServer interface {
	List(context.Context, *ProjectListRequest) (*ProjectListResponse, error)
	Members(context.Context, *ProjectMemberListRequest) (*ProjectMemberListResponse, error)
}

func RegisterProjectsServer(s *grpc.Server, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Projects/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).List(ctx, req.(*ProjectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_Members_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectMemberListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).Members(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Projects/Members",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).Members(ctx, req.(*ProjectMemberListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.auth.v1beta1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Projects_List_Handler,
		},
		{
			MethodName: "Members",
			Handler:    _Projects_Members_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}

func init() { proto.RegisterFile("project.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x55, 0x32, 0x65, 0x92, 0xdc, 0xe1, 0x51, 0x8c, 0x80, 0x74, 0x4a, 0x1f, 0x8a, 0x10, 0x94,
	0x87, 0x12, 0x3a, 0x20, 0xe8, 0x92, 0x76, 0xdd, 0xc2, 0x28, 0xdd, 0xb1, 0x89, 0x3c, 0x89, 0xd5,
	0xb8, 0xcc, 0xc4, 0x69, 0xec, 0xb4, 0xaa, 0x10, 0x12, 0xaa, 0xf8, 0x01, 0xc4, 0x7f, 0xb0, 0xe3,
	0x3b, 0x58, 0xf0, 0x0b, 0xac, 0xf9, 0x06, 0xe4, 0x47, 0xe6, 0x21, 0xda, 0x32, 0x62, 0x13, 0x39,
	0xe7, 0x9e, 0xeb, 0x7b, 0x7c, 0xef, 0xb1, 0xe1, 0x5a, 0x59, 0xb1, 0x43, 0x92, 0x8a, 0xb0, 0xac,
	0x98, 0x60, 0xe8, 0x36, 0x2e, 0x4b, 0x9e, 0xb2, 0x8c, 0x84, 0xb8, 0x16, 0x79, 0x78, 0xbc, 0x39,
	0x20, 0x02, 0x6f, 0x76, 0xef, 0x1d, 0x30, 0x76, 0x30, 0x24, 0x11, 0x2e, 0x69, 0x84, 0x8b, 0x82,
	0x09, 0x2c, 0x28, 0x2b, 0xb8, 0x4e, 0xea, 0xae, 0x36, 0x49, 0x17, 0xc4, 0xd7, 0x66, 0xe2, 0x99,
	0x38, 0x2d, 0x09, 0x8f, 0xd4, 0x57, 0x13, 0x82, 0xb7, 0x80, 0xfa, 0x5a, 0xc6, 0x2e, 0xe5, 0x22,
	0x26, 0x47, 0x35, 0xe1, 0x02, 0xad, 0x41, 0xe7, 0x84, 0x8a, 0x3c, 0x19, 0x91, 0xd1, 0x80, 0x54,
	0xbe, 0xb5, 0x6e, 0x6d, 0xb8, 0x31, 0x48, 0x68, 0x4f, 0x21, 0xc8, 0x07, 0x47, 0xc7, 0xb8, 0x6f,
	0xaf, 0xb7, 0x36, 0xbc, 0xb8, 0xf9, 0x0d, 0x3e, 0x59, 0x70, 0x6b, 0x66, 0x47, 0x5e, 0xb2, 0x82,
	0x13, 0x14, 0x41, 0x9b, 0x0b, 0x2c, 0x6a, 0xae, 0x76, 0xeb, 0xf4, 0xee, 0x86, 0xe3, 0xf3, 0x6a,
	0x59, 0xe1, 0xbe, 0x0a, 0xc7, 0x86, 0x86, 0xb6, 0xc0, 0x51, 0x0d, 0x12, 0xba, 0x44, 0xa7, 0xb7,
	0x1a, 0x9e, 0xdb, 0xa1, 0xd0, 0x54, 0x8b, 0x1b, 0x7a, 0xf0, 0x14, 0x7c, 0x83, 0x69, 0xb5, 0xd3,
	0x27, 0x5b, 0x84, 0x56, 0x4d, 0x33, 0xa5, 0xc1, 0x8b, 0xe5, 0x32, 0xf8, 0x6c, 0xc1, 0xd2, 0x39,
	0xf4, 0xff, 0x95, 0xfd, 0x12, 0xda, 0x5a, 0x87, 0x6f, 0xab, 0x84, 0x7f, 0xa9, 0x36, 0xec, 0xe0,
	0xb7, 0x0d, 0x8e, 0xc1, 0x10, 0x82, 0x85, 0x02, 0x8f, 0x88, 0x51, 0xa9, 0xd6, 0x12, 0x2b, 0x73,
	0x9a, 0xa9, 0x5d, 0xbd, 0x58, 0xad, 0x25, 0x26, 0x35, 0xf8, 0x2d, 0x8d, 0xc9, 0x35, 0xba, 0x33,
	0x16, 0xbc, 0xa0, 0xd0, 0x46, 0xd7, 0x1a, 0x74, 0x8e, 0x29, 0x39, 0x49, 0x4a, 0x36, 0xa4, 0xe9,
	0xa9, 0x7f, 0x45, 0x05, 0x41, 0x42, 0x7d, 0x85, 0x48, 0x02, 0xc9, 0xa8, 0x68, 0x08, 0x6d, 0x4d,
	0x90, 0xd0, 0x84, 0x70, 0xc8, 0x68, 0xd1, 0x10, 0x1c, 0x4d, 0x90, 0x90, 0x21, 0x3c, 0x81, 0x9b,
	0xc6, 0x05, 0x39, 0x2d, 0x93, 0x21, 0x4b, 0xdf, 0x93, 0xcc, 0x77, 0x95, 0x77, 0x16, 0x27, 0x81,
	0x5d, 0x85, 0xa3, 0x87, 0x70, 0x23, 0xc7, 0x3c, 0xe1, 0xf5, 0xc0, 0x5c, 0x03, 0xee, 0x7b, 0x8a,
	0x7a, 0x3d, 0xc7, 0x7c, 0x7f, 0x82, 0xa2, 0x57, 0x13, 0xab, 0x81, 0xf2, 0xc1, 0xca, 0x05, 0x1d,
	0xd5, 0xd3, 0x1b, 0x3b, 0x11, 0xad, 0x00, 0xa4, 0x15, 0xc1, 0x82, 0x64, 0x09, 0x16, 0x7e, 0x67,
	0xdd, 0xda, 0x68, 0xc5, 0x9e, 0x41, 0xb6, 0x45, 0x70, 0x04, 0x6d, 0x63, 0xe6, 0xa6, 0xb5, 0xd6,
	0x54, 0x6b, 0x97, 0xc1, 0xab, 0x39, 0xa9, 0x12, 0x35, 0x07, 0xdd, 0x73, 0x57, 0x02, 0x6f, 0xe4,
	0x2c, 0x96, 0xc1, 0xab, 0x08, 0x1e, 0xea, 0xa0, 0x6e, 0xbe, 0x2b, 0x01, 0x15, 0x5c, 0x02, 0x97,
	0xf2, 0x04, 0x67, 0x23, 0x5a, 0xa8, 0x11, 0xb8, 0xb1, 0x43, 0xf9, 0xb6, 0xfc, 0xed, 0xfd, 0xb0,
	0xc1, 0xed, 0x37, 0xe7, 0xfa, 0x62, 0xc1, 0x82, 0xb4, 0x1a, 0x7a, 0x74, 0xb9, 0x43, 0xa6, 0xdc,
	0xdb, 0x7d, 0x3c, 0x0f, 0x55, 0x3b, 0x37, 0xe8, 0x9d, 0x7d, 0xf7, 0x6d, 0xd7, 0x3a, 0xfb, 0xf9,
	0xeb, 0xab, 0xfd, 0x00, 0xdd, 0x8f, 0x92, 0xd9, 0xa7, 0xa2, 0x16, 0x79, 0x64, 0xd2, 0xa3, 0x71,
	0xaf, 0xbf, 0x59, 0xe0, 0xec, 0x99, 0xf6, 0x45, 0x97, 0xd7, 0xfa, 0xeb, 0x6a, 0x75, 0x9f, 0xcd,
	0x9f, 0x60, 0x24, 0xbe, 0x9e, 0x92, 0xf8, 0x02, 0xf5, 0xe6, 0x91, 0x18, 0x7d, 0xa8, 0x69, 0xf6,
	0x31, 0x32, 0x33, 0xde, 0xd9, 0x82, 0x95, 0x94, 0x8d, 0xa6, 0x0a, 0x97, 0x74, 0xa6, 0xf8, 0xce,
	0x55, 0x53, 0xbd, 0x2f, 0x5f, 0xbb, 0xbe, 0xf5, 0xce, 0x31, 0x81, 0x41, 0x5b, 0xbd, 0x7f, 0xcf,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xc1, 0xd3, 0xb0, 0x86, 0x05, 0x00, 0x00,
}
