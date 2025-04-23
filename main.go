package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
)

func main() {
	// $.abc.def[2]
	// support linux pipe
	fileFlag := flag.String("file", "", "JSON File")
	flag.Parse()
	fmt.Println(*fileFlag)

	byt, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var dat map[string]interface{}
	if err = json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	s1 := dat["key1"].(string)
	fmt.Println(s1)

	t2 := dat["key2"].(float64)
	fmt.Println(t2)

	var result map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(byt))
	decoder.UseNumber() // 使用 UseNumber(),保留float64精度

	err = decoder.Decode(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result["key2"])
	fmt.Println(reflect.TypeOf(result["key2"]))
}
