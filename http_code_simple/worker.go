package main

import (
	"fmt"
	"github.com/myzhan/boomer"
	gorequest "github.com/parnurzeal/Gorequest"
	"github.com/tidwall/gjson"
	"time"
)

var request = gorequest.New()

func Worker() {
	url := "http://www.baidu.com"
	startTime := time.Now()
	_, body, errs := request.Get(url).
		Set("Authorization", "bearer add1ae9b-47de-4758-ad6e-e6c5d7801edd").End()
	t := time.Since(startTime)
	if len(errs) != 0 {
		boomer.RecordFailure("HTTP", "app list", t.Milliseconds(), errs[0].Error())
		return
	}
	if value := gjson.Get(body, "code"); value.Int() != 1 {
		boomer.RecordFailure("HTTP", "app list", t.Milliseconds(), string(rune(len(body))))
		return
	}
	boomer.RecordSuccess("HTTP", "app list", t.Milliseconds(), int64(len(body)))
	fmt.Println(len(body))
}

//func Worker() {
//	url := "http://login-uat.wenxiangcn.com/api/upmsx/admin/application/current/v2"
//	result := make(map[string]string)
//	resp, err := Get(url)
//	if err != nil{
//		return
//	}
//	fmt.Println(resp.Json())
//	fmt.Println(result)
//	println(resp.Text())
//}
