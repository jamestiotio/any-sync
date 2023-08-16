// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.33
// source: identityrepo/identityrepoproto/protos/identityrepo.proto

package identityrepoproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto struct{}

func (drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCIdentityRepoClient interface {
	DRPCConn() drpc.Conn

	DataPut(ctx context.Context, in *DataPutRequest) (*Ok, error)
	DataDelete(ctx context.Context, in *DataDeleteRequest) (*Ok, error)
	DataPull(ctx context.Context, in *DataPullRequest) (*DataPullResponse, error)
}

type drpcIdentityRepoClient struct {
	cc drpc.Conn
}

func NewDRPCIdentityRepoClient(cc drpc.Conn) DRPCIdentityRepoClient {
	return &drpcIdentityRepoClient{cc}
}

func (c *drpcIdentityRepoClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcIdentityRepoClient) DataPut(ctx context.Context, in *DataPutRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/identityRepo.IdentityRepo/DataPut", drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcIdentityRepoClient) DataDelete(ctx context.Context, in *DataDeleteRequest) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/identityRepo.IdentityRepo/DataDelete", drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcIdentityRepoClient) DataPull(ctx context.Context, in *DataPullRequest) (*DataPullResponse, error) {
	out := new(DataPullResponse)
	err := c.cc.Invoke(ctx, "/identityRepo.IdentityRepo/DataPull", drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCIdentityRepoServer interface {
	DataPut(context.Context, *DataPutRequest) (*Ok, error)
	DataDelete(context.Context, *DataDeleteRequest) (*Ok, error)
	DataPull(context.Context, *DataPullRequest) (*DataPullResponse, error)
}

type DRPCIdentityRepoUnimplementedServer struct{}

func (s *DRPCIdentityRepoUnimplementedServer) DataPut(context.Context, *DataPutRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCIdentityRepoUnimplementedServer) DataDelete(context.Context, *DataDeleteRequest) (*Ok, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCIdentityRepoUnimplementedServer) DataPull(context.Context, *DataPullRequest) (*DataPullResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCIdentityRepoDescription struct{}

func (DRPCIdentityRepoDescription) NumMethods() int { return 3 }

func (DRPCIdentityRepoDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/identityRepo.IdentityRepo/DataPut", drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCIdentityRepoServer).
					DataPut(
						ctx,
						in1.(*DataPutRequest),
					)
			}, DRPCIdentityRepoServer.DataPut, true
	case 1:
		return "/identityRepo.IdentityRepo/DataDelete", drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCIdentityRepoServer).
					DataDelete(
						ctx,
						in1.(*DataDeleteRequest),
					)
			}, DRPCIdentityRepoServer.DataDelete, true
	case 2:
		return "/identityRepo.IdentityRepo/DataPull", drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCIdentityRepoServer).
					DataPull(
						ctx,
						in1.(*DataPullRequest),
					)
			}, DRPCIdentityRepoServer.DataPull, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterIdentityRepo(mux drpc.Mux, impl DRPCIdentityRepoServer) error {
	return mux.Register(impl, DRPCIdentityRepoDescription{})
}

type DRPCIdentityRepo_DataPutStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcIdentityRepo_DataPutStream struct {
	drpc.Stream
}

func (x *drpcIdentityRepo_DataPutStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCIdentityRepo_DataDeleteStream interface {
	drpc.Stream
	SendAndClose(*Ok) error
}

type drpcIdentityRepo_DataDeleteStream struct {
	drpc.Stream
}

func (x *drpcIdentityRepo_DataDeleteStream) SendAndClose(m *Ok) error {
	if err := x.MsgSend(m, drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCIdentityRepo_DataPullStream interface {
	drpc.Stream
	SendAndClose(*DataPullResponse) error
}

type drpcIdentityRepo_DataPullStream struct {
	drpc.Stream
}

func (x *drpcIdentityRepo_DataPullStream) SendAndClose(m *DataPullResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_identityrepo_identityrepoproto_protos_identityrepo_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}