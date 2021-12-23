package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "grpcserver/proto"
	//"time"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGetDataServer
}

func guardar_data(data string) {
	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI("**ruta de conexión con mongo**"))
	if err != nil {
		log.Fatal(err)
	}

	databases, err := mongoclient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

	Database := mongoclient.Database("tallerGRPC")
	Collection := Database.Collection("comentarios")

	var bdoc interface{}

	errb := bson.UnmarshalExtJSON([]byte(data), true, &bdoc)
	fmt.Println(errb)

	insertResult, err := Collection.InsertOne(ctx, bdoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult)*/
	fmt.Print(data)
}

func (s *server) ReturnData(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	guardar_data(in.GetData())
	fmt.Printf(">> Hemos recibido la data del cliente: %v\n", in.GetData())
	return &pb.ResponseData{Data: ">> Hola Cliente, he recibido el comentario: " + in.GetData()}, nil
}

func main() {
	escucha, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fallo al levantar el servidor: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetDataServer(s, &server{})
	if err := s.Serve(escucha); err != nil {
		log.Fatalf("Fallo al levantar el servidor: %v", err)
	}
}