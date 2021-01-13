package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"log"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
}

func main() {
	var jsonStr string
	var person Person
	lzh := Person{Name: "lzh", Age: 27, Gender: 1, Height: 173.5, Weight: 85}
	// 序列化: struct -> json
	if result, err := json.Marshal(lzh); err == nil {
		jsonStr = string(result)
		log.Println(jsonStr)
	} else {
		log.Panicln(err)
	}
	// 反序列化: json -> struct
	// 没有的字段会给对应类型的零值
	value := `{"name":"lzh","age":27,"height":173.5,"weight":85}`
	// 可以使用结构体，也可以使用map
	// var dict map[string]interface{}
	if err := json.Unmarshal([]byte(value), &person); err == nil {
		log.Printf("%+v", person)
	} else {
		log.Println(err)
	}

	// gjson
	result := gjson.Get(value, "age")
	logrus.Info(result)

}
