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
	address = "grpc-server:5000"
)

func conectar_server(wri http.ResponseWriter, req *http.Request) {
	wri.Header().Set("Access-Control-Allow-Origin", "*")
	wri.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	wri.Header().Set("Content-Type", "application/json")

	if req.Method == "GET" {
		wri.WriteHeader(http.StatusOK)
		wri.Write([]byte("{\"mensaje\": \"ok green\"}"))
		return
	}
	datos, _ := ioutil.ReadAll(req.Body)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		json.NewEncoder(wri).Encode("Error, no se puede conectar con el servidor grpc")
		log.Fatalf("No se puede conectar con el server :c (%v)", err)
	}
	defer conn.Close()
	cl := pb.NewGetDataClient(conn)
	received_data := string(datos)
	if len(os.Args) > 1 {
		received_data = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	ret, err := cl.ReturnData(ctx, &pb.RequestData{Data: received_data})

	if err != nil {
		json.NewEncoder(wri).Encode("Error, no  se puede retornar la información.")
		log.Fatalf("No se puede retornar la información :c (%v)", err)
	}
	log.Printf("Respuesta del server: %s\n", ret.GetData())
	json.NewEncoder(wri).Encode("green deployment")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", conectar_server)
	fmt.Println("Cliente se levanto en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
