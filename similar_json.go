package main

import (
	"encoding/json"
	"fmt"
)

/*
{
	"traceId": "",
	"code": "000000",
	"message": "请求成功",
	"data": {
		"page": {
			"total": 186,
			"size": 10
		},
		"content": [{
			"companyName": "国泰君安证券股份有限公司",
			"latestReadTime": "2018-11-29"
		}]
	}
}
*/
type content struct {
	companyName    string
	latestReadTime string
}
type page struct {
	total int
	size  int
}
type jsonData struct {
	traceId string
	code    string
	message string
	data    struct {
		page    page
		content []content
	}
}

func main() {
	m := jsonData{
		traceId: "10029",
		code:    "200",
		message: "请求成功",
		data: struct {
			page    page
			content []content
		}{page: page{total: 100, size: 10}, content: []content{
			{"小米科技", "2018-10-10"},
			{"希望集团", "2018-05-12"},
		}},
	}
	m.code = "1002"
	fmt.Println(m)
	data,_ := json.Marshal(m)
	fmt.Println(string(data))
}
