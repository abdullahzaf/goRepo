package vault

import (
    "golang.org/x/net/context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/abdullahzaf/goRepo/vault/pb"
)
// NewGRPCServer gets a new pb.VaultServer.
func NewGRPCServer(ctx context.Context, endpoints Endpoints) pb.VaultServer {
	return &grpcServer{
		hash: grpctransport.NewServer(
			ctx,
			endpoints.HashEndpoint,
			DecodeGRPCHashRequest,
			EncodeGRPCHashResponse,
		),
		validate: grpctransport.NewServer(
			ctx,
			endpoints.ValidateEndpoint,
			DecodeGRPCValidateRequest,
			EncodeGRPCValidateResponse,
		),
	}
}

type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

func (s *grpcServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HashResponse), nil
}

func (s *grpcServer) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ValidateResponse), nil
}
//EncodeGRPCHashRequest : encode request
func EncodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(hashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}
//DecodeGRPCHashRequest : decode request
func DecodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HashRequest)
	return hashRequest{Password: req.Password}, nil
}
//EncodeGRPCHashResponse : encode
func EncodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(hashResponse)
	return &pb.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}
//DecodeGRPCHashResponse : decode
func DecodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.HashResponse)
	return hashResponse{Hash: res.Hash, Err: res.Err}, nil
}
//EncodeGRPCValidateRequest : validate request
func EncodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(validateRequest)
	return &pb.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}
//DecodeGRPCValidateRequest : decode
func DecodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ValidateRequest)
	return validateRequest{Password: req.Password, Hash: req.Hash}, nil
}
//EncodeGRPCValidateResponse : encode
func EncodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(validateResponse)
	return &pb.ValidateResponse{Valid: res.Valid}, nil
}
//DecodeGRPCValidateResponse : decode
func DecodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ValidateResponse)
	return validateResponse{Valid: res.Valid}, nil
}
