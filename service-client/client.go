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
	"fmt"
	"github.com/frankstar007/go-improve/service-api/github.com/frankstar007/improve/go-improve/service-api/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

const address = "localhost:50001"

func main()  {

	//建立连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("connect error : %v", err)
	}
	defer conn.Close()

	userClient := user.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	userIndex, err := userClient.GetUser(ctx, &user.UserRequest{
		Page: 0,
		PageSize: 10,

	})

	if err != nil {
		log.Printf("invalid userRequest %v", err)
	}

	if 0 == userIndex.Err {
		log.Printf("user index success %s", userIndex.Msg)
		// 包含 UserEntity 的数组列表
		userEntityList := userIndex.Data
		for _, row := range userEntityList {
			fmt.Println(row.Name, row.Age)
		}
	}else {
		log.Printf("user index error: %d", userIndex.Err)
	}

	// UserView 请求
	userViewResponse, err := userClient.UserView(ctx, &user.UserViewRequest{Uid: 1})
	if err != nil {
		log.Printf("user view could not greet: %v", err)
	}

	if 0 == userViewResponse.Err {
		log.Printf("user view success: %s", userViewResponse.Msg)
		userEntity := userViewResponse.Data
		fmt.Println(userEntity.Name, userEntity.Age)
	} else {
		log.Printf("user view error: %d", userViewResponse.Err)
	}

	// UserPost 请求
	userPostReponse, err := userClient.UserPost(ctx, &user.UserPostRequest{
		Uid: 9})
	if err != nil {
		log.Printf("user post could not greet: %v", err)
	}

	if 0 == userPostReponse.Err {
		log.Printf("user post success: %s", userPostReponse.Msg)
	} else {
		log.Printf("user post error: %d", userPostReponse.Err)
	}

	// UserDelete 请求
	userDeleteReponse, err := userClient.UserDelete(ctx, &user.UserDeleteRequest{Uid: 1})
	if err != nil {
		log.Printf("user delete could not greet: %v", err)
	}

	if 0 == userDeleteReponse.Err {
		log.Printf("user delete success: %s", userDeleteReponse.Msg)
	} else {
		log.Printf("user delete error: %d", userDeleteReponse.Err)
	}





}