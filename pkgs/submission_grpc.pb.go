// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: submission.proto

package pkgs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Submission_SubmitSnapshot_FullMethodName = "/submission.Submission/SubmitSnapshot"
)

// SubmissionClient is the client API for Submission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionClient interface {
	SubmitSnapshot(ctx context.Context, opts ...grpc.CallOption) (Submission_SubmitSnapshotClient, error)
}

type submissionClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionClient(cc grpc.ClientConnInterface) SubmissionClient {
	return &submissionClient{cc}
}

func (c *submissionClient) SubmitSnapshot(ctx context.Context, opts ...grpc.CallOption) (Submission_SubmitSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &Submission_ServiceDesc.Streams[0], Submission_SubmitSnapshot_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &submissionSubmitSnapshotClient{stream}
	return x, nil
}

type Submission_SubmitSnapshotClient interface {
	Send(*SnapshotSubmission) error
	CloseAndRecv() (*SubmissionResponse, error)
	grpc.ClientStream
}

type submissionSubmitSnapshotClient struct {
	grpc.ClientStream
}

func (x *submissionSubmitSnapshotClient) Send(m *SnapshotSubmission) error {
	return x.ClientStream.SendMsg(m)
}

func (x *submissionSubmitSnapshotClient) CloseAndRecv() (*SubmissionResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SubmissionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubmissionServer is the server API for Submission service.
// All implementations must embed UnimplementedSubmissionServer
// for forward compatibility
type SubmissionServer interface {
	SubmitSnapshot(Submission_SubmitSnapshotServer) error
	mustEmbedUnimplementedSubmissionServer()
}

// UnimplementedSubmissionServer must be embedded to have forward compatible implementations.
type UnimplementedSubmissionServer struct {
}

func (UnimplementedSubmissionServer) SubmitSnapshot(Submission_SubmitSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method SubmitSnapshot not implemented")
}
func (UnimplementedSubmissionServer) mustEmbedUnimplementedSubmissionServer() {}

// UnsafeSubmissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionServer will
// result in compilation errors.
type UnsafeSubmissionServer interface {
	mustEmbedUnimplementedSubmissionServer()
}

func RegisterSubmissionServer(s grpc.ServiceRegistrar, srv SubmissionServer) {
	s.RegisterService(&Submission_ServiceDesc, srv)
}

func _Submission_SubmitSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SubmissionServer).SubmitSnapshot(&submissionSubmitSnapshotServer{stream})
}

type Submission_SubmitSnapshotServer interface {
	SendAndClose(*SubmissionResponse) error
	Recv() (*SnapshotSubmission, error)
	grpc.ServerStream
}

type submissionSubmitSnapshotServer struct {
	grpc.ServerStream
}

func (x *submissionSubmitSnapshotServer) SendAndClose(m *SubmissionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *submissionSubmitSnapshotServer) Recv() (*SnapshotSubmission, error) {
	m := new(SnapshotSubmission)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Submission_ServiceDesc is the grpc.ServiceDesc for Submission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Submission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "submission.Submission",
	HandlerType: (*SubmissionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitSnapshot",
			Handler:       _Submission_SubmitSnapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "submission.proto",
}
