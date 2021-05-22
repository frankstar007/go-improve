/**
 * created by
 * @project go-improve
 * @author frankstar
 * @date 2021/5/22
 * @contact frankstarye@tencent.com
 **/

package main

import (
	"context"
	"github.com/frankstar007/go-improve/service-api/github.com/frankstar007/improve/go-improve/service-api/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50001"

)


type UserService struct {
	//实现User服务的业务的对象
}

func (UserService *UserService) GetUser(ctx context.Context, in *user.UserRequest) (out *user.UserResponse, err error)  {
	log.Printf("page %d, pageSize %d", in.GetPage(), in.GetPageSize())
	return &user.UserResponse{
		Err: 0,
		Msg: "Success",
		Data: []*user.UserModel {
			{Name: "frankstar", Age: 30},
			{Name: "zoe", Age: 29},
		},
	}, nil
}

func (UserService *UserService) UserView(ctx context.Context, in *user.UserViewRequest) (*user.UserViewResponse, error)  {

	log.Printf("request uid %d", in.Uid)
	return &user.UserViewResponse{
		Err: 0,
		Msg: "Success",
		Data: &user.UserModel{
			Name: "frankstar", Age: 30,
		},
	},nil
}

func (UserService *UserService) UserPost(ctx context.Context, in *user.UserPostRequest) (*user.UserPostResponse, error)  {
	log.Printf("UserPost uid %d", in.Uid)
	return &user.UserPostResponse{
		Err: 0,
		Msg: "Success",
	},nil
}

func (UserService *UserService) UserDelete(ctx context.Context, in *user.UserDeleteRequest) (*user.UserDeleteResponse, error)  {
	log.Printf("UserDelete uid %d", in.Uid)
	return &user.UserDeleteResponse{
		Err: 0,
		Msg: "Success",
	}, nil
}



func main()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 RPC 服务容器
	grpcServer := grpc.NewServer()

	// 为 User 服务注册业务实现 将 User 服务绑定到 RPC 服务容器上
	user.RegisterUserServer(grpcServer, &UserService{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}
