package main

import (
	"fmt"

	"encoding/json"

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
	inicioRangos()
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
			set(mensaje)
			sumoRango(p.Age)
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

func inicioRangos() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
	}
	a, err := conn.Do("SET", "rango1", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(a)
	}
	b, err := conn.Do("SET", "rango2", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(b)
	}
	c, err := conn.Do("SET", "rango3", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(c)
	}
	d, err := conn.Do("SET", "rango4", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(d)
	}
	e, err := conn.Do("SET", "rango5", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(e)
	}
	f, err := conn.Do("SET", "rango6", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(f)
	}
	g, err := conn.Do("SET", "rango7", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(g)
	}
	h, err := conn.Do("SET", "rango8", "0")
	if err != nil {
		fmt.Println(err)
		fmt.Println(h)
	}
}

func sumoRango(rango int) {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	if rango <= 17 && rango >= 12 {
		a, err := c.Do("INCR", "rango1")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 24 && rango >= 18 {
		a, err := c.Do("INCR", "rango2")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 29 && rango >= 25 {
		a, err := c.Do("INCR", "rango3")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 39 && rango >= 30 {
		a, err := c.Do("INCR", "rango4")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 49 && rango >= 40 {
		a, err := c.Do("INCR", "rango5")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 59 && rango >= 50 {
		a, err := c.Do("INCR", "rango6")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 69 && rango >= 60 {
		a, err := c.Do("INCR", "rango7")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
	if rango <= 100 && rango >= 70 {
		a, err := c.Do("INCR", "rango8")
		if err != nil {
			fmt.Println(err)
			fmt.Println(a)
		}
	}
}
