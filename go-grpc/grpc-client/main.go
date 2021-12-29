package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	pb "grpclient/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	address = ":5000"
)

func conectar_server(wri http.ResponseWriter, req *http.Request) {
	fmt.Println("ENTRANDO A CONECTAR SERVER")
	wri.Header().Set("Access-Control-Allow-Origin", "*")
	wri.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	wri.Header().Set("Content-Type", "application/json")

	if req.Method == "GET" {
		wri.WriteHeader(http.StatusOK)
		wri.Write([]byte("{\"mensaje\": \"ok\"}"))
		return
	}
	fmt.Println("0")
	datos, _ := ioutil.ReadAll(req.Body)
	fmt.Println("1")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	fmt.Println("2")
	if err != nil {
		json.NewEncoder(wri).Encode("Error, no se puede conectar con el servidor grpc")
		log.Fatalf("No se puede conectar con el server :c (%v)", err)
	}
	fmt.Println("3")
	defer conn.Close()
	fmt.Println("4")
	cl := pb.NewGetDataClient(conn)
	fmt.Println("5")
	received_data := string(datos)
	fmt.Print("received : ")
	fmt.Println(received_data)
	if len(os.Args) > 1 {
		received_data = os.Args[1]
	}
	fmt.Println("6")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	fmt.Println("7")
	defer cancel()
	fmt.Println("8")
	fmt.Println("ANTES RETURN")
	ret, err := cl.ReturnData(ctx, &pb.RequestData{Data: received_data})
	fmt.Println("DESPUES RETURN")

	if err != nil {
		json.NewEncoder(wri).Encode("Error, no  se puede retornar la información.")
		log.Fatalf("No se puede retornar la información :c (%v)", err)
	}
	fmt.Println("9")
	log.Printf("Respuesta del server: %s\n", ret.GetData())
	fmt.Println("10")
	json.NewEncoder(wri).Encode("Se ha almacenado la información")
	fmt.Println("11")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", conectar_server)
	fmt.Println("Cliente se levanto en el puerto 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
