package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.zhihu.com/api/v3/feed/topstory/hot-list-wx?limit=50"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
