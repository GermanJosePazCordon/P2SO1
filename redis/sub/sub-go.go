package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gomodule/redigo/redis"
)

type Persona struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	Vaccine_type string `json:"vaccine_type"`
	N_dose       string `json:"n_dose"`
}

func main() {

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	//inicioRangos()
	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe("vacunados")

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			mensaje := string([]byte(v.Data))
			//m := strings.ReplaceAll(mensaje, "\n", "")
			//fmt.Println(mensaje)
			//fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			var p Persona
			json.Unmarshal([]byte(mensaje), &p)
			//fmt.Println(string(p.Title))
			out, _ := json.Marshal(&p)
			fmt.Println(string(out))
			//set(mensaje)
			sumoRango(p.Age)
			guardar_data(mensaje)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}

}

/*func set(mensaje string) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
	}
	a, err := conn.Do("lpush", "personas", mensaje)
	if err != nil {
		fmt.Println(err)
		fmt.Println(a)
	}
}*/

func sumoRango(rango int) {
	c, err := redis.Dial("tcp", "localhost:6379")
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

//--------------------------------------------MONGO------------------------------------------------------------------
func guardar_data(data string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
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
