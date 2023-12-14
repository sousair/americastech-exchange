package user_client

import (
	"context"

	"github.com/sousair/americastech-exchange/pkg/user_client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type ()

type UserServiceGrpcClient struct {
	client pb.UserServiceClient
}

func NewUserServiceGRPCClient(userApiUrl string) *UserServiceGrpcClient {
	conn, err := grpc.Dial(userApiUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := pb.NewUserServiceClient(conn)

	return &UserServiceGrpcClient{
		client: client,
	}
}

func (c UserServiceGrpcClient) ValidateUserToken(ctx context.Context, token string) (bool, error) {
	resp, err := c.client.ValidateUserToken(ctx, &pb.ValidateUserTokenRequest{
		Token: token,
	})

	if err != nil {
		return false, err
	}

	return resp.Valid, nil
}
