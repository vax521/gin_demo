package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Event struct {
	Type     string              `json:"type"`
	Database string              `json:"database"`
	Table    string              `json:"table"`
	Data     []map[string]string `json:"data"`
}

type Blog struct {
	BlogId string `mapstructure:"blogId"`
	Title  string `mapstructure:"title"`
	Uid    int32  `mapstructure:"uid"`
}

func main() {
	msg := []byte(`{
       "type":"UPDATE",
        "database":"blog",
        "table":"blogs",
         "data":[{
          "blogId":"100001",
          "title":"title",
           "uid":"100"
           }]

     }`)
	e := Event{}

	if err := json.Unmarshal(msg, &e); err != nil {
		panic(err)
	}
	fmt.Println(e)

	if data, err := json.Marshal(e); err == nil {
		fmt.Println(string(data))
	}

	if e.Table == "blogs" {
		var blogs []Blog
		//if err:= mapstructure.Decode(e.Data,&blogs); err!=nil{
		if err := mapstructure.WeakDecode(e.Data, &blogs); err != nil {
			panic(err)
		}
		fmt.Println(blogs)
	}

}
