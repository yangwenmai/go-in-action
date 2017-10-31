package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var u User
	h := `{"name":"小张", "age":22}`

	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
	newJson, err := json.Marshal(&u)
	fmt.Println(string(newJson))

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag)
		fmt.Println(sf.Tag.Get("json"), sf.Tag.Get("bson"))
	}
}

// User struct
type User struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}
