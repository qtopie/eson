package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// $[\\.[w+]\\]
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

	pattern := "$.key3"
	if !validatePattern(pattern) {
		panic("invalid pattern")
	}

	var data interface{}
	data = dat
	s := strings.TrimPrefix(pattern, "$.")
	parts := strings.Split(s, ".")
	for _, part := range parts {
		key := part
		// index := -1
		if strings.HasSuffix(part, "]") {
			key = string(part[0:len(part) - 3])
			// index = 
		}
	
		// parse data
		data = (data.(map[string]interface{}))[key]
	}

	fmt.Println(data)

	// var result map[string]interface{}
	// decoder := json.NewDecoder(bytes.NewReader(byt))
	// decoder.UseNumber() // 使用 UseNumber(),保留float64精度

	// err = decoder.Decode(&result)
	// if err != nil {
	// 	panic(err)
	// }

}

func validatePattern(s string) bool {
	pattern := "^\\$(\\.\\w+(\\[\\d+\\])?)+$"
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}

	return matched
}
