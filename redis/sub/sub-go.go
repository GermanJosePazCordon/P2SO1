package main

import (
	"fmt"

	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

type Persona struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

func main() {

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}

	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe("example")

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			mensaje := string([]byte(v.Data))
			//m := strings.ReplaceAll(mensaje, "\n", "")
			//fmt.Println(mensaje)
			//fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			var p Persona
			json.Unmarshal([]byte(mensaje), &p)
			fmt.Println(string(p.Title))
			out, _ := json.Marshal(&p)
			fmt.Println(string(out))
			set(mensaje)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}

}

func set(mensaje string) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
	}
	a, err := conn.Do("lpush", "personas", mensaje)
	if err != nil {
		fmt.Println(err)
		fmt.Println(a)
	}
}
