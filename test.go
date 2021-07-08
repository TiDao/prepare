package main

import (
	"flag"
	"fmt"
)

type URLValue struct {
	URL string
}

func (v URLValue) String() string {
	return ""
}

func (v URLValue) Set(s string) error {
	v.URL = s
	return nil
}


func main() {
	u := &URLValue{}
	u.URL = "test2"
	flag.Var(u, "url", "URL to parse")
	flag.Parse()
	fmt.Println(u)

}
