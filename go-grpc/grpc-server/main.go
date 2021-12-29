package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "grpcserver/proto"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/gomodule/redigo/redis"
)

const (
	port = ":5000"
	url  = "redis://miguelesdb@34.125.174.190:6379"
)

type server struct {
	pb.UnimplementedGetDataServer
}

type Persona struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	Vaccine_type string `json:"vaccine_type"`
	N_dose       string `json:"n_dose"`
}

func guardar_data(data string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://miel:miguelesdb@34.125.174.190:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}

	databases, err := mongoclient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

	Database := mongoclient.Database("vacunadosData")
	Collection := Database.Collection("vacunados")

	var bdoc interface{}

	errb := bson.UnmarshalExtJSON([]byte(data), true, &bdoc)
	fmt.Println(errb)

	insertResult, err := Collection.InsertOne(ctx, bdoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult)
}

func (s *server) ReturnData(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	guardar_data(in.GetData())
	set(in.GetData())
	mensaje := string(in.GetData())
	var p Persona
	json.Unmarshal([]byte(mensaje), &p)
	sumoRango(p.Age)
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

//-------------------------------------REDIS---------------------------------------------------------------
func set(mensaje string) {
	conn, err := redis.DialURL(url)
	if err != nil {
		fmt.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
	}
	a, err := conn.Do("lpush", "personas", mensaje)
	if err != nil {
		fmt.Println(err)
		fmt.Println(a)
	}
}

func sumoRango(rango int) {
	c, err := redis.DialURL(url)
	if err != nil {
		fmt.Println(err)
	}
	if rango <= 11 && rango >= 0 {
		a, err := c.Do("INCR", "rango1")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 18 && rango >= 12 {
		a, err := c.Do("INCR", "rango2")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 26 && rango >= 19 {
		a, err := c.Do("INCR", "rango3")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 59 && rango >= 27 {
		a, err := c.Do("INCR", "rango4")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango >= 60 {
		a, err := c.Do("INCR", "rango5")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}

}
