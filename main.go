package main

import (
	"context"
	"fmt"
	user "gRpc/proto"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"log"
	"net"
)

type serv struct {


}

func main()  {

	lis ,errr := net.Listen("tcp","localhost:8080")

	if errr !=nil {
		fmt.Println(errr)
	}

	server := grpc.NewServer()
	user.RegisterUserServiceServer(server,&serv{})
	fmt.Println("Grpc Server Başladı")
	err := server.Serve(lis)
	log.Fatal(err)

}

func (*serv) UserWriteRedis(ctx context.Context, req *user.User) (*user.UserResponse, error){
	fmt.Println("redis eklendi")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 1,
	})

	pong ,err := client.Ping().Result()
	fmt.Println(pong,err)

	err = client.Set(string(req.Id),req.String(),0).Err()

	if err!= nil {
		fmt.Println(err)
		return &user.UserResponse{
			Success: false,
		},nil;
	}
	return &user.UserResponse{
		Success: true,
	},nil;

}
