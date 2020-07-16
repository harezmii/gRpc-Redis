package main

import (
	"context"
	"encoding/json"
	"fmt"
	user "gRpc/proto"
	"google.golang.org/grpc"
	"io/ioutil"
)

type UserModel []struct {
	Id int `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name string `json:"last_name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	Ip_Adress string `json:"ip_adress"`
	User_Name string `json:"user_name"`
	Agent string `json:"agent"`
	Country string `json:"country"`

}



func main()  {
	conn , err := grpc.Dial("localhost:8080",grpc.WithInsecure())

	if err != nil {
		print("bağlantıda hata var")
	}
	data ,errrr :=ioutil.ReadFile("data/1.json")
	if errrr!= nil {

	}
	var pd UserModel
	json.Unmarshal(data,&pd)


	for i:=0; i<1000;i++{
		c := user.NewUserServiceClient(conn)
		message := user.User{
			Id: int32(pd[i].Id),
			Country: pd[i].Country,
			Agent: pd[i].Agent,
			Gender: pd[i].Gender,
			Email: pd[i].Email,
			UserName: pd[i].User_Name,
			LastName: pd[i].Last_Name,
			IpAdress: pd[i].Ip_Adress,
			FirstName: pd[i].First_Name,

		}

		response ,error := c.UserWriteRedis(context.Background(),&message)
		if error!= nil{
			fmt.Println("hata var")
		}
		fmt.Print(response)
	}

	defer conn.Close()

}
