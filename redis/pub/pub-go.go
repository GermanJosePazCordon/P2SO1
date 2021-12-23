package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

func conectar_server(wri http.ResponseWriter, req *http.Request) {
	wri.Header().Set("Access-Control-Allow-Origin", "*")
	wri.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	wri.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		wri.WriteHeader(http.StatusOK)
		wri.Write([]byte("{\"mensaje\": \"ok\"}"))
		fmt.Println("aca entre")
		return
	}
	fmt.Println("es un post")
	datos, _ := ioutil.ReadAll(req.Body)
	fmt.Println("Respuesta del server: ")
	fmt.Println(datos)
	json.NewEncoder(wri).Encode("Se ha almacenado la información")
	bodyString := string(datos)
	log.Print(bodyString)
	publish(bodyString)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", conectar_server)
	fmt.Println("Cliente se levanto en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func publish(mensaje string) {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	c.Do("PUBLISH", "example", mensaje)
}
