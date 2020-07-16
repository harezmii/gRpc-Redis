import grpc
import client.user_pb2 as d
import client.user_pb2_grpc as grpc2
import json


channel = grpc.insecure_channel("localhost:8080")
stub = grpc2.UserServiceStub(channel)
stub.UserWriteRedis(json.dumps('{"id":2,"first_name":"Ingar","last_name":"Ashment","email":"iashment1@ning.com","gender":"Male","ip_address":"70.1.87.28","user_name":"iashment1","agent":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.20 Safari/535.1","country":"Sweden"}'))